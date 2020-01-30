package basic

import (
	"github.com/elitecodegroovy/gnetwork/apps/micro/rpc2/basic/config"
	"github.com/elitecodegroovy/gnetwork/apps/micro/rpc2/basic/db"
	"github.com/elitecodegroovy/gnetwork/apps/micro/rpc2/basic/redis"
)

func Init() {
	config.Init()
	db.Init()
	redis.Init()
}
