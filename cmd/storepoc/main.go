package main

import (
	"log"
	"os"

	"github.com/mvrilo/storepoc"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "storepoc",
		Usage: "storepoc cli",
		Action: func(c *cli.Context) error {
			poc, err := storepoc.New()
			if err != nil {
				return err
			}
			poc.Start()
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
