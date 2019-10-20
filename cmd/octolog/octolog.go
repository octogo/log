package main

import (
	"log"
	"os"

	"github.com/urfave/cli"
)

// Version defines the version of the build.
var Version = "0.2.0"

func main() {
	app := cli.NewApp()
	app.Name = "octolog"
	app.Usage = "Go logging for human beings"
	app.Version = Version

	app.Commands = []cli.Command{
		genconfCmd,
		gensrcCmd,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
