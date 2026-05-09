package main

import (
	"fmt"
	"log"

	"github.com/SaurabPoudel/smart-parking/aggregator/client"
)

const KafkaTopic = "anprData"

func main() {
	var (
		svc       ParkingSessionServicer
		err       error
		aggClient *client.Client
	)
	svc = NewParkingSessionService()
	svc = NewLogMiddleware(svc)

	aggClient = client.NewClient("http://localhost:3000/aggregate")

	kafkaConsumer, err := NewKafkaConsumer(KafkaTopic, svc, aggClient)
	if err != nil {
		log.Fatal(err)
	}
	kafkaConsumer.Start()
	fmt.Println("working fine")

}
