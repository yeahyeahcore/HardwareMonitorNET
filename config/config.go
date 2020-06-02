package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

var cfg config

var (
	Server = &cfg.Server
	Client = &cfg.Client
)

func Load(path string) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(data, &cfg)
	if err != nil {
		log.Fatal(err)
	}
}

func Save(path string) error {
	file, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0775)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(cfg)
}
