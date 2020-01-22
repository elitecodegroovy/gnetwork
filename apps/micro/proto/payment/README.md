
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

对于micro工具包，只需使用内置标志来设置ttl和间隔。
```

micro --register_ttl=30 --register_interval=15 api
```