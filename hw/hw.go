package hw

import (
	"time"

	"github.com/andreandradecosta/rpimonitor/models"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
)

//NewSample collects the current host data.
func NewSample() models.Sample {
	now := time.Now()
	s := models.Sample{
		LocalTime: now,
		Timestamp: now.Unix(),
		Metrics:   make(models.Info),
	}
	s.Metrics["cpu_times"] = getData(cpu.CPUTimes(false))
	s.Metrics["disk_io"] = getData(disk.DiskIOCounters())
	s.Metrics["load"] = getData(load.LoadAvg())
	s.Metrics["virtual_memory"] = getData(mem.VirtualMemory())
	s.Metrics["swap_memory"] = getData(mem.SwapMemory())
	s.Metrics["net_io"] = getData(net.NetIOCounters(true))
	s.Metrics["net_proto"] = getData(net.NetProtoCounters(nil))
	s.Metrics["temperature"] = getData(GetTemperature())
	return s
}

//NewStatus collects the current host status data.
func NewStatus() models.Info {
	i := make(models.Info)
	i["cpu_info"] = getData(cpu.CPUInfo())
	i["disk_usage"] = getData(disk.DiskUsage("/"))
	i["disk_part"] = getData(disk.DiskPartitions(true))
	i["host"] = getData(host.HostInfo())
	i["users"] = getData(host.Users())
	return i
}

func getData(data interface{}, err error) interface{} {
	if err != nil {
		return err.Error()
	}
	return data
}
