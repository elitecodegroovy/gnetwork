package main

import (
	"fmt"
	"github.com/elitecodegroovy/gnetwork/apps/micro/rpc5/auth/handler"
	"github.com/elitecodegroovy/gnetwork/apps/micro/rpc5/auth/model"
	"github.com/elitecodegroovy/gnetwork/apps/micro/rpc5/auth/model/redis"
	s "github.com/elitecodegroovy/gnetwork/apps/micro/rpc5/auth/proto/auth"
	cfg "github.com/elitecodegroovy/gnetwork/apps/micro/rpc5/basic/config"
	l "github.com/elitecodegroovy/goutil/logger"
	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/etcd"
	"go.uber.org/zap"
)

var (
	log           = l.GetLogger()
	appName       = "auth"
	moduleName    = "auth_srv"
	authorizedCfg = &cfg.AppCfg{}
	// 根据需要改成可配置的app列表
	configurationFileNames = []string{"auth"}
)

func Init() {
	//loading configuration files
	initCfg()
	//init authorized application
	initAuth()

	//Redis init func
	redis.Init()

}

func registryOptions(ops *registry.Options) {
	etcdCfg := &cfg.Etcd{}
	err := cfg.GetConfigurator().Path("etcd", etcdCfg)
	if err != nil {
		panic(err)
	}
	ops.Addrs = []string{fmt.Sprintf("%s:%d", etcdCfg.Host, etcdCfg.Port)}
	log.Info("[registryOptions] 配置", zap.Any("Addrs:", ops.Addrs))
}

func initCfg() {
	log.Info("[initCfg]...from main")

	cfg.SetAppName(appName)
	// 加载每个应用的配置文件
	cfg.LoadConfigurationFile(configurationFileNames)
}

func initAuth() {
	if err := cfg.GetConfigurator().Path("auth_srv", authorizedCfg); err != nil {
		panic("can't get the value of the file " + appName + ".yml")
	}

	log.Info("[initCfg] 配置", zap.Any("cfg", authorizedCfg))

	return
}

func main() {
	// 初始化配置、数据库等信息
	Init()

	// 使用etcd注册
	micReg := etcd.NewRegistry(registryOptions)

	// 新建服务
	service := micro.NewService(
		micro.Name(authorizedCfg.Name),
		micro.Name("mu.micro.book.srv.auth"),
		micro.Registry(micReg),
		micro.Version(authorizedCfg.Version),
		micro.Address(authorizedCfg.Addr()),
	)

	// 服务初始化
	service.Init(
		micro.Action(func(c *cli.Context) {
			// 初始化handler
			model.Init()
			// 初始化handler
			handler.Init()
		}),
	)
	// 注册服务
	s.RegisterServiceHandler(service.Server(), new(handler.Service))

	// 启动服务
	if err := service.Run(); err != nil {
		log.Error("[main] error")
		panic(err)
	}
}
