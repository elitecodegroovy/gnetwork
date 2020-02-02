package zap

import (
	"github.com/micro/go-micro/util/log"
	"go.uber.org/zap"
)

// Options 配置项
type Options struct {
	zap.Config
	LogFileDir    string `json:logFileDir`
	AppName       string `json:"appName"`
	ErrorFileName string `json:"errorFileName"`
	WarnFileName  string `json:"warnFileName"`
	InfoFileName  string `json:"infoFileName"`
	DebugFileName string `json:"debugFileName"`
	MaxSize       int    `json:"maxSize"` // megabytes
	MaxBackups    int    `json:"maxBackups"`
	MaxAge        int    `json:"maxAge"` // days
}

var (
	_ = f0()
	a = c + b + z // == 9
	b = f()       // == 4
	c = f()       // == 5
	d = 3         // == 5 after initialization has finished
	z = f()       // == 5

)

func f0() int {
	log.Logf("...f0()")
	return d
}

func f() int {
	log.Logf("...f()")
	d++
	return d
}

func init() {
	log.Logf("options init ... %d, %d, %d, %d, %d", a, b, c, d, z)
}
