package main

import (
	"fmt"
	"log"
	"os"

	"github.com/octogo/log/pkg/config"
	"github.com/urfave/cli"
)

var (
	gensrcFlags = []cli.Flag{
		cli.StringFlag{
			Name:  "pkg, p",
			Usage: "Name of the Go package the source file will be used in. Defaults to 'main'",
		},
		cli.StringFlag{
			Name:  "name, n",
			Usage: "Name of the source file to create (without the .go suffix). Defaults to 'log'",
		},
		cli.BoolFlag{
			Name:  "stdout, o",
			Usage: "Print sample source file to STDOUT rather than writing it to disk",
		},
	}
	gensrcCmd = cli.Command{
		Name:   "gensrc",
		Usage:  "Creates a Go source file with sample code in your CWD",
		Flags:  gensrcFlags,
		Action: gensrcRun,
	}
)

func gensrcRun(c *cli.Context) error {
	if c.Bool("stdout") {
		fmt.Println(config.GetSampleSource(c.String("pkg")))
	}

	srcFile := c.String("name")
	if srcFile == "" {
		srcFile = "log"
	}
	srcFile += ".go"
	if _, err := os.Stat(srcFile); !os.IsNotExist(err) {
		log.Fatalf("aborting because file already exists: %s", srcFile)
	}
	f, err := os.OpenFile(srcFile, os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}
	config.WriteSampleSource(c.String("pkg"), f)
	log.Printf("created: %s\n", srcFile)
	return nil
}
