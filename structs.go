package YTElkProducer

import "time"

type LogDocument struct {
	Timestamp time.Time   `json:"timestamp"`
	Log       interface{} `json:"log"`
}
