package log

import (
	"errors"
	"os"
	"strings"

	"github.com/octogo/log/internal/lib"
	"github.com/octogo/log/pkg/config"
	"github.com/octogo/log/pkg/level"
)

// Init initializes the package.
func Init() {
	NewFileOutput(
		os.Stdout,
		[]level.Level{
			level.INFO,
			level.NOTICE,
		},
		DefaultLogFormat,
	)
	NewFileOutput(
		os.Stderr,
		[]level.Level{
			level.WARNING,
			level.ERROR,
		},
		DefaultLogFormat,
	)
	for i := range DefaultOutputs {
		loadOutput(DefaultOutputs[i])
	}
	defaultLogger = NewLogger(LoggerName, nil)
}

// Configure configures this package according to the given configuration.
func Configure(c *config.Config) {
	DefaultLogFormat = c.DefaultFormat
	LoggerName = c.LoggerName
	if c.DefaultOutputs != nil && len(c.DefaultOutputs) > 0 {
		DefaultOutputs = c.DefaultOutputs
	}
	// call Init() after configuring defaults
	loadOutputs(c.Outputs...)
	// Todo: loadLoggers
	Init()
}

func loadOutput(url string) Output {
	schema, uri, err := lib.ParseURL(url)
	if err != nil {
		panic(err)
	}
	switch strings.ToLower(schema) {
	case "file":
		switch uri {
		case os.Stdout.Name():
			if output := GetOutput("file://" + os.Stdout.Name()); output == nil {
				return NewFileOutput(os.Stdout, nil, DefaultDebugFormat)
			} else {
				return output
			}
		case os.Stderr.Name():
			if output := GetOutput("file://" + os.Stderr.Name()); output == nil {
				return NewFileOutput(os.Stderr, nil, DefaultDebugFormat)
			} else {
				return output
			}
		default:
			f := lib.OpenFile(uri)
			if output := GetOutput(f.Name()); output == nil {
				return NewFileOutput(f, nil, DefaultDebugFormat)
			} else {
				return output
			}
		}
	default:
		panic(errors.New("unsupported schema in URL: " + schema))
	}
}

func loadOutputs(configuredOutputs ...config.Output) []Output {
	if configuredOutputs == nil || len(configuredOutputs) == 0 {
		return []Output{}
	}
	outputs := make([]Output, len(configuredOutputs))
	for i := range configuredOutputs {
		o := loadOutput(configuredOutputs[i].URL)
		o.SetWants(lib.ParseLevels(configuredOutputs[i].Wants...))

		var format string
		if configuredOutputs[i].Format == "" {
			format = DefaultLogFormat
		} else {
			format = configuredOutputs[i].Format
		}
		o.SetFormat(format)
		outputs[i] = o
	}
	return outputs
}
