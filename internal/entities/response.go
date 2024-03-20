package entities

import "time"

type Response struct {
	ServiceName    string    `json:"service_name"`
	ServiceVersion string    `json:"service_version"`
	Status         string    `json:"status"`
	Timestamp      time.Time `json:"timestamp"`
}
