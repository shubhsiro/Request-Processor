package config

import "os"

// Config holds configuration values for the application.
type Config struct {
	ServerAddr  string
	LogFilePath string
	RedisAddr   string
}

// LoadConfig loads configuration from environment variables or defaults.
func LoadConfig() Config {
	return Config{
		ServerAddr:  getEnv("SERVER_ADDR", ":8080"),
		LogFilePath: getEnv("LOG_FILE_PATH", "./logs/app.log"),
		RedisAddr:   getEnv("REDIS_ADDR", "localhost:6379"),
	}
}

// getEnv gets the value of an environment variable or returns a default value.
func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return fallback
}
