package main

import (
	"log"
	"os"

	"github.com/restechnica/anyreleaser/cmd"
)

func main() {
	var err error
	var app = cmd.NewApp()

	if err = app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
