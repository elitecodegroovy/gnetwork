如果使用默认的http broker，请运行：

```
go run main.go
```
如果想使用其他消息队列服务，例如nats，请运行：
```
export MICRO_BROKER=nats
go run main.go
```
OR 
```
go run main.go --broker=nats --broker_address=127.0.0.1:4222
```