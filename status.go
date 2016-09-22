package rpimonitor

import "time"

type Status struct {
	LocalTime time.Time `json:"localTime" bson:"localTime"`
	Metrics   Info      `json:"metrics" bson:"metrics"`
}

type StatusReader interface {
	ReadStatus() (*Status, error)
}
