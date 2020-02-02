package broadcast

import (
	"github.com/elitecodegroovy/gnetwork/apps/micro/conf-center/proto/config"
)

type Broadcast interface {
	Send(namespace *config.ConfigResponse) error
	Watch() Watcher
}

type Watcher interface {
	Next() (*config.ConfigResponse, error)
	Stop() error
}

var broadcast Broadcast

func Init(b Broadcast) {
	broadcast = b
}

func GetBroadcast() Broadcast {
	return broadcast
}
