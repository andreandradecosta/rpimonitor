package device

import (
	"time"

	"github.com/andreandradecosta/rpimonitor"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
)

//Device represents the machine that is monitored.
type Device struct {
}

//ReadSample extract various metrics in the current time.
func (d *Device) ReadSample() (*rpimonitor.Sample, error) {
	now := time.Now()
	s := &rpimonitor.Sample{
		LocalTime: now,
		Timestamp: now.Unix(),
		Metrics:   make(rpimonitor.Info),
	}
	s.Metrics["cpuTimes"] = getData(cpu.Times(false))
	s.Metrics["diskIO"] = getData(disk.IOCounters())
	s.Metrics["load"] = getData(load.Avg())
	s.Metrics["virtualMemory"] = getData(mem.VirtualMemory())
	s.Metrics["swapMemory"] = getData(mem.SwapMemory())
	s.Metrics["netIO"] = getData(net.IOCounters(true))
	s.Metrics["netProto"] = getData(net.ProtoCounters(nil))
	s.Metrics["temperature"] = getData(GetTemperature())
	return s, nil
}

//ReadStatus extract static info about this device.
func (d *Device) ReadStatus() (*rpimonitor.Status, error) {
	now := time.Now()
	s := &rpimonitor.Status{
		LocalTime: now,
		Metrics:   make(rpimonitor.Info),
	}
	s.Metrics["cpuInfo"] = getData(cpu.Info())
	s.Metrics["diskUsage"] = getData(disk.Usage("/"))
	s.Metrics["diskPart"] = getData(disk.Partitions(true))
	s.Metrics["host"] = getData(host.Info())
	s.Metrics["users"] = getData(host.Users())
	return s, nil
}

func getData(data interface{}, err error) interface{} {
	if err != nil {
		return err.Error()
	}
	return data
}
