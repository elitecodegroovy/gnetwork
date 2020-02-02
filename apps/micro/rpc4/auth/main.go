package main

import (
	"fmt"

	"github.com/elitecodegroovy/gnetwork/apps/micro/rpc4/auth/handler"
	"github.com/elitecodegroovy/gnetwork/apps/micro/rpc4/auth/model"
	_ "github.com/elitecodegroovy/gnetwork/apps/micro/rpc4/auth/plugin"
	s "github.com/elitecodegroovy/gnetwork/apps/micro/rpc4/auth/proto/auth"
	"github.com/elitecodegroovy/gnetwork/apps/micro/rpc4/basic"
	"github.com/elitecodegroovy/gnetwork/apps/micro/rpc4/basic/common"
	"github.com/elitecodegroovy/gnetwork/apps/micro/rpc4/basic/config"
	z "github.com/elitecodegroovy/gnetwork/apps/micro/rpc4/plugins/zap"
	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/etcd"
	"github.com/micro/go-plugins/config/source/grpc"
	"go.uber.org/zap"
)

var (
	log     = z.GetLogger()
	appName = "auth_srv"
	cfg     = &authCfg{}
)

type authCfg struct {
	common.AppCfg
}

func main() {
	// 初始化配置、数据库等信息
	initCfg()

	// 使用etcd注册
	micReg := etcd.NewRegistry(registryOptions)

	// 新建服务
	service := micro.NewService(
		micro.Name(cfg.Name),
		micro.Registry(micReg),
		micro.Version(cfg.Version),
		micro.Address(cfg.Addr()),
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

func registryOptions(ops *registry.Options) {
	etcdCfg := &common.Etcd{}
	err := config.C().App("etcd", etcdCfg)
	if err != nil {
		panic(err)
	}
	ops.Addrs = []string{fmt.Sprintf("%s:%d", etcdCfg.Host, etcdCfg.Port)}
}

func initCfg() {
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
	log.Info("[initCfg]", zap.String("appName", appName))
	log.Info("[initCfg] 配置", zap.Any("cfg", cfg))

	return
}