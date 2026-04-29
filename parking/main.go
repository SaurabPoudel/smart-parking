package main

import (
	"fmt"
	"log"
)

const KafkaTopic = "anprData"

func main() {
	var (
		svc ParkingSessionServicer
		err error
	)
	svc = NewParkingSessionService()
	svc = NewLogMiddleware(svc)
	kafkaConsumer, err := NewkafkaConsumer(KafkaTopic, svc)
	if err != nil {
		log.Fatal(err)
	}
	kafkaConsumer.Start()
	fmt.Println("working fine")

}
