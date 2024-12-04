package utils

import (
	"os"
	"time"
)

// GetEnv retrieves the value of an environment variable or returns a default value.
func GetEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return fallback
}

func GetCurrentMinute() string {
	return time.Now().Format("2006-01-02 15:04")
}
