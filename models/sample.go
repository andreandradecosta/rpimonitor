package models

import (
	"time"

	"github.com/andreandradecosta/rpimonitor/hw"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
)

type Sample struct {
	LocalTime time.Time `json:"local_time"`
	Timestamp int64     `json:"timestamp"`
	Metrics   Info      `json:"metrics"`
}

func NewSample() Sample {
	now := time.Now()
	s := Sample{
		LocalTime: now,
		Timestamp: now.Unix(),
		Metrics:   make(Info),
	}
	s.Metrics["cpu_times"] = getData(cpu.CPUTimes(false))
	s.Metrics["disk_io"] = getData(disk.DiskIOCounters())
	s.Metrics["load"] = getData(load.LoadAvg())
	s.Metrics["virtual_memory"] = getData(mem.VirtualMemory())
	s.Metrics["swap_memory"] = getData(mem.SwapMemory())
	s.Metrics["net_io"] = getData(net.NetIOCounters(true))
	s.Metrics["net_proto"] = getData(net.NetProtoCounters(nil))
	s.Metrics["temperature"] = getData(hw.GetTemperature())
	return s
}
