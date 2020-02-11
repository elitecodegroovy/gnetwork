package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/micro/go-micro/broker"
	"github.com/micro/go-micro/client/selector"
	"github.com/micro/go-micro/registry"
	"net/http"
	"time"

	auth "github.com/elitecodegroovy/gnetwork/apps/micro/rpc4/auth/proto/auth"
	"github.com/elitecodegroovy/gnetwork/apps/micro/rpc4/plugins/session"
	us "github.com/elitecodegroovy/gnetwork/apps/micro/rpc4/user-srv/proto/user"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/util/log"
)

var (
	serviceClient us.UserService
	authClient    auth.Service
	topic         = "go.micro.web.topic.hi"
)

// Error 错误结构体
type Error struct {
	Code   string `json:"code"`
	Detail string `json:"detail"`
}

// 过滤器，指定版本
// 也可改成传入数字，低于或大于某版本号的不调用，可灵活配置
func Filter(v string) client.CallOption {
	filter := func(services []*registry.Service) []*registry.Service {
		var filtered []*registry.Service

		for _, service := range services {
			if service.Version == v {
				filtered = append(filtered, service)
			}
		}

		return filtered
	}

	return client.WithSelectOption(selector.WithFilter(filter))
}

func Init() {
	//cli := client.NewClient(
	//	client.RequestTimeout(time.Second * 3),
	//)
	//
	//// 定义服务，可以传入其它可选参数
	//commonService := micro.NewService(
	//	micro.Name("timeout.client"),
	//	micro.Client(cli))
	// service.Init()

	serviceClient = us.NewUserService("mu.micro.book.srv.user", client.DefaultClient)
	authClient = auth.NewService("mu.micro.book.srv.auth", client.DefaultClient)
}

// Login 登录入口
func Login(w http.ResponseWriter, r *http.Request) {
	log.Logf("req: %+v ", r)
	// 只接受POST请求
	if r.Method != "POST" {
		log.Logf("非法请求")
		http.Error(w, "非法请求", 400)
		return
	}

	r.ParseForm()

	// 调用后台服务
	rsp, err := serviceClient.QueryUserByName(context.TODO(), &us.Request{
		UserName: r.Form.Get("userName"),
	}, Filter("v1.0.0"))
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// 返回结果
	response := map[string]interface{}{
		"ref": time.Now().UnixNano(),
	}

	if rsp.User.Pwd == r.Form.Get("pwd") {
		response["success"] = true

		// 干掉密码返回
		rsp.User.Pwd = ""
		response["data"] = rsp.User
		log.Logf("[Login] 密码校验完成，生成token...")

		// 生成token
		rsp2, err := authClient.MakeAccessToken(context.TODO(), &auth.Request{
			UserId:   rsp.User.Id,
			UserName: rsp.User.Name,
		})
		if err != nil {
			log.Logf("[Login] 创建token失败，err：%s", err)
			http.Error(w, err.Error(), 500)
			return
		}

		log.Logf("[Login] token %s", rsp2.Token)
		response["token"] = rsp2.Token

		// 同时将token写到cookies中
		w.Header().Add("set-cookie", "application/json; charset=utf-8")
		// 过期30分钟
		expire := time.Now().Add(30 * time.Minute)
		cookie := http.Cookie{Name: "remember-me-token", Value: rsp2.Token, Path: "/", Expires: expire, MaxAge: 90000}
		http.SetCookie(w, &cookie)

		// 同步到session中
		sess := session.GetSession(w, r)
		sess.Values["userId"] = rsp.User.Id
		sess.Values["userName"] = rsp.User.Name
		_ = sess.Save(r, w)
	} else {
		response["success"] = false
		response["error"] = &Error{
			Detail: "密码错误",
		}
	}

	w.Header().Add("Content-Type", "application/json; charset=utf-8")

	// 返回JSON结构
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

// Logout 退出登录
func Logout(w http.ResponseWriter, r *http.Request) {
	// 只接受POST请求
	if r.Method != "POST" {
		log.Logf("非法请求")
		http.Error(w, "非法请求", 400)
		return
	}

	tokenCookie, err := r.Cookie("remember-me-token")
	if err != nil {
		log.Logf("token获取失败")
		http.Error(w, "非法请求", 400)
		return
	}

	// 删除token
	_, err = authClient.DelUserAccessToken(context.TODO(), &auth.Request{
		Token: tokenCookie.Value,
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// 清除cookie
	cookie := http.Cookie{Name: "remember-me-token", Value: "", Path: "/", Expires: time.Now().Add(0 * time.Second), MaxAge: 0}
	http.SetCookie(w, &cookie)

	w.Header().Add("Content-Type", "application/json; charset=utf-8")

	// 返回结果
	response := map[string]interface{}{
		"ref":     time.Now().UnixNano(),
		"success": true,
	}

	// 返回JSON结构
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

func TestSession(w http.ResponseWriter, r *http.Request) {
	sess := session.GetSession(w, r)

	if v, ok := sess.Values["path"]; !ok {
		sess.Values["path"] = r.URL.Query().Get("path")
		log.Logf("path:" + r.URL.Query().Get("path"))
	} else {
		log.Logf(v.(string))
	}

	log.Logf(sess.ID)
	log.Logf(sess.Name())

	w.Write([]byte("OK"))
}

//It doesn't work.
func PushMsg(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json; charset=utf-8")

	_ = r.ParseForm()
	// 返回结果
	response := map[string]interface{}{
		"ref":  time.Now().UnixNano(),
		"data": "Hello! " + r.Form.Get("name"),
	}

	// 返回JSON结构
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	pub(r.Form.Get("name"))
}

func pub(name string) {
	msg := &broker.Message{
		Header: map[string]string{
			"name": fmt.Sprintf("%s", name),
		},
		Body: []byte(fmt.Sprintf("%s: %s", name, time.Now().String())),
	}
	if err := broker.Publish(topic, msg); err != nil {
		log.Logf("[pub] 发布消息失败： %s", err)
	} else {
		log.Logf("[pub] 发布消息：%s", string(msg.Body))
	}
}
