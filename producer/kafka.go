package main

import (
	"context"
	"encoding/json"
	"github.com/segmentio/kafka-go"
)

var writer *kafka.Writer

func InitKafka() {
	writer = &kafka.Writer{
		Addr:     kafka.TCP("localhost:9092"),
		Topic:    "delivery.location.v1",
		Balancer: &kafka.Hash{},
	}
}

func Publish(event LocationEvent) error {
	bytes, _ := json.Marshal(event)
	return writer.WriteMessages(context.Background(),
		kafka.Message{
			Key:   []byte(event.PartnerID),
			Value: bytes,
		},
	)
}
