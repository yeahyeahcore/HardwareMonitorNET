package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
	"path/filepath"
	"strconv"
	"time"

	"github.com/yeahyeahcore/HardwareMonitorNET/config"
)

func main() {
	config.Load(filepath.Join("..", "..", "conf", "config.json"))
	timeslice, err := strconv.Atoi(config.Client.Time)
	if err != nil {
		fmt.Println("parse timeslice err")
		log.Fatal(err)
	}

	for true {
		cmd := exec.Command("InfoCheck.exe")
		stdout, err := cmd.StdoutPipe()
		if err != nil {
			fmt.Println("pipe err")
			log.Fatal(err)
		}

		if err := cmd.Start(); err != nil {
			fmt.Println("start err")
			log.Fatal(err)
		}

		buf, err := ioutil.ReadAll(stdout)
		if err != nil {
			fmt.Println("readall err")
			log.Fatal(err)
		}

		resp, err := http.Post(
			fmt.Sprintf("http://%s:%s/", config.Server.Host, config.Server.Port),
			"application/json",
			bytes.NewBuffer(buf),
		)
		if err != nil {
			fmt.Println("post err")
			log.Fatal(err)
		}
		if resp.StatusCode != 200 {
			fmt.Println("status err", resp.Status)
		}

		time.Sleep(time.Duration(timeslice) * time.Second)
	}
}
