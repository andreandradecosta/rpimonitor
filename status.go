package rpimonitor

import "time"

type Status struct {
	LocalTime time.Time `json:"localTime"`
	Metrics   Info      `json:"metrics"`
}

type StatusReader interface {
	ReadStatus() (*Status, error)
}
