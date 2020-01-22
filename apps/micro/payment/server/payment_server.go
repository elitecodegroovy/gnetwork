package main

import (
	"fmt"
	payment "github.com/elitecodegroovy/gnetwork/apps/micro/proto/payment"
	"github.com/micro/go-micro"
	"golang.org/x/net/context"
)

const (
	paymentService = "go.micro.paymentService"
)

type Payment struct{}

func (g *Payment) Pay(ctx context.Context, req *payment.PayReq, rsp *payment.PayResp) error {
	rsp.Msg = "success: " + req.GetOrderId()
	rsp.Code = 100
	rsp.Success = "OK! " + req.GetAccount()
	return nil
}

func main() {
	// Create a new service. Optionally include some options here.
	//TTL 30s, register internal 15s
	service := micro.NewService(
		micro.Name(paymentService),
		//micro.RegisterTTL(time.Second*30),
		//micro.RegisterInterval(time.Second*15),
	)

	// Init will parse the command line flags.
	service.Init()

	// Register handler
	payment.RegisterPaymentServiceHandler(service.Server(), new(Payment))

	// Run the server
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}

	//grpc running with the parameters '--registry=mdns --server_address=localhost:9090'
}
