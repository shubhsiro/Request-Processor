package logger

import (
	"log"
	"os"
)

var appLogger *log.Logger

// InitLogger initializes the logger with a file output.
func InitLogger(logFilePath string) {
	file, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}

	appLogger = log.New(file, "", log.LstdFlags)
}

func Log(message string) {
	appLogger.Println(message)
}
