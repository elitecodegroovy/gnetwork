
Support registry etcd3 , run src like this:
```
go run payment_server.go –registrer=etcdv3 –registrer-address= http://192.168.1.147:2379
```


linux

```
# cd gnetowrk/apps/mciro
protoc --proto_path=$GOPATH/src:. --micro_out=. --go_out=. proto/payment/payment.proto
```

windows
```
# cd gnetowrk/apps/mciro
protoc --proto_path=. --go_out=. --micro_out=. proto/payment/payment.proto
```

grpc
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
