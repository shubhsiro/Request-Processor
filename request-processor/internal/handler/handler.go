package handler

import (
	"fmt"
	"net/http"

	"request-processor/internal/logger"
	"request-processor/internal/metrics"
	"request-processor/internal/request"
	"request-processor/internal/storage"
)

// AcceptHandler handles /api/verve/accept requests.
func AcceptHandler(redisClient *storage.RedisClient, kafkaProducer *storage.KafkaProducer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		endpoint := r.URL.Query().Get("endpoint")
		extension := r.URL.Query().Get("extension")

		if id == "" {
			http.Error(w, "Missing 'id' query parameter", http.StatusBadRequest)
			return
		}

		// Deduplicate ID
		unique, err := storage.IsUnique(redisClient, id)
		if err != nil {
			logger.Log("Failed to check ID uniqueness: " + err.Error())
			http.Error(w, "failed", http.StatusInternalServerError)
			return
		}

		if !unique {
			http.Error(w, "failed", http.StatusConflict)
			return
		}

		metrics.LogUniqueRequest(id)

		// Process extensions
		if endpoint != "" {
			count := metrics.GetCurrentMinuteCount()
			err := request.SendRequest(endpoint, extension, count)
			if err != nil {
				logger.Log("Failed to send request to endpoint: " + err.Error())
				http.Error(w, "failed", http.StatusInternalServerError)
				return
			}
		}

		// Handle Extension 3 - Send count to Kafka
		if extension == "3" {
			count := metrics.GetCurrentMinuteCount()
			message := fmt.Sprintf("Unique count in the last minute: %d", count)

			err := kafkaProducer.SendMessage(message)
			if err != nil {
				logger.Log("Failed to send unique count to Kafka: " + err.Error())
				http.Error(w, "failed", http.StatusInternalServerError)

				return
			}
		}

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("ok"))
	}
}
