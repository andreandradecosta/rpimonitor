package models

import (
	"time"

	"github.com/shirou/gopsutil/mem"
)

type Sample struct {
	LocalTime time.Time              `json:"local_time"`
	Metrics   map[string]interface{} `json:"metrics"`
}

func NewSample() Sample {
	s := Sample{
		LocalTime: time.Now(),
		Metrics:   make(map[string]interface{}),
	}
	vmem, _ := mem.VirtualMemory()
	s.Metrics["Memory"] = vmem

	return s
}
