package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
	"strconv"
	"time"

	"github.com/yeahyeahcore/HardwareMonitorNET/config"
)

func main() {

	config.Load("config.json")

	timeslice, err := strconv.Atoi(config.Client.Time)
	if err != nil {
		fmt.Println("parse timeslice err")
		log.Fatal(err)
	}

	path, err := exec.LookPath("InfoCheck.exe")
	if err != nil {
		fmt.Printf("Файл %s не найден", path)
	}

	fmt.Printf("Доступ к файлу %s\n", path)
	fmt.Println("Подключение... (если консоль горит, значит подключено успешно!)")

	for true {
		cmd := exec.Command("InfoCheck.exe")

		stdout, err := cmd.StdoutPipe()
		if err != nil {
			log.Fatal(err)
		}
		if err := cmd.Start(); err != nil {
			log.Fatal(err)
		}

		buf, err := ioutil.ReadAll(stdout)
		if err != nil {
			fmt.Println("readall err")
			log.Fatal(err)
		}

		resp, err := http.Post(
			fmt.Sprintf("http://%s:%s/post_info", config.Server.Host, config.Server.Port),
			"application/json",
			bytes.NewBuffer(buf),
		)
		if err != nil {
			fmt.Println("Ошибка подключения")
			time.Sleep(time.Duration(timeslice) * time.Second)
			log.Fatal(err)
		}
		if resp.StatusCode != 200 {
			fmt.Println("status err", resp.Status)
		}

		time.Sleep(time.Duration(timeslice) * time.Second)
	}

}
