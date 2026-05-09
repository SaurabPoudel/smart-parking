package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/SaurabPoudel/smart-parking/aggregator/client"
	"github.com/SaurabPoudel/smart-parking/types"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type KafkaConsumer struct {
	consumer       *kafka.Consumer
	IsRunning      bool
	parkingService ParkingSessionServicer
	aggClient      *client.Client
}

func NewKafkaConsumer(topic string, svc ParkingSessionServicer, aggClient *client.Client) (*KafkaConsumer, error) {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost",
		"group.id":          "myGroup",
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		fmt.Println("error in creating consumer!")
		return nil, err
	}

	c.SubscribeTopics([]string{topic}, nil)

	return &KafkaConsumer{
		consumer:       c,
		parkingService: svc,
		aggClient:      aggClient,
	}, nil
}

func (c *KafkaConsumer) Start() {
	logrus.Info(" Kafka Transport started")
	c.IsRunning = true
	c.ReadMessageLoop()
}

func (c *KafkaConsumer) ReadMessageLoop() {
	for c.IsRunning {
		msg, err := c.consumer.ReadMessage(-1)
		if err != nil {
			logrus.Errorf("kafka consume error: %v (%v)\n", err, msg)
			continue
		}

		var data types.ANPRData
		if err := json.Unmarshal(msg.Value, &data); err != nil {
			logrus.Errorf("JSON serialization error: %v", err)
			continue
		}

		if data.Event == types.EventEntry {
			if err := c.parkingService.ProcessEntry(data); err != nil {
				logrus.Errorf("Error processing entry for plate %s: %v", data.Plate, err)
				continue
			}
			logrus.Infof("Entry processed for plate: %s", data.Plate)
			continue
		}

		if data.Event == types.EventExit {
			if err := c.parkingService.ProcessExit(data); err != nil {
				logrus.Errorf("Error processing exit for plate %s: %v", data.Plate, err)
				continue
			}

			fee, err := c.parkingService.GetSessionFee(data.Plate)
			if err != nil {
				logrus.Errorf("Error getting session fee for plate %s: %v", data.Plate, err)
				continue
			}

			invoice := types.Invoice{
				InvoiceID:     uuid.New().String(),
				Plate:         data.Plate,
				TotalSessions: 1,
				TotalAmount:   fee,
				PeriodStart:   time.Now(),
				PeriodEnd:     time.Now(),
				Status:        "PENDING",
			}

			if err := c.aggClient.AggregateInvoice(invoice); err != nil {
				logrus.Errorf("Aggregate error for plate %s: %v", data.Plate, err)
				continue
			}
			logrus.Infof("Exit processed and invoice sent for plate: %s, fee: %.2f", data.Plate, fee)
		}
	}
}
