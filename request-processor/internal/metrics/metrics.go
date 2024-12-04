package metrics

import (
	"strconv"
	"sync"
	"time"

	"request-processor/internal/logger"
	"request-processor/internal/storage"
)

var (
	uniqueRequests sync.Map
	mu             sync.Mutex
)

// LogUniqueRequest logs a unique request ID.
func LogUniqueRequest(id string) {
	mu.Lock()

	defer mu.Unlock()

	uniqueRequests.Store(id, struct{}{})
}

// GetCurrentMinuteCount returns the count of unique requests in the current minute.
func GetCurrentMinuteCount() int {
	count := 0
	uniqueRequests.Range(func(_, _ interface{}) bool {
		count++
		return true
	})

	return count
}

// StartMetricsLogger logs unique request counts every minute.
func StartMetricsLogger(redisClient *storage.RedisClient) {
	for {
		time.Sleep(1 * time.Minute)

		count := GetCurrentMinuteCount()
		logger.Log("Unique requests in the last minute: " + strconv.Itoa(count))

		mu.Lock()
		uniqueRequests = sync.Map{}
		mu.Unlock()
	}
}
