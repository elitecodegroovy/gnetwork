package main

import (
	"fmt"

	"github.com/elitecodegroovy/gnetwork/apps/micro/rpc2/basic"
	"github.com/elitecodegroovy/gnetwork/apps/micro/rpc2/basic/config"
	"github.com/elitecodegroovy/gnetwork/apps/micro/rpc2/payment-srv/handler"
	"github.com/elitecodegroovy/gnetwork/apps/micro/rpc2/payment-srv/model"
	s "github.com/elitecodegroovy/gnetwork/apps/micro/rpc2/payment-srv/proto/payment"
	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/etcd"
	"github.com/micro/go-micro/util/log"
)

func main() {
	// 初始化配置、数据库等信息
	basic.Init()

	// 使用etcd注册
	micReg := etcd.NewRegistry(registryOptions)

	// 新建服务
	service := micro.NewService(
		micro.Name("mu.micro.book.srv.payment"),
		micro.Registry(micReg),
		micro.Version("latest"),
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
	s.RegisterPaymentHandler(service.Server(), new(handler.Service))

	// 启动服务
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

func registryOptions(ops *registry.Options) {
	etcdCfg := config.GetEtcdConfig()
	ops.Addrs = []string{fmt.Sprintf("%s:%d", etcdCfg.GetHost(), etcdCfg.GetPort())}
}
