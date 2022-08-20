package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Action: func(cCtx *cli.Context) error {
			fmt.Printf("Args: %v\n", cCtx.Args().Slice())
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
