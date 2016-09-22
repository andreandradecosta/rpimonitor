package rpimonitor

import "time"

//Sample holds data collected at some time.
type Sample struct {
	LocalTime time.Time `json:"localTime" bson:"localTime"`
	Timestamp int64     `json:"timestamp" bson:"timestamp"`
	Metrics   Info      `json:"metrics" bson:"metrics"`
}

type SampleReader interface {
	ReadSample() (*Sample, error)
}

type SampleFetcher interface {
	Query(start, end time.Time) ([]Sample, error)
}

type SampleWriter interface {
	Write(sample *Sample) error
}
