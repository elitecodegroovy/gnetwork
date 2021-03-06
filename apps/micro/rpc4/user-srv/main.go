package main

import (
	"fmt"
	"github.com/elitecodegroovy/gnetwork/apps/micro/rpc4/basic"
	"github.com/elitecodegroovy/gnetwork/apps/micro/rpc4/basic/common"
	"github.com/elitecodegroovy/gnetwork/apps/micro/rpc4/basic/config"
	"github.com/elitecodegroovy/gnetwork/apps/micro/rpc4/user-srv/handler"
	"github.com/elitecodegroovy/gnetwork/apps/micro/rpc4/user-srv/model"
	_ "github.com/elitecodegroovy/gnetwork/apps/micro/rpc4/user-srv/plugin"
	s "github.com/elitecodegroovy/gnetwork/apps/micro/rpc4/user-srv/proto/user"
	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/broker"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/etcd"
	"go.uber.org/zap"

	l "github.com/elitecodegroovy/goutil/logger"
	"github.com/micro/go-plugins/config/source/grpc"
)

var (
	log     = l.GetLogger()
	appName = "user_srv"
	cfg     = &userCfg{}
	topic   = "go.micro.web.topic.hi"
)

type userCfg struct {
	common.AppCfg
}

func main() {
	// 初始化配置、数据库等信息
	initCfg()

	// 使用etcd注册
	micReg := etcd.NewRegistry(registryOptions)

	bk := broker.NewBroker(
		broker.Addrs(fmt.Sprintf("%s:%d", "127.0.0.1", 11089)),
	)

	_, err := bk.Subscribe(topic, func(p broker.Event) error {
		log.Info(fmt.Sprintf("[sub] Received Body: %s, Header: %s", string(p.Message().Body), p.Message().Header))
		return nil
	})
	if err != nil {
		log.Info("[ERR] err: %s", zap.String("err", err.Error()))
	}

	// 新建服务
	service := micro.NewService(
		micro.Name("mu.micro.book.srv.user"),
		micro.Version("v1.0.0"),
		micro.Registry(micReg),
		micro.Broker(bk),
		micro.WrapHandler(handler.LogWrapper1, handler.LogWrapper2),
	)

	// 服务初始化
	service.Init(
		micro.Action(func(c *cli.Context) {
			// 初始化模型层
			model.Init()
			// 初始化handler
			handler.Init()
		}),
	)

	// 注册服务
	s.RegisterUserHandler(service.Server(), new(handler.Service))

	// 启动服务
	if err := service.Run(); err != nil {
		log.Fatal("", zap.String("error: ", err.Error()))
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

	basic.Init(config.WithSource(source))

	err := config.C().App(appName, cfg)
	if err != nil {
		panic(err)
	}

	log.Info("[initCfg] 配置，cfg：%v", zap.Any("cfg:", cfg))

	return
}
