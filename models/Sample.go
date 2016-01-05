package models

import (
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
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
	s.Metrics["CPUTimes"], _ = cpu.CPUTimes(false)
	s.Metrics["DiskIO"], _ = disk.DiskIOCounters()
	s.Metrics["Load"], _ = load.LoadAvg()
	s.Metrics["VirtualMemory"], _ = mem.VirtualMemory()
	s.Metrics["SwapMemory"], _ = mem.SwapMemory()
	s.Metrics["NetIO"], _ = net.NetIOCounters(true)
	s.Metrics["NetProto"], _ = net.NetProtoCounters(nil)
	return s
}
