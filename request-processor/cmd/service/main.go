package main

import (
	"log"
	"net/http"
	"time"

	"request-processor/internal/config"
	"request-processor/internal/handler"
	"request-processor/internal/logger"
	"request-processor/internal/metrics"
	"request-processor/internal/storage"
)

func main() {
	cfg := config.LoadConfig()
	logger.InitLogger(cfg.LogFilePath)

	redisClient := storage.NewRedisClient(cfg.RedisAddr)

	kafkaProducer := storage.NewKafkaProducer("localhost:9092", "unique-count-topic")

	go metrics.StartMetricsLogger(redisClient)

	mux := http.NewServeMux()
	mux.HandleFunc("/api/verve/accept", handler.AcceptHandler(redisClient, kafkaProducer))

	server := &http.Server{
		Addr:         cfg.ServerAddr,
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Printf("Starting server on %s", cfg.ServerAddr)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
