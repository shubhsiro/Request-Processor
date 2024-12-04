package request

import (
	"fmt"
	"log"

	"github.com/IBM/sarama"
)

func SendToStreamingService(uniqueCount int) error {
	producer, err := sarama.NewSyncProducer([]string{"localhost:9092"}, nil)
	if err != nil {
		return fmt.Errorf("failed to create Kafka producer: %v", err)
	}
	defer producer.Close()

	message := fmt.Sprintf("Unique Count: %d", uniqueCount)
	kafkaMessage := &sarama.ProducerMessage{
		Topic: "unique-counts",
		Value: sarama.StringEncoder(message),
	}

	_, _, err = producer.SendMessage(kafkaMessage)
	if err != nil {
		return fmt.Errorf("failed to send message to Kafka: %v", err)
	}

	log.Printf("Successfully sent unique count to streaming service: %d", uniqueCount)

	return nil
}

// SendRequest handles sending requests based on the extension type.
func SendRequest(endpoint string, extension string, uniqueCount int) error {
	switch extension {
	case "1": // Extension 1: POST request
		payload := map[string]interface{}{
			"unique_count": uniqueCount,
		}
		return SendPOST(endpoint, payload)
	case "2": // Extension 2: POST request for deduplication
		payload := map[string]interface{}{
			"unique_count": uniqueCount,
			"source":       "load_balancer",
		}
		return SendPOST(endpoint, payload)
	case "3": // Extension 3: Send unique count to a streaming service
		err := SendToStreamingService(uniqueCount)
		if err != nil {
			return fmt.Errorf("failed to send to streaming service: %v", err)
		}
		return nil
	default: // Default case: GET request
		return SendGET(endpoint, uniqueCount)
	}
}
