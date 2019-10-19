package main

import (
	"github.com/octogo/log"
	"github.com/octogo/log/pkg/config"
	"github.com/octogo/log/pkg/level"
	octolog "github.com/octogo/log/pkg/log"
)

func main() {
	c := &config.Config{
		// set the log-format to the default debug-format
		DefaultFormat: octolog.DefaultDebugFormat,
		// set the name of the standard logger to 'main'
		LoggerName: "main",
		// define two custom log-levels
		Levels: []config.Level{
			config.Level{
				Name:  "CUSTOM1",
				Color: "magenta", // see: https://en.wikipedia.org/wiki/ANSI_escape_code
			},
			config.Level{
				Name:  "CUSTOM2",
				Color: "5;41", // custom ANSII escape, e.g.: blinking text on red background
			},
		},
		// configure the STDOUT output to log the above custom log-levels
		Outputs: []config.Output{
			config.Output{
				URL: "file:///dev/stdout",
				Wants: []string{
					"custom1",
					"custom2",
					"debug",
					"info",
					"notice",
				},
			},
		},
	}
	log.InitWithConfig(c)
	l := log.New("", nil)
	l.Debug("This is a DEBUG log...")
	l.Info("This is an INFO log...")
	l.Notice("This is a NOTICE log...")
	l.Warning("This is a WARNING log...")
	l.Error("This is an ERROR log...")

	custom1, err := level.Parse("custom1")
	if err != nil {
		log.Fatal(err)
	}
	custom2, err := level.Parse("custom2")
	if err != nil {
		log.Fatal(err)
	}
	l.Log(custom1, "This is a CUSTOM1 log...")
	l.Log(custom2, "This is a CUSTOM2 log...")
}
