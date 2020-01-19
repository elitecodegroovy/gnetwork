
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
