package main

import (
	"context"
	"fmt"

	"github.com/elitecodegroovy/gnetwork/apps/micro/rpc4/auth/handler"
	"github.com/elitecodegroovy/gnetwork/apps/micro/rpc4/auth/model"
	_ "github.com/elitecodegroovy/gnetwork/apps/micro/rpc4/auth/plugin"
	s "github.com/elitecodegroovy/gnetwork/apps/micro/rpc4/auth/proto/auth"
	"github.com/elitecodegroovy/gnetwork/apps/micro/rpc4/basic"
	"github.com/elitecodegroovy/gnetwork/apps/micro/rpc4/basic/common"
	"github.com/elitecodegroovy/gnetwork/apps/micro/rpc4/basic/config"
	l "github.com/elitecodegroovy/goutil/logger"
	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/etcd"
	"github.com/micro/go-plugins/config/source/grpc"
	"go.uber.org/zap"
)

var (
	log     = l.GetLogger()
	appName = "auth_srv"
	cfg     = &authCfg{}
)

type authCfg struct {
	common.AppCfg
}

var appService micro.Service

var cancel context.CancelFunc
var ctx context.Context

func startBootstrap() {

	// 获取上下文，并声明关闭函数cancel
	ctx, cancel = context.WithCancel(context.Background())

	// 使用etcd注册
	micReg := etcd.NewRegistry(registryOptions)

	// 新建服务
	appService := micro.NewService(
		micro.Name(cfg.Name),
		micro.Name("mu.micro.book.srv.auth"),
		micro.Registry(micReg),
		micro.Version(cfg.Version),
		micro.Address(cfg.Addr()),
		micro.Context(ctx),
	)

	// 服务初始化
	appService.Init(
		micro.Action(func(c *cli.Context) {
			// 初始化handler
			model.Init()
			// 初始化handler
			handler.Init()
		}),
	)
	// 注册服务
	s.RegisterServiceHandler(appService.Server(), new(handler.Service))

	// 启动服务
	if err := appService.Run(); err != nil {
		log.Error("[main] error")
		panic(err)
	}
}

func registryOptions(ops *registry.Options) {
	etcdCfg := &common.Etcd{}
	err := config.C().App("etcd", etcdCfg)
	if err != nil {
		panic(err)
	}
	ops.Addrs = []string{fmt.Sprintf("%s:%d", etcdCfg.Host, etcdCfg.Port)}
}

func initCfg() {
	log.Info("[initCfg]...from main")
	source := grpc.NewSource(
		grpc.WithAddress("127.0.0.1:9600"),
		grpc.WithPath("micro"),
	)

	basic.Init(
		config.WithSource(source),
		config.WithApp(appName),
	)

	err := config.C().App(appName, cfg)
	if err != nil {
		panic(err)
	}

	log.Info("[initCfg] 配置", zap.Any("cfg", cfg))

	return
}

func listenChangedConf() {
	go func() {
		for {
			switch {
			case <-config.GetChangedChan():
				cancel()
				startBootstrap()
			}
		}

	}()
}

func main() {
	// 初始化配置、数据库等信息
	initCfg()
	// conf change listen
	listenChangedConf()
	//start service
	startBootstrap()

}
