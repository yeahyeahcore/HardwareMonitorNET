package main

import (
	"github.com/yeahyeahcore/HardwareMonitorNET/config"
	"github.com/yeahyeahcore/HardwareMonitorNET/server"
)

func main() {
	config.Load("config.json")
	server.Init()
	server.Start()
}
