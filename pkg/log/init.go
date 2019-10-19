package log

import (
	"errors"
	"os"
	"strings"

	"github.com/octogo/log/internal/lib"
	"github.com/octogo/log/pkg/color"
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
	loadLevels(c.Levels)
	loadOutputs(c.Outputs...)
	// Todo: loadLoggers
	Init()
}

func loadLevels(levels []config.Level) {
	if levels == nil || len(levels) == 0 {
		return
	}
	for i := range levels {
		switch strings.ToUpper(levels[i].Color) {
		case "BLACK":
			level.Register(levels[i].Name, color.New(color.NormalDisplay, color.Black))
		case "RED":
			level.Register(levels[i].Name, color.New(color.NormalDisplay, color.Red))
		case "GREEN":
			level.Register(levels[i].Name, color.New(color.NormalDisplay, color.Green))
		case "YELLOW":
			level.Register(levels[i].Name, color.New(color.NormalDisplay, color.Yellow))
		case "BLUE":
			level.Register(levels[i].Name, color.New(color.NormalDisplay, color.Blue))
		case "MAGENTA":
			level.Register(levels[i].Name, color.New(color.NormalDisplay, color.Magenta))
		case "CYAN":
			level.Register(levels[i].Name, color.New(color.NormalDisplay, color.Cyan))
		case "WHITE":
			level.Register(levels[i].Name, color.New(color.NormalDisplay, color.White))
		default:
			level.Register(levels[i].Name, color.NewLiteral(levels[i].Color))
		}
	}
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
			var output Output
			if output = GetOutput("file://" + os.Stdout.Name()); output == nil {
				return NewFileOutput(os.Stdout, nil, DefaultDebugFormat)
			}
			return output
		case os.Stderr.Name():
			var output Output
			if output = GetOutput("file://" + os.Stderr.Name()); output == nil {
				return NewFileOutput(os.Stderr, nil, DefaultDebugFormat)
			}
			return output
		default:
			f := lib.OpenFile(uri)
			var output Output
			if output = GetOutput(f.Name()); output == nil {
				return NewFileOutput(f, nil, DefaultDebugFormat)
			}
			return output
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
