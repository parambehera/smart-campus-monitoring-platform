package kafka

import (
	"context"
	"encoding/json"
	"fmt"

	"device-service/internal/config"

	kafkago "github.com/segmentio/kafka-go"
)

var Writer *kafkago.Writer

func InitProducer() {

	fmt.Println("Kafka Broker:", config.GetEnv("KAFKA_BROKER"))
	fmt.Println("Kafka Topic:", config.GetEnv("KAFKA_TOPIC"))

	Writer = &kafkago.Writer{
		Addr:     kafkago.TCP(config.GetEnv("KAFKA_BROKER")),
		Topic:    config.GetEnv("KAFKA_TOPIC"),
		Balancer: &kafkago.LeastBytes{},
	}

	fmt.Println("✅ Kafka Producer Connected")
}

func PublishEvent(event Event) error {

	data, err := json.Marshal(event)
	if err != nil {
		return err
	}

	err = Writer.WriteMessages(
		context.Background(),
		kafkago.Message{
			Value: data,
		},
	)

	return err
}

func CloseProducer() error {
	return Writer.Close()
}