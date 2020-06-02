package main

import (
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.App{
		Name:   "System monitor",
		Action: run,

		Commands: []cli.Command{
			{
				Name:   "init",
				Action: initialize,
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
