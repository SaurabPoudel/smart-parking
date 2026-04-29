package main

import (
	"encoding/json"

	"github.com/SaurabPoudel/smart-parking/types"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/sirupsen/logrus"
)

type KafkaConsumer struct {
	consumer  *kafka.Consumer
	IsRunning bool
	service   ParkingSessionServicer
}

func NewkafkaConsumer(topic string, svc ParkingSessionServicer) (*KafkaConsumer, error) {

	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost",
		"group.id":          "myGroup",
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		return nil, err
	}

	c.SubscribeTopics([]string{topic}, nil)

	// c.Close()

	return &KafkaConsumer{
		consumer: c,
		service:  svc,
	}, nil
}

func (c *KafkaConsumer) Start() {
	logrus.Info("Kafka transport started")
	c.IsRunning = true
	c.ReadMessageLoop()
}

func (c *KafkaConsumer) ReadMessageLoop() {
	for c.IsRunning {
		msg, err := c.consumer.ReadMessage(-1)
		if err != nil {
			logrus.Errorf("kafka consume error %s", err)
			continue
		}

		var anprData types.ANPRData
		if err := json.Unmarshal(msg.Value, &anprData); err != nil {
			logrus.Errorf("failed to unmarshal message: %v", err)
			continue
		}

		if anprData.Event == types.EventEntry {
			if err := c.service.ProcessEntry(anprData); err != nil {
				logrus.Errorf("error processing entry: %v", err)
			}
		} else if anprData.Event == types.EventExit {
			if err := c.service.ProcessExit(anprData); err != nil {
				logrus.Errorf("error processing exit: %v", err)
			}
		}
	}

}
