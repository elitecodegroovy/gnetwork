package redis

import (
	cfg "github.com/elitecodegroovy/gnetwork/apps/micro/rpc5/basic/config"
	z "github.com/elitecodegroovy/goutil/logger"
	r "github.com/go-redis/redis"
	logger "github.com/micro/go-micro/util/log"
	"go.uber.org/zap"
	"strings"
	"sync"
)

var (
	log    = z.GetLogger()
	client *r.Client
	m      sync.RWMutex
	inited bool
)

// redis redis 配置
type redis struct {
	Enabled  bool           `json:"enabled"`
	Conn     string         `json:"conn"`
	Password string         `json:"password"`
	DBNum    int            `json:"dbNum"`
	Timeout  int            `json:"timeout"`
	Sentinel *RedisSentinel `json:"sentinel"`
}

type RedisSentinel struct {
	Enabled bool   `json:"enabled"`
	Master  string `json:"master"`
	XNodes  string `json:"nodes"`
	nodes   []string
}

// Nodes redis 哨兵节点列表
func (s *RedisSentinel) GetNodes() []string {
	if len(s.XNodes) != 0 {
		for _, v := range strings.Split(s.XNodes, ",") {
			v = strings.TrimSpace(v)
			s.nodes = append(s.nodes, v)
		}
	}
	return s.nodes
}

// init 初始化Redis
func Init() {
	logger.Logf("init redis...")
	initRedis()
}

func initRedis() {
	m.Lock()
	defer m.Unlock()

	if inited {
		log.Info("[initRedis] 已经初始化过Redis...")
		return
	}

	log.Info("[initRedis] 初始化Redis...")

	c := cfg.GetConfigurator()
	redisConf := &redis{}
	err := c.Path("redis", redisConf)
	if err != nil {
		log.Info("[initRedis] ", zap.Any("error: ", err))
	}

	if !redisConf.Enabled {
		log.Info("[initRedis] 未启用redis")
		return
	}

	// 加载哨兵模式
	if redisConf.Sentinel != nil && redisConf.Sentinel.Enabled {
		log.Info("[initRedis] 初始化Redis，哨兵模式...")
		initSentinel(redisConf)
	} else { // 普通模式
		log.Info("[initRedis] 初始化Redis，普通模式...")
		initSingle(redisConf)
	}

	log.Info("[initRedis] 初始化Redis，检测连接...")

	pong, err := client.Ping().Result()
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Info("[initRedis]", zap.String(" 初始化Redis，检测连接Ping...", pong))
}

// Redis 获取redis
func Redis() *r.Client {
	return client
}

func initSentinel(redisConfig *redis) {
	client = r.NewFailoverClient(&r.FailoverOptions{
		MasterName:    redisConfig.Sentinel.Master,
		SentinelAddrs: redisConfig.Sentinel.GetNodes(),
		DB:            redisConfig.DBNum,
		Password:      redisConfig.Password,
	})

}

func initSingle(redisConfig *redis) {
	client = r.NewClient(&r.Options{
		Addr:     redisConfig.Conn,
		Password: redisConfig.Password, // no password set
		DB:       redisConfig.DBNum,    // use default DB
	})
}
