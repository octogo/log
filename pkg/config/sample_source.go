package config

import "os"

var sampleSource = `
import (
	"github.com/octogo/log"
	"github.com/octogo/log/pkg/config"
	octolog "github.com/octogo/log/pkg/log"
)

var octologConfig = &config.Config{
	DefaultFormat: octolog.DefaultLogFormat,
	LoggerName:    "main",
	Levels: []config.Level{
		config.Level{
			Name: "custom1",
			Color: "magenta",
		},
		config.Level{
			Name: "custom2",
			Color: "5;41",
		},
	},
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
	Loggers: []config.Logger{
		config.Logger{
			Name: "my-custom-logger",
			Wants: []string{"custom1", "custom2"},
		},
	},
}

func initOctolog() {
	log.InitWithConfig(octologConfig)
}
`

// GetSampleSource returns a sample source file.
func GetSampleSource(pkg string) string {
	if pkg == "" {
		pkg = "main"
	}
	out := "package " + pkg + "\n"
	out += sampleSource
	return out
}

// WriteSampleSource writes the sample source into the given file.
func WriteSampleSource(pkg string, f *os.File) {
	f.WriteString(GetSampleSource(pkg))
}
