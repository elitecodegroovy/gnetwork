package main

import (
	"fmt"
	proto "github.com/elitecodegroovy/gnetwork/apps/micro/proto/greeter"
	"github.com/micro/go-micro"

	"golang.org/x/net/context"
)

const (
	paymentService = "go.micro.paymentService"
)

type Payment struct{}

func (g *Payment) Pay(ctx context.Context, req *proto.PayReq, rsp *proto.PayResp) error {
	rsp.Msg = "success: " + req.GetOrderId()
	rsp.Code = 100
	rsp.Success = "OK"
	return nil
}

func main() {
	// Create a new service. Optionally include some options here.
	service := micro.NewService(
		micro.Name(paymentService),
	)

	// Init will parse the command line flags.
	service.Init()

	// Register handler
	proto.RegisterPaymentServiceHandler(service.Server(), new(Payment))

	// Run the server
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
