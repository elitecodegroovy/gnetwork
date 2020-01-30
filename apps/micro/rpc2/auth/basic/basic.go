package basic

import (
	"github.com/elitecodegroovy/gnetwork/apps/micro/rpc/user-srv/basic/config"
	"github.com/elitecodegroovy/gnetwork/apps/micro/rpc/user-srv/basic/db"
)

func Init() {
	config.Init()
	db.Init()
}
