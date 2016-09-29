package rpimonitor

import "time"

// Sample holds data collected at some time.
type Sample struct {
	LocalTime time.Time `json:"localTime" bson:"localTime"`
	Timestamp int64     `json:"timestamp" bson:"timestamp"`
	Metrics   Info      `json:"metrics" bson:"metrics"`
}

// SampleReader is the interface that must implemented by the device.
type SampleReader interface {
	ReadSample() (*Sample, error)
}

// SampleFetcher is the interface that must be implemented by a database service
// for querying data.
type SampleFetcher interface {
	Query(start, end time.Time) ([]Sample, error)
}

// SampleWriter is the interface that must be implemented by a database service
// for persisting data.
type SampleWriter interface {
	Write(sample *Sample) error
}
