package main

import (
	"github.com/elitecodegroovy/gnetwork/apps/micro/conf-center/agent-api/config"
	"github.com/elitecodegroovy/gnetwork/apps/micro/conf-center/agent-api/handler"
	pconfig "github.com/elitecodegroovy/gnetwork/apps/micro/conf-center/proto/config"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/web"
)

func main() {
	service := web.NewService(
		web.Name("go.micro.api.agent"),
	)

	if err := service.Init(); err != nil {
		panic(err)
	}

	client := pconfig.NewConfigService("go.micro.srv.config", service.Options().Service.Client())

	config.Init(client, 1024*1024)
	router := Router()
	service.Handle("/", router)

	if err := service.Run(); err != nil {
		panic(err)
	}
}

func Router() *gin.Engine {
	router := gin.Default()
	r := router.Group("/agent/api/v1")
	r.GET("/config", handler.ReadConfig)
	r.GET("/watch", handler.WatchUpdate)

	return router
}
