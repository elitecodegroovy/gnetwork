# User Service

This is the User service

Generated with

```
micro new github.com/elitecodegroovy/gnetwork/apps/micro/rpc/user-web --namespace=mu.micro.book --alias=user --type=web
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: mu.micro.book.web.user
- Type: web
- Alias: user

## Dependencies

Micro services depend on service discovery. The default is multicast DNS, a zeroconf system.

In the event you need a resilient multi-host setup we recommend etcd.

```
# install etcd
brew install etcd

# run etcd
etcd
```

Run the user-web and user-srv with following steps.

运行api

```
$ micro --registry=etcd --api_namespace=mu.micro.book.web  api --handler=web
```
运行user-srv

```
$ cd user-srv
$ go run main.go plugin.go 
```

运行user-web

```
$ cd user-web
$ go run main.go
```

运行auth
```
$ cd auth
$ go run main.go
```

请求登录
```
$ curl --request POST   --url http://127.0.0.1:8080/user/login   --header 'Content-Type: application/x-www-form-urlencoded'  --data 'userName=micro&pwd=123'
```
返回结果
```
{
    "data": {
        "id": 10001,
        "name": "micro"
    },
    "ref": 1555428438879936000,
    "success": false,
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1ODMwMjY1OTgsImp0aSI6IjEwMDAxIiwiaWF0IjoxNTgwNDM0NTk4LCJpc3MiOiJib29rLm1pY3JvLm11IiwibmJmIjoxNTgwNDM0NTk4LCJzdWIiOiIxMDAwMSJ9.coOFsjIBxagc1lnlNAAdygoAJF-8L12CPEL_ZyADwxU"
}
```
退出登录（需要将token换成实际的）

```
curl --request POST \
  --url http://127.0.0.1:8080/user/logout \
  --cookie 'remember-me-token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1ODMwMjY1OTgsImp0aSI6IjEwMDAxIiwiaWF0IjoxNTgwNDM0NTk4LCJpc3MiOiJib29rLm1pY3JvLm11IiwibmJmIjoxNTgwNDM0NTk4LCJzdWIiOiIxMDAwMSJ9.coOFsjIBxagc1lnlNAAdygoAJF-8L12CPEL_ZyADwxU'
```

返回：
```
{"ref":1580434849259172067,"success":true}
```