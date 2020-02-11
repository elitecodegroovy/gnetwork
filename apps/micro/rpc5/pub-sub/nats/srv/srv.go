package main

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/broker"
	"github.com/micro/go-micro/util/log"
)

func Handler(event broker.Event) error {
	log.Logf("[sub] 收到消息，请查收: %v, %v", event.Message().Header, string(event.Message().Body))
	return nil
}

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.pubsub"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	broker := service.Server().Options().Broker

	if err := broker.Init(); err != nil {
		log.Fatalf("Broker Init error: %v", err)
	}
	if err := broker.Connect(); err != nil {
		log.Fatalf("Broker Connect error: %v", err)
	}

	if _, err := broker.Subscribe("go.micro.pubsub.topic.event", Handler); err != nil {
		log.Fatalf("broker.Subscribe error: %v", err)
	}

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

//go run srv.go --broker=nats --broker_address=127.0.0.1:4222
