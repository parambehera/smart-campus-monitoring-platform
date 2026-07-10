package kafka

import "time"

type Event struct {
	EventType string      `json:"eventType"`
	Source    string      `json:"source"`
	Timestamp time.Time   `json:"timestamp"`
	Payload   interface{} `json:"payload"`
}