package jwt

import "github.com/elitecodegroovy/gnetwork/apps/micro/rpc4/basic"

// Jwt 配置 接口
type Jwt struct {
	SecretKey string `json:"secretKey"`
}

// init 初始化Redis
func init() {
	basic.Register(initJwt)
}

func initJwt() {

}
