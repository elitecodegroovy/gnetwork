
grpc :
```

protoc -I/home/app/protoc/include -I. \
       -I$GOPATH/src \
       -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
       --go_out=plugins=grpc:. \
       proto/grpc/payment.proto
  
  
protoc -I/home/app/protoc/include -I. \
       -I$GOPATH/src \
       -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
       --grpc-gateway_out=logtostderr=true:. \
        proto/grpc/payment.proto
```

srv running :
```cgo
go run payment_server.go --registry=mdns --server_address=localhost:9090
```

srv api:
```cgo
go run grpc
```

http client :

```cgo

curl -d '{"orderId": "1222222222222222222222222222222", "account": "2010lllllllalsfasdfasdfasdfasdf"}' http://localhost:8080/order/pay

```
result :
```cgo
{"code":100,"success":"OK! 2010lllllllalsfasdfasdfasdfasdf","msg":"success: 1222222222222222222222222222222"}
```