package models

import (
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
)

func NewStatus() Info {
	i := make(Info)
	i["cpu_info"] = getData(cpu.CPUInfo())
	i["disk_usage"] = getData(disk.DiskUsage("/"))
	i["disk_part"] = getData(disk.DiskPartitions(true))
	i["host"] = getData(host.HostInfo())
	i["users"] = getData(host.Users())
	return i
}
