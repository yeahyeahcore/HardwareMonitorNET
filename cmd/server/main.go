package main

import (
	"path/filepath"

	"github.com/yeahyeahcore/HardwareMonitorNET/config"
	"github.com/yeahyeahcore/HardwareMonitorNET/server"
)

func main() {
	config.Load(filepath.Join("..", "..", "conf", "config.json"))
	server.Init()
	server.Start()
}