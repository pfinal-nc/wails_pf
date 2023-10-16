package main

import (
	"fmt"
	"testing"
	"wails_demo/pkg/sys"
)

/**
 * @Author: PFinal南丞
 * @Author: lampxiezi@163.com
 * @Date: 2023/10/13
 * @Desc:
 * @Project: wails_demo
 */

func TestSysMemInfo(t *testing.T) {
	mem := sys.NewMemory()
	fmt.Println(mem.MemInfo())
}

func TestSysCPUInfo(t *testing.T) {
	cpu := sys.NewCpu()
	fmt.Println(cpu.CPUInfo())
}

func TestSysHostInfo(t *testing.T) {
	host := sys.NewHost()
	fmt.Println(host.HostInfo())
}

func TestSysDiskInfo(t *testing.T) {
	disk := sys.NewDisk()
	for _, disk := range disk.DiskInfo() {
		fmt.Println(disk.Device)
	}
}
