package main

import (
	"log"
	"os"

	"github.com/mvrilo/storepoc/pkg/server"
	"github.com/urfave/cli/v2"

	"github.com/mvrilo/storepoc/core/health"
	"github.com/mvrilo/storepoc/core/home"
	"github.com/mvrilo/storepoc/core/store"
)

func main() {
	app := &cli.App{
		Name:  "storepoc",
		Usage: "storepoc cli",
		Action: func(c *cli.Context) error {
			server, err := server.New()
			if err != nil {
				return err
			}

			err = server.Load(
				&home.Home{},
				&health.Health{},
				&store.Store{},
			)
			if err != nil {
				return err
			}

			server.Start()
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
