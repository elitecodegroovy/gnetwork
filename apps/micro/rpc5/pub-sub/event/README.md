## Event handler

micro API会把http请求映射到匹配的事件处理服务上.
## 运行

因为我们在代码中声明了事件主题topic是go.micro.evt.user，即是说事件服务的命名所属空间是go.micro.evt，所以我们的API也要是这个命名空间，这样API才能找到它。
```
micro api --handler=event --namespace=go.micro.evt
```
运行服务

```
go run main.go
```

发送事件

```
curl -d '{"message": "Hello, Micro中国"}' http://localhost:8080/user/login
```