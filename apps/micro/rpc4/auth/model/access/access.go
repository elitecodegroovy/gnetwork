package access

import (
	"fmt"
	"go.uber.org/zap"
	"sync"

	"github.com/elitecodegroovy/gnetwork/apps/micro/rpc4/basic/config"
	"github.com/elitecodegroovy/gnetwork/apps/micro/rpc4/plugins/jwt"
	"github.com/elitecodegroovy/gnetwork/apps/micro/rpc4/plugins/redis"
	z "github.com/elitecodegroovy/goutil/logger"
	r "github.com/go-redis/redis"
)

var (
	log = z.GetLogger()
	s   *service
	ca  *r.Client
	m   sync.RWMutex
	cfg = &jwt.Jwt{}
)

// service 服务
type service struct {
}

// Service 用户服务类
type Service interface {
	// MakeAccessToken 生成token
	MakeAccessToken(subject *Subject) (ret string, err error)

	// GetCachedAccessToken 获取缓存的token
	GetCachedAccessToken(subject *Subject) (ret string, err error)

	// DelUserAccessToken 清除用户token
	DelUserAccessToken(token string) (err error)
}

// GetService 获取服务类
func GetService() (Service, error) {
	if s == nil {
		return nil, fmt.Errorf("[GetService] GetService 未初始化")
	}
	return s, nil
}

// Init 初始化用户服务层
func Init() {
	m.Lock()
	defer m.Unlock()

	if s != nil {
		return
	}

	err := config.C().App("jwt", cfg)
	if err != nil {
		panic(err)
	}

	log.Info("[initCfg] 配置", zap.Any("cfg: ", cfg))

	ca = redis.Redis()

	s = &service{}
}
