package handler

import (
	"context"
	"net/http"

	auth "github.com/elitecodegroovy/gnetwork/apps/micro/rpc2/auth/proto/auth"
	"github.com/elitecodegroovy/gnetwork/apps/micro/rpc2/basic/common"
	"github.com/elitecodegroovy/gnetwork/apps/micro/rpc2/plugins/session"
	"github.com/micro/go-micro/util/log"
)

// AuthWrapper 认证wrapper
func AuthWrapper(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ck, _ := r.Cookie(common.RememberMeCookieName)
		// token不存在，则状态异常，无权限
		if ck == nil {
			http.Error(w, "非法请求", 400)
			return
		}

		sess := session.GetSession(w, r)
		log.Logf("1_____________%+v", sess)
		if sess.ID != "" {
			// 检测是否通过验证
			if sess.Values["valid"] != nil {
				h.ServeHTTP(w, r)
				return
			} else {
				log.Logf("2______%+v", sess)
				_userId := sess.Values["userId"]
				if _userId == nil {
					log.Logf("[AuthWrapper]，cookie userId doesn't exist! ")
					http.Error(w, "非法请求", 400)
					return
				}
				userId := _userId.(int64)
				if userId != 0 {
					rsp, err := authClient.GetCachedAccessToken(context.Background(), &auth.AuthRequest{
						UserId: userId,
					})
					if err != nil {
						log.Logf("[AuthWrapper]，err：%s", err)
						http.Error(w, "非法请求", 400)
						return
					}

					// token不一致
					if rsp.Token != ck.Value {
						log.Logf("[AuthWrapper]，token不一致")
						http.Error(w, "非法请求", 400)
						return
					}
				} else {
					log.Logf("[AuthWrapper]，session不合法，无用户id")
					http.Error(w, "非法请求", 400)
					return
				}
			}
		} else {
			http.Error(w, "非法请求", 400)
			return
		}

		h.ServeHTTP(w, r)
	})
}
