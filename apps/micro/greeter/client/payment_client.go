package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro"
	"os"
	"time"

	proto "github.com/elitecodegroovy/gnetwork/apps/micro/proto/payment"
)

func main() {
	// Create a new service
	service := micro.NewService(micro.Name("greeter.client"))
	// Initialise the client and parse command line flags
	service.Init()

	// Create new greeter client
	cli := proto.NewPaymentService("go.micro.paymentService", service.Client())

	// Call the greeter
	rsp, err := cli.Pay(context.Background(), &proto.PayReq{OrderId: "20100000000000100010", Account: "order2020"})
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	// Print response
	fmt.Println(fmt.Sprintf("%d, %s, %s", rsp.Code, rsp.Success, rsp.Msg))
	time.Sleep(1 * time.Second)
}
