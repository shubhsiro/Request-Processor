package request

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSendGET(t *testing.T) {
	testCases := []struct {
		name        string
		url         string
		count       int
		expectError bool
	}{
		{
			name: "Successful GET",
			url: httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.Method != http.MethodGet {
					t.Errorf("Expected GET method, got %s", r.Method)
				}
				if r.URL.Query().Get("unique_count") != "5" {
					t.Errorf("Expected unique_count=5, got %s", r.URL.Query().Get("unique_count"))
				}
				w.WriteHeader(http.StatusOK)
			})).URL,
			count:       5,
			expectError: false,
		},
		{
			name:        "Invalid URL",
			url:         "invalid://url",
			count:       5,
			expectError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := SendGET(tc.url, tc.count)
			if (err != nil) != tc.expectError {
				t.Errorf("Expected error: %v, got: %v", tc.expectError, err)
			}
		})
	}
}

func TestSendPOST(t *testing.T) {
	testCases := []struct {
		name        string
		url         string
		payload     map[string]interface{}
		expectError bool
	}{
		{
			name: "Successful POST",
			url: httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.Method != http.MethodPost {
					t.Errorf("Expected POST method, got %s", r.Method)
				}

				var payload map[string]interface{}
				err := json.NewDecoder(r.Body).Decode(&payload)
				if err != nil {
					t.Errorf("Failed to decode payload: %v", err)
				}

				if payload["key"] != "value" {
					t.Errorf("Unexpected payload: %v", payload)
				}

				w.WriteHeader(http.StatusOK)
			})).URL,
			payload: map[string]interface{}{
				"key": "value",
			},
			expectError: false,
		},
		{
			name:        "JSON Marshaling Error",
			url:         httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})).URL,
			payload:     map[string]interface{}{"invalid": make(chan int)},
			expectError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := SendPOST(tc.url, tc.payload)
			if (err != nil) != tc.expectError {
				t.Errorf("Expected error: %v, got: %v", tc.expectError, err)
			}
		})
	}
}
