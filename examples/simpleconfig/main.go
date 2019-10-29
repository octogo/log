package main

import (
	"github.com/octogo/log"
	"github.com/octogo/log/pkg/config"
	"github.com/octogo/log/pkg/level"
	octolog "github.com/octogo/log/pkg/log"
)

func main() {
	log.InitWithConfig(&config.Config{
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
		// optionally configure a dedicated logger that only logs the above two custom
		// log-levels.
		Loggers: []config.Logger{
			config.Logger{
				Name:  "my-custom-logger",
				Wants: []string{"custom1", "custom2"},
			},
		},
	})

	l1 := log.New("", nil)
	l2 := log.New("my-custom-logger", nil)
	l1.Debug("This is a DEBUG log...")
	l1.Info("This is an INFO log...")
	l1.Notice("This is a NOTICE log...")
	l1.Warning("This is a WARNING log...")
	l1.Error("This is an ERROR log...")

	custom1, err := level.Parse("custom1")
	if err != nil {
		log.Fatal(err)
	}
	custom2, err := level.Parse("custom2")
	if err != nil {
		log.Fatal(err)
	}
	l2.Log(custom1, "This is a CUSTOM1 log...")
	l2.Log(custom2, "This is a CUSTOM2 log...")
}
