package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro"
	"time"

	proto "github.com/elitecodegroovy/gnetwork/apps/micro/proto/greeter"
)

func main() {
	// Create a new service
	service := micro.NewService(micro.Name("greeter.client"))
	// Initialise the client and parse command line flags
	service.Init()

	// Create new greeter client
	greeter := proto.NewGreeterService("greeter", service.Client())

	// Call the greeter
	rsp, err := greeter.Hello(context.TODO(), &proto.HelloRequest{Name: "John"})
	if err != nil {
		fmt.Println(err)
	}

	// Print response
	fmt.Println(fmt.Sprintf("%v", rsp.Greeting))
	time.Sleep(2 * time.Second)
}
