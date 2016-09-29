package rpimonitor

import "time"

// Status holds static data about the device like CPU info, users, etc.
type Status struct {
	LocalTime time.Time `json:"localTime" bson:"localTime"`
	Metrics   Info      `json:"metrics" bson:"metrics"`
}

// StatusReader is the interface the must be implemented by the device.
type StatusReader interface {
	ReadStatus() (*Status, error)
}
