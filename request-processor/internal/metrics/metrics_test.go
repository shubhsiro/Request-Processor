package metrics

import (
	"fmt"
	"sync"
	"testing"
)

func TestLogUniqueRequest(t *testing.T) {
	uniqueRequests = sync.Map{}

	LogUniqueRequest("test-id-1")
	LogUniqueRequest("test-id-2")
	LogUniqueRequest("test-id-1") // Duplicate

	count := GetCurrentMinuteCount()
	if count != 2 {
		t.Errorf("Expected 2 unique requests, got %d", count)
	}
}

func TestGetCurrentMinuteCount(t *testing.T) {
	uniqueRequests = sync.Map{}

	for i := 0; i < 5; i++ {
		LogUniqueRequest(fmt.Sprintf("id-%d", i))
	}

	count := GetCurrentMinuteCount()
	if count != 5 {
		t.Errorf("Expected 5 unique requests, got %d", count)
	}
}

func TestConcurrentUniqueRequests(t *testing.T) {
	uniqueRequests = sync.Map{}

	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			LogUniqueRequest(fmt.Sprintf("concurrent-id-%d", id))
		}(i)
	}

	wg.Wait()

	count := GetCurrentMinuteCount()
	if count != 100 {
		t.Errorf("Expected 100 unique requests, got %d", count)
	}
}
