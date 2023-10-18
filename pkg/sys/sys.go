package sys

import (
	"context"
	"fmt"
	"github.com/dustin/go-humanize"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	"github.com/wailsapp/wails/v2/pkg/logger"
	"time"
)

/**
 * @Author: PFinal南丞
 * @Author: lampxiezi@163.com
 * @Date: 2023/10/13
 * @Desc:
 * @Project: wails_demo
 */

type Host struct {
	ctx context.Context
}

type CPU struct {
	ctx context.Context
}

type Memory struct {
	ctx context.Context
	log logger.Logger
}

type Disk struct {
	ctx context.Context
}

type MemoryInfo struct {
	Total       string  `json:"total"`
	Available   string  `json:"available"`
	Used        string  `json:"used"`
	UsedPercent float64 `json:"used_percent"`
}

type CPUInfo struct {
	CpuPercent float64 `json:"cpu_percent"`
	CpuNumber  int     `json:"cpu_number"`
}

type HostInfo struct {
	Hostname string `json:"hostname"`
	Platform string `json:"platform"`
}

type DiskInfo struct {
	Device     string `json:"device"`
	MountPoint string `json:"mount_point"`
	Total      string `json:"total"`
	UsageRate  string `json:"usage_rate"`
}

func NewMemory() *Memory {
	return &Memory{}
}

func (m *Memory) MemInfo() *MemoryInfo {
	memInfo, err := mem.VirtualMemory()
	if err != nil {
		m.log.Error("unable to get memory stats: " + err.Error())
		return nil
	}
	return &MemoryInfo{
		Total:       humanize.Bytes(memInfo.Total),
		Available:   humanize.Bytes(memInfo.Available),
		Used:        humanize.Bytes(memInfo.Used),
		UsedPercent: memInfo.UsedPercent,
	}
}

func NewCpu() *CPU {
	return &CPU{}
}

func (c *CPU) CPUInfo() *CPUInfo {
	cpuPercent, _ := cpu.Percent(time.Second, true)
	// fmt.Printf("CPU使用率: %.3f%% \n", cpuPercent[0])
	cpuNumber, _ := cpu.Counts(true)
	//fmt.Printf("CPU核心数: %v \n", cpuNumber)
	// fmt.Println(cpu.Info())
	return &CPUInfo{
		CpuNumber:  cpuNumber,
		CpuPercent: cpuPercent[0],
	}
}

func NewHost() *Host {
	return &Host{}
}

func (h *Host) HostInfo() *HostInfo {
	hostInfo, err := host.Info()
	//fmt.Println(hostInfo)
	if err != nil {
		fmt.Println("get host info fail, error: ", err)
	}
	// fmt.Printf("hostname is: %v, os platform: %v \n", hostInfo.Hostname, hostInfo.Platform)
	return &HostInfo{
		Hostname: hostInfo.Hostname,
		Platform: hostInfo.Platform,
	}
}

func NewDisk() *Disk {
	return &Disk{}
}

func (d *Disk) DiskInfo() (diskInfo []*DiskInfo) {
	diskPart, err := disk.Partitions(false)
	if err != nil {
		fmt.Println(err)
	}

	// fmt.Println(diskPart)
	for _, dp := range diskPart {
		//fmt.Println(dp)
		diskUsed, _ := disk.Usage(dp.Mountpoint)
		//fmt.Printf("分区总大小: %d MB \n", diskUsed.Total/1024/1024)
		//fmt.Printf("分区使用率: %.3f %% \n", diskUsed.UsedPercent)
		//fmt.Printf("分区inode使用率: %.3f %% \n", diskUsed.InodesUsedPercent)
		info := &DiskInfo{
			Device:     dp.Device,
			MountPoint: dp.Mountpoint,
			Total:      fmt.Sprintf("%d MB", diskUsed.Total/1024/1024),
			UsageRate:  fmt.Sprintf("%.3f %% \n", diskUsed.UsedPercent),
		}
		diskInfo = append(diskInfo, info)
	}
	return diskInfo
}
