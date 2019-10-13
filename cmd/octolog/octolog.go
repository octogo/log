package main

import (
	"fmt"
	"log"
	"os"

	"github.com/octogo/log/pkg/config"
	"github.com/urfave/cli"
)

// Version defines the version of the build.
var Version = "0.1.0"

func main() {
	app := cli.NewApp()
	app.Name = "octolog"
	app.Usage = "Go logging for human beings"
	app.Version = Version

	app.Commands = []cli.Command{
		{
			Name:  "genconf",
			Usage: "Creates a sample configuration file in your CWD.",
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "stdout, o",
					Usage: "Write sample configuration to STDOUT.",
				},
			},
			Action: func(c *cli.Context) error {
				if c.Bool("stdout") {
					fmt.Println(config.GetSampleConfig(Version))
					return nil
				}
				configFile := "logging.yml"
				if _, err := os.Stat(configFile); !os.IsNotExist(err) {
					log.Fatalf("aboring because: %s already exists", configFile)
				}
				f, err := os.OpenFile(configFile, os.O_CREATE|os.O_WRONLY, 0600)
				if err != nil {
					log.Fatal(err)
				}
				config.WriteSampleToFile(Version, f)
				log.Printf("Created: %s\n", configFile)
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
