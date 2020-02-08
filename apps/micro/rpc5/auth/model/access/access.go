package access

import (
	"fmt"
	redis2 "github.com/elitecodegroovy/gnetwork/apps/micro/rpc5/auth/model/redis"
	"go.uber.org/zap"
	"sync"

	"github.com/elitecodegroovy/gnetwork/apps/micro/rpc5/basic/config"
	"github.com/elitecodegroovy/gnetwork/apps/micro/rpc5/plugins/jwt"
	l "github.com/elitecodegroovy/goutil/logger"
	r "github.com/go-redis/redis"
)

var (
	log = l.GetLogger()
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

	err := config.GetConfigurator().Path("jwt", cfg)
	if err != nil {
		panic(err)
	}

	log.Info("[initCfg] 配置", zap.Any("cfg: ", cfg))
	ca = redis2.Redis()
	s = &service{}
}
