package main

import (
	"github.com/urfave/cli/v2"
	"github.io/cbuschka/go-secret-tool/internal/cliproc"
	"log"
	"os"
)

func main() {

	app := &cli.App{
		Name:  "secret-tool",
		Usage: "access keyring secrets",
		Commands: []*cli.Command{
			{
				Name:    "list",
				Aliases: []string{"l"},
				Usage:   "list all secrets",
				Action: func(cCtx *cli.Context) error {
					return cliproc.ListSecrets()
				},
			},
			{
				Name:    "get",
				Aliases: []string{"g"},
				Usage:   "get a secret",
				Action: func(cCtx *cli.Context) error {
					return cliproc.GetSecret(cCtx.Args().First())
				},
			}}}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
