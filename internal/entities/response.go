package entities

import "time"

type Response struct {
	ServiceName    string
	Status         int
	Timestamp      time.Time
	ServiceVersion string
}
