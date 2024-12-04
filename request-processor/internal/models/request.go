package models

type RequestPayload struct {
	UniqueCount int    `json:"unique_count"`
	Source      string `json:"source,omitempty"`
}
