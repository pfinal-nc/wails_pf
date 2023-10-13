package sys

import (
	"context"
	"github.com/dustin/go-humanize"
	"github.com/shirou/gopsutil/mem"
	"github.com/wailsapp/wails/v2/pkg/logger"
)

/**
 * @Author: PFinal南丞
 * @Author: lampxiezi@163.com
 * @Date: 2023/10/13
 * @Desc:
 * @Project: wails_demo
 */

type CPU struct {
	ctx context.Context
}

type Memory struct {
	ctx context.Context
	log logger.Logger
}

type MemoryInfo struct {
	Total       string  `json:"total"`
	Available   string  `json:"available"`
	Used        string  `json:"used"`
	UsedPercent float64 `json:"used_percent"`
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
	// fmt.Println(memInfo)
	return &MemoryInfo{
		Total:       humanize.Bytes(memInfo.Total),
		Available:   humanize.Bytes(memInfo.Available),
		UsedPercent: memInfo.UsedPercent,
	}
}
