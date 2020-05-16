package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
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
