package storage

import (
	"context"
	"log"

	"github.com/segmentio/kafka-go"
)

type KafkaProducer struct {
	writer *kafka.Writer
}

// NewKafkaProducer creates a new Kafka producer.
func NewKafkaProducer(broker string, topic string) *KafkaProducer {
	writer := &kafka.Writer{
		Addr:     kafka.TCP(broker),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}

	return &KafkaProducer{writer: writer}
}

// SendMessage sends a message to the Kafka topic.
func (p *KafkaProducer) SendMessage(message string) error {
	err := p.writer.WriteMessages(context.Background(),
		kafka.Message{
			Value: []byte(message),
		},
	)
	if err != nil {
		log.Println("Failed to send message to Kafka:", err)
		return err
	}

	return nil
}

// Close closes the Kafka writer connection.
func (p *KafkaProducer) Close() {
	if err := p.writer.Close(); err != nil {
		log.Println("Failed to close Kafka writer:", err)
	}
}
