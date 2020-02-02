package cache

import "github.com/elitecodegroovy/gnetwork/apps/micro/conf-center/proto/config"

type Cache interface {
	Set(config *config.ConfigResponse) error
	Get(config *config.QueryConfigRequest) (v *config.ConfigResponse, ok bool)
	Clear()
}

func New(cacheSize int) Cache {
	return newFreeCache(cacheSize)
}
