package handler

import (
	"context"
	"strconv"

	"github.com/elitecodegroovy/gnetwork/apps/micro/rpc2/auth/model/access"
	auth "github.com/elitecodegroovy/gnetwork/apps/micro/rpc2/auth/proto/auth"
	"github.com/micro/go-micro/util/log"
)

var (
	accessService access.Service
)

// Init 初始化handler
func Init() {
	var err error
	accessService, err = access.GetService()
	if err != nil {
		log.Fatal("[Init] 初始化Handler错误，%s", err)
		return
	}
}

type Service struct{}

func (s *Service) GetCachedAccessToken(context.Context, *auth.AuthRequest, *auth.AuthResponse) error {
	panic("implement me")
}

// MakeAccessToken 生成token
func (s *Service) MakeAccessToken(ctx context.Context, req *auth.AuthRequest, rsp *auth.AuthResponse) error {
	log.Log("[MakeAccessToken] 收到创建token请求")

	token, err := accessService.MakeAccessToken(&access.Subject{
		ID:   strconv.FormatInt(req.UserId, 10),
		Name: req.UserName,
	})
	if err != nil {
		rsp.Error = &auth.Error{
			Detail: err.Error(),
		}

		log.Logf("[MakeAccessToken] token生成失败，err：%s", err)
		return err
	}

	rsp.Token = token
	return nil
}

// DelUserAccessToken 清除用户token
func (s *Service) DelUserAccessToken(ctx context.Context, req *auth.AuthRequest, rsp *auth.AuthResponse) error {
	log.Log("[DelUserAccessToken] 清除用户token")
	err := accessService.DelUserAccessToken(req.Token)
	if err != nil {
		rsp.Error = &auth.Error{
			Detail: err.Error(),
		}

		log.Logf("[DelUserAccessToken] 清除用户token失败，err：%s", err)
		return err
	}

	return nil
}
