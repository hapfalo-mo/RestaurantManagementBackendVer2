package config

import (
	"RestuarantBackend/models/events"
	"context"
	"encoding/json"
	"github.com/segmentio/kafka-go"
	"time"
)

type KafkaClient struct {
	writer *kafka.Writer
}

// Generate New Kafka Client
func GenerateNewKafkaClient(brokers []string, topic string) (*KafkaClient, error) {
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  brokers,
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	})
	return &KafkaClient{writer: w}, nil
}

// Publish Order Service

func (k *KafkaClient) PublistOrderService(ctx context.Context, evt events.OrderMessage) error {
	data, err := json.Marshal(evt)
	if err != nil {
		return err
	}
	msg := &kafka.Message{
		Key:   []byte(evt.OrderId),
		Value: data,
		Time:  time.Now(),
	}
	return k.writer.WriteMessages(ctx, *msg)
}

// Close Kafka
func (kc *KafkaClient) Close() {
	kc.writer.Close()
}
