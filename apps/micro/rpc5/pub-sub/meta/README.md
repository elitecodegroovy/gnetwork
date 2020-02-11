
## HTTP TO RPC Handler
使用metadata模式下的Micro API，以下简称API。go-micro支持将请求路由到服务元数据声明的方法，也即是基于元数据的服务发现。

## Opts
使用protoc生成go代码

```
protoc --go_out=. --micro_out=. proto/api.proto
```

运行API网关，可以看到，API启动时，并没有声明handler模式，故而使用的RPC模式。所以Meta API其实是在RPC模式的基础上，通过在接口层声明端点元数据而指定服务的。

```
micro api --address=0.0.0.0:8080
```

运行示例程序，在代码中注册服务时，我们在endpoint参数中写入了元数据，声明接口为 /example和 /foo/bar

```
go run meta.go  
```

向 /example POST请求时，会被转到go.micro.api.example的Example.Call方法。 

```
curl -H 'Content-Type: application/json' -d '{"name": "john"}' "http://localhost:8080/example"
curl -XGET "http://localhost:8080/example?name=john"
```
向 /example POST请求时，会被转到go.micro.api.example的Foo.Bar方法。

```
curl -H 'Content-Type: application/json' -d '{}' http://localhost:8080/foo/bar
curl -XGET "http://localhost:8080/foo/bar"
curl -XDELETE "http://localhost:8080/foo/bar"
```