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
