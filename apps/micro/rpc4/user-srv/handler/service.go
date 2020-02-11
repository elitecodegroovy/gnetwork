package handler

import (
	"context"
	us "github.com/elitecodegroovy/gnetwork/apps/micro/rpc4/user-srv/model/user"
	s "github.com/elitecodegroovy/gnetwork/apps/micro/rpc4/user-srv/proto/user"
	"github.com/micro/go-micro/server"
	"github.com/micro/go-micro/util/log"
)

type Service struct{}

var (
	userService us.Service
)

// Init 初始化handler
func Init() {

	var err error
	userService, err = us.GetService()
	if err != nil {
		log.Fatal("[Init] 初始化Handler错误")
		return
	}
}

// QueryUserByName 通过参数中的名字返回用户
func (e *Service) QueryUserByName(ctx context.Context, req *s.Request, rsp *s.Response) error {
	//time.Sleep(3 * time.Second)
	log.Logf("......v1.0.0")
	user, err := userService.QueryUserByName(req.UserName)
	if err != nil {
		rsp.Error = &s.Error{
			Code:   500,
			Detail: err.Error(),
		}

		return nil
	}

	rsp.User = user
	return nil
}

// logWrapper1 包装HandlerFunc类型的接口
func LogWrapper1(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, rsp interface{}) error {
		log.Logf("[logWrapper1] %s 收到请求", req.Endpoint())
		err := fn(ctx, req, rsp)
		return err
	}
}

// logWrapper2 包装HandlerFunc类型的接口
func LogWrapper2(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, rsp interface{}) error {
		log.Logf("[logWrapper2] %s 收到请求", req.Endpoint())
		err := fn(ctx, req, rsp)
		return err
	}
}
