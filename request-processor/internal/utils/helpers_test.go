package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetEnv(t *testing.T) {
	tests := []struct {
		key      string
		fallback string
		expected string
	}{
		{"TEST_KEY", "default_value", "default_value"},
	}

	for _, tt := range tests {
		t.Run(tt.key, func(t *testing.T) {
			result := GetEnv(tt.key, tt.fallback)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestGetCurrentMinute(t *testing.T) {
	currentMinute := GetCurrentMinute()
	assert.NotEmpty(t, currentMinute)
}
