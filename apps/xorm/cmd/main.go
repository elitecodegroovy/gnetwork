package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/elitecodegroovy/gnetwork/pkg/bus"
	"github.com/elitecodegroovy/gnetwork/pkg/infra/localcache"
	"github.com/elitecodegroovy/gnetwork/pkg/infra/log"
	"github.com/elitecodegroovy/gnetwork/pkg/middleware"
	"github.com/elitecodegroovy/gnetwork/pkg/registry"
	"github.com/elitecodegroovy/gnetwork/pkg/routing"
	"github.com/elitecodegroovy/gnetwork/pkg/setting"
	"github.com/facebookgo/inject"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"golang.org/x/sync/errgroup"
	sLog "log"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"runtime/trace"
	"strconv"
	"syscall"
	"time"
)

var engine *xorm.Engine

var version = "1.0.0"
var commit = "NA"
var buildBranch = "master"
var buildstamp string

var configFile = flag.String("config", "", "path to config file")
var homePath = flag.String("homepath", "", "path to gnetwork install/home path, defaults to working directory")
var pidFile = flag.String("pidfile", "", "path to pid file")
var packaging = flag.String("packaging", "unknown", "describes the way gnetwork was installed")

func Init() {
	mLog := log.New("xorm_log")
	mLog.Info("DB Initialization starts !", "0", "++")

	var err error
	//
	engine, err = xorm.NewEngine("mysql", "gc:test123456#G@192.168.1.229:3306/oa-case?charset=utf8")
	if err != nil {
		log.Info(" get engine with an error : {}", err.Error())
		os.Exit(1)
	}

	//setting of log
	f, err := os.Create("mysql_log.log")
	if err != nil {
		log.Info(" mysql_log : {}", err.Error())
		return
	}
	mLog.Info("DB Initialization starts !", "0", "--")
	engine.SetLogger(xorm.NewSimpleLogger(f))
	//show sql statement
	engine.ShowSQL(true)
	engine.ShowExecTime(true)
	//
	engine.SetMaxOpenConns(0)
	engine.SetMaxIdleConns(2)
	engine.SetConnMaxLifetime(time.Second * time.Duration(14400))

	mLog.Info("Init DB successfully!")
}

func validPackaging(packaging string) string {
	validTypes := []string{"dev", "deb", "rpm", "docker", "brew", "hosted", "unknown"}
	for _, vt := range validTypes {
		if packaging == vt {
			return packaging
		}
	}
	return "unknown"
}

//Setting system signal handling
func listenToSystemSignals(server *GNetworkServerImpl) {
	signalChan := make(chan os.Signal, 1)
	sighupChan := make(chan os.Signal, 1)

	signal.Notify(sighupChan, syscall.SIGHUP)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)

	for {
		select {
		case <-sighupChan:
			log.Reload()
		case sig := <-signalChan:
			server.Shutdown(fmt.Sprintf("System signal: %s", sig))
		}
	}
}

func NewGNetworkServer() *GNetworkServerImpl {
	//a copy of parent
	rootCtx, shutdownFn := context.WithCancel(context.Background())
	//a new Group
	childRoutines, childCtx := errgroup.WithContext(rootCtx)

	return &GNetworkServerImpl{
		context:       childCtx,
		shutdownFn:    shutdownFn,
		childRoutines: childRoutines,
		log:           log.New("server"),
		cfg:           setting.NewCfg(),
	}
}

type GNetworkServerImpl struct {
	context            context.Context
	shutdownFn         context.CancelFunc
	childRoutines      *errgroup.Group
	log                log.Logger
	cfg                *setting.Cfg
	shutdownReason     string
	shutdownInProgress bool
}

func (g *GNetworkServerImpl) loadConfiguration() {
	err := g.cfg.Load(&setting.CommandLineArgs{
		Config:   *configFile,
		HomePath: *homePath,
		Args:     flag.Args(),
	})

	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to start grafana. error: %s\n", err.Error())
		os.Exit(1)
	}

	g.log.Info("Starting "+setting.ApplicationName, "version", version, "commit", commit, "branch", buildBranch, "compiled", time.Unix(setting.BuildStamp, 0))
	g.cfg.LogConfigSources()
	g.log.Info("loading configuration successfully!")
}

func (g *GNetworkServerImpl) Run() error {
	var err error
	g.loadConfiguration()

	return err
}

func (g *GNetworkServerImpl) Shutdown(reason string) {
	g.log.Info("Shutdown started", "reason", reason)
	g.shutdownReason = reason
	g.shutdownInProgress = true

	// call cancel func on root context
	g.shutdownFn()

	// wait for child routines
	g.childRoutines.Wait()
}

func (g *GNetworkServerImpl) Exit(reason error) int {
	// default exit code is 1
	code := 1

	if reason == context.Canceled && g.shutdownReason != "" {
		reason = fmt.Errorf(g.shutdownReason)
		code = 0
	}

	g.log.Error("Server shutdown", "reason", reason)
	return code
}

func main() {
	sLog.SetOutput(os.Stdout)
	sLog.SetFlags(0)

	v := flag.Bool("v", false, "prints current version and exits")
	profile := flag.Bool("profile", false, "Turn on pprof profiling")
	profilePort := flag.Int("profile-port", 6060, "Define custom port for profiling")
	flag.Parse()
	if *v {
		fmt.Printf("Version %s (commit: %s, branch: %s)\n", version, commit, buildBranch)
		os.Exit(0)
	}

	if *profile {
		runtime.SetBlockProfileRate(1)
		go func() {
			err := http.ListenAndServe(fmt.Sprintf("localhost:%d", *profilePort), nil)
			if err != nil {
				panic(err)
			}
		}()

		f, err := os.Create("trace.out")
		if err != nil {
			panic(err)
		}
		defer f.Close()

		err = trace.Start(f)
		if err != nil {
			panic(err)
		}
		defer trace.Stop()
	}

	buildstampInt64, _ := strconv.ParseInt(buildstamp, 10, 64)
	if buildstampInt64 == 0 {
		buildstampInt64 = time.Now().Unix()
	}

	setting.BuildVersion = version
	setting.BuildCommit = commit
	setting.BuildStamp = buildstampInt64
	setting.BuildBranch = buildBranch
	setting.Packaging = validPackaging(*packaging)
	sLog.Printf("Version: %s, Commit Version: %s, Package Iteration: %s\n", version, setting.BuildCommit, setting.BuildBranch)

	//Init()

	server := NewGNetworkServer()

	go listenToSystemSignals(server)

	err := server.Run()

	code := server.Exit(err)
	trace.Stop()
	log.Close()

	os.Exit(code)

}
