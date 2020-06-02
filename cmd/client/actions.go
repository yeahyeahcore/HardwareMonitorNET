package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os/exec"
	"time"

	"github.com/rs/xid"

	"github.com/urfave/cli"
	"github.com/yeahyeahcore/HardwareMonitorNET/config"
	"github.com/yeahyeahcore/HardwareMonitorNET/storage"
)

func initialize(c *cli.Context) error {
	config.Server.Host = "Enter Server Host"
	config.Server.Port = "Enter Server Post"

	config.Client.ID = xid.New().String()
	config.Client.PauseSec = 30

	return config.Save("config.json")
}

func run(c *cli.Context) error {
	config.Load("config.json")

	path, err := exec.LookPath("InfoCheck.exe")
	if err != nil {
		fmt.Printf("Файл %s не найден", path)
	}

	fmt.Printf("Доступ к файлу %s\n", path)
	fmt.Println("Подключение... (если консоль горит, значит подключено успешно!)")

	cmd := exec.Command("InfoCheck.exe")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}

	if err := cmd.Start(); err != nil {
		return err
	}

	buf, err := ioutil.ReadAll(stdout)
	if err != nil {
		fmt.Println("readall err")
		return err
	}

	resp, err := http.Post(
		fmt.Sprintf("http://%s:%s/config_info", config.Server.Host, config.Server.Port),
		"application/json",
		bytes.NewBuffer(buf),
	)
	if err != nil {
		fmt.Println("Ошибка подключения")
		return err
	}

	if resp.StatusCode != 200 {
		fmt.Println("status err", resp.Status)
	}

	pause := func() {
		time.Sleep(time.Duration(config.Client.PauseSec) * time.Second)
	}

	for true {
		cmd := exec.Command("InfoCheck.exe")

		stdout, err := cmd.StdoutPipe()
		if err != nil {
			return err
		}
		if err := cmd.Start(); err != nil {
			return err
		}

		var params storage.Parameter
		err = json.NewDecoder(stdout).Decode(&params)
		if err != nil {
			fmt.Println("decode err")
			return err
		}

		params.DeviceID = config.Client.ID

		buf := bytes.NewBuffer(nil)
		err = json.NewEncoder(buf).Encode(params)

		if err != nil {
			fmt.Println("encode err")
			return err
		}

		resp, err := http.Post(
			fmt.Sprintf("http://%s:%s/post_info", config.Server.Host, config.Server.Port),
			"application/json",
			buf,
		)

		if err != nil {
			fmt.Println("Ошибка подключения")
			return err
		}

		if resp.StatusCode != 200 {
			fmt.Println("status err", resp.Status)
		}
		pause()
	}
	return nil
}
