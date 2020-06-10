package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/mvrilo/app-poc/pkg/server"
	"github.com/urfave/cli/v2"

	"github.com/mvrilo/app-poc/core/health"
	"github.com/mvrilo/app-poc/core/home"
	"github.com/mvrilo/app-poc/core/store"
)

func main() {
	app := &cli.App{
		Name:  "app-poc",
		Usage: "app-poc cli",
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

	err := godotenv.Load()
	if err != nil {
		log.Println(err)
	}

	err = app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
