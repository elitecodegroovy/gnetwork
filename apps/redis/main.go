package main

import (
	"log"
	"time"
)

type valueEx struct {
	Name  string
	Email string
}

func main() {
	//Use your actually ip address here
	redisCluterClient := initialize("47.244.118.5:6379,47.244.118.5:6380,47.244.118.5:6381")
	key1 := "sampleKey"
	value1 := &valueEx{Name: "someName", Email: "someemail@abc.com"}
	err := redisCluterClient.setKey(key1, value1, time.Minute*1)
	if err != nil {
		log.Fatalf("Error: %v", err.Error())
	}
	value2 := &valueEx{}
	err = redisCluterClient.getKey(key1, value2)
	if err != nil {
		log.Fatalf("Error: %v", err.Error())
	}
	log.Printf("Name: %s", value2.Name)
	log.Printf("Email: %s", value2.Email)
}
