package main

import (
	"context"
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"github.com/wailsapp/wails/v2/pkg/logger"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"math"
	"time"
	"wails_demo/pkg/sys"
)

// App struct
type App struct {
	log logger.Logger
	ctx context.Context
}

type CPUUsage struct {
	Average int `json:"avg"`
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called at application startup
func (a *App) startup(ctx context.Context) {
	// Perform your setup here
	a.ctx = ctx
	// 监控 CPU
	go func() {
		for {
			runtime.EventsEmit(a.ctx, "cpu_usage", a.GetCPUUsage())
			time.Sleep(1 * time.Second)
		}
	}()
	// 监控 内存
	go func() {
		for {
			runtime.EventsEmit(a.ctx, "mem_usage", a.GetMemUsage())
			time.Sleep(1 * time.Second)
		}
	}()
}

// domReady is called after front-end resources have been loaded
func (a App) domReady(ctx context.Context) {
	// Add your action here
}

// beforeClose is called when the application is about to quit,
// either by clicking the window close button or calling runtime.Quit.
// Returning true will cause the application to continue, false will continue shutdown as normal.
func (a *App) beforeClose(ctx context.Context) (prevent bool) {
	return false
}

// shutdown is called at application termination
func (a *App) shutdown(ctx context.Context) {
	// Perform your teardown here
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

// GetCPUUsage cpu 使用结果
func (a *App) GetCPUUsage() *CPUUsage {
	percent, err := cpu.Percent(1*time.Second, false)
	if err != nil {
		a.log.Error("unable to get cpu stats: " + err.Error())
		return nil
	}

	return &CPUUsage{
		Average: int(math.Round(percent[0])),
	}
}

func (a *App) GetMemUsage() *sys.MemoryInfo {
	mem := sys.NewMemory()
	return mem.MemInfo()
}
