
linux

```
# cd gnetowrk/apps/mciro
protoc --proto_path=$GOPATH/src:. --micro_out=. --go_out=. proto/greeter/greeter.proto
```

windows
```
# cd gnetowrk/apps/mciro
protoc --proto_path=. --go_out=. --micro_out=. proto/greeter/greeter.proto
```

grpc
```

protoc -I/home/app/protoc/include -I. \
       -I$GOPATH/src \
       -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
       --go_out=plugins=grpc:. \
       proto/greeter/hello.proto
  
  
protoc -I/home/app/protoc/include -I. \
       -I$GOPATH/src \
       -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
       --grpc-gateway_out=logtostderr=true:. \
        proto/greeter/hello.proto
```