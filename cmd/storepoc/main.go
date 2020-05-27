package main

import (
	"log"
	"os"

	"github.com/mvrilo/storepoc"
	"github.com/mvrilo/storepoc/core/health"
	"github.com/mvrilo/storepoc/core/store"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "storepoc",
		Usage: "storepoc cli",
		Action: func(c *cli.Context) error {
			server, err := storepoc.New()
			if err != nil {
				return err
			}

			err = server.Load(
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
