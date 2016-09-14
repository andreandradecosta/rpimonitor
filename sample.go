package rpimonitor

import "time"

//Sample holds data collected at some time.
type Sample struct {
	LocalTime time.Time `json:"localTime"`
	Timestamp int64     `json:"timestamp"`
	Metrics   Info      `json:"metrics"`
}

type SampleReader interface {
	ReadSample() (*Sample, error)
}

type SampleFetcher interface {
	Query(start, end time.Time) ([]Sample, error)
}
