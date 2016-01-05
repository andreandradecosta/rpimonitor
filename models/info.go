package models

import (
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
)

type Info map[string]interface{}

func NewInfo() Info {
	i := make(Info)
	i["CPUInfo"], _ = cpu.CPUInfo()
	i["DiskUsage"], _ = disk.DiskUsage("/")
	i["DiskPart"], _ = disk.DiskPartitions(true)
	i["Host"], _ = host.HostInfo()
	i["Users"], _ = host.Users()
	return i
}
