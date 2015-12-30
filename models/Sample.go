package models

import (
	"time"

	"github.com/shirou/gopsutil/mem"
)

type Sample struct {
	LocalTime time.Time              `json:"local_time"`
	Timestamp int64                  `json:"timestamp"`
	Metrics   map[string]interface{} `json:"metrics"`
}

func NewSample() Sample {
	now := time.Now()
	s := Sample{
		LocalTime: now,
		Timestamp: now.Unix(),
		Metrics:   make(map[string]interface{}),
	}
	vmem, _ := mem.VirtualMemory()
	s.Metrics["Memory"] = vmem

	return s
}
