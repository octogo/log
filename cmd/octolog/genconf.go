package main

import (
	"fmt"
	"log"
	"os"

	"github.com/octogo/log/pkg/config"
	"github.com/urfave/cli"
)

var (
	genconfFlags = []cli.Flag{
		cli.BoolFlag{
			Name:  "stdout, o",
			Usage: "Write sample configuration to STDOUT",
		},
	}
	genconfCmd = cli.Command{
		Name:   "genconf",
		Usage:  "Creates a configuration file in your CWD",
		Flags:  genconfFlags,
		Action: genconfRun,
	}
)

func genconfRun(c *cli.Context) error {
	if c.Bool("stdout") {
		fmt.Println(config.GetSampleConfig(Version))
		return nil
	}
	configFile := "logging.yml"
	if _, err := os.Stat(configFile); !os.IsNotExist(err) {
		return err
	}
	f, err := os.OpenFile(configFile, os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}
	config.WriteSampleToFile(Version, f)
	log.Printf("Created: %s\n", configFile)
	return nil
}
