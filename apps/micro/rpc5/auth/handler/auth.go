package handler

import (
	"context"
	"go.uber.org/zap"
	"strconv"

	"github.com/elitecodegroovy/gnetwork/apps/micro/rpc5/auth/model/access"
	auth "github.com/elitecodegroovy/gnetwork/apps/micro/rpc5/auth/proto/auth"
	l "github.com/elitecodegroovy/goutil/logger"
)

var (
	log           = l.GetLogger()
	accessService access.Service
)

// Init 初始化handler
func Init() {
	var err error
	accessService, err = access.GetService()
	if err != nil {
		log.Fatal("[Init] 初始化Handler错误", zap.String("err:", err.Error()))
		return
	}
}

type Service struct{}

// MakeAccessToken 生成token
func (s *Service) MakeAccessToken(ctx context.Context, req *auth.Request, rsp *auth.Response) error {
	log.Info("[MakeAccessToken] 收到创建token请求")

	token, err := accessService.MakeAccessToken(&access.Subject{
		ID:   strconv.FormatInt(req.UserId, 10),
		Name: req.UserName,
	})
	if err != nil {
		rsp.Error = &auth.Error{
			Detail: err.Error(),
		}

		log.Error("[MakeAccessToken] token生成失败", zap.String("err:", err.Error()))
		return err
	}

	rsp.Token = token
	return nil
}

// DelUserAccessToken 清除用户token
func (s *Service) DelUserAccessToken(ctx context.Context, req *auth.Request, rsp *auth.Response) error {
	log.Info("[DelUserAccessToken] 清除用户token")
	err := accessService.DelUserAccessToken(req.Token)
	if err != nil {
		rsp.Error = &auth.Error{
			Detail: err.Error(),
		}

		log.Info("[DelUserAccessToken] 清除用户token失败，err：%s", zap.Any("err:", err))
		return err
	}

	return nil
}

// GetCachedAccessToken 获取缓存的token
func (s *Service) GetCachedAccessToken(ctx context.Context, req *auth.Request, rsp *auth.Response) error {
	log.Info("[GetCachedAccessToken] 获取缓存", zap.Int64("token:", req.UserId))
	token, err := accessService.GetCachedAccessToken(&access.Subject{
		ID: strconv.FormatInt(req.UserId, 10),
	})
	if err != nil {
		rsp.Error = &auth.Error{
			Detail: err.Error(),
		}

		log.Info("[GetCachedAccessToken] 获取缓存的token失败，err：%s", zap.String("err:", err.Error()))
		return err
	}

	rsp.Token = token
	return nil
}
