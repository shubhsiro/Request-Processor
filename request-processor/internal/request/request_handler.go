package request

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// SendGET sends a GET request to the specified endpoint.
func SendGET(endpoint string, count int) error {
	url := fmt.Sprintf("%s?unique_count=%d", endpoint, count)

	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}

// SendPOST sends a POST request to the specified endpoint with a payload.
func SendPOST(endpoint string, payload map[string]interface{}) error {
	body, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	resp, err := http.Post(endpoint, "application/json", bytes.NewReader(body))
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}
