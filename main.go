package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
)

//Info hui
type Info struct {
	NameCPU        string   `json:"NameCPU,omitempty"`
	TemperatureCPU []string `json:"TemperatureCPU,omitempty"`
	NameGPU        string   `json:"NameGPU,omitempty"`
	TemperatureGPU string   `json:"TemperatureGPU,omitempty"`
	NameHDD        string   `json:"NameHDD,omitempty"`
	TemperatureHDD string   `json:"TemperatureHDD,omitempty"`
	LocalIPAddress string   `json:"LocalIpAddress,omitempty"`
	MACaddress     string   `json:"MACaddress,omitempty"`
}

func main() {
	cmd := exec.Command("InfoCheck.exe")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	var ob Info
	if err := json.NewDecoder(stdout).Decode(&ob); err != nil {
		log.Fatal(err)
	}
	if err := cmd.Wait(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(ob)

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
}
