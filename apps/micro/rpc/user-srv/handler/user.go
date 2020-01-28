package handler

import (
	"context"
	"time"

	us "github.com/elitecodegroovy/gnetwork/apps/micro/rpc/user-srv/model/user"
	s "github.com/elitecodegroovy/gnetwork/apps/micro/rpc/user-srv/proto/user"
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
func (e *Service) QueryUserByName(ctx context.Context, req *s.UserRequest, rsp *s.UserResponse) error {
	t1 := time.Now()
	user, err := userService.QueryUserByName(req.UserName)
	if err != nil {
		rsp.Success = false
		rsp.Error = &s.Error{
			Code:   500,
			Detail: err.Error(),
		}

		return nil
	}

	rsp.User = user
	rsp.Success = true
	log.Logf("req :%+v, rsp: %+v, time elapsed %d ms", req, rsp, (time.Since(t1)).Milliseconds())
	return nil
}
