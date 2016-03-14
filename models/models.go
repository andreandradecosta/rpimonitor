package models

import "time"

//Info is a basic data type
type Info map[string]interface{}

//Sample holds data collected at some time.
type Sample struct {
	LocalTime time.Time `json:"local_time"`
	Timestamp int64     `json:"timestamp"`
	Metrics   Info      `json:"metrics"`
}
