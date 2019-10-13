package config

import "github.com/jinzhu/configor"

// Config is a data container for loaded configuration.
type Config struct {
	DefaultFormat  string `default:"{{.Date}} {{.Time}} {{.Logger}} {{.Level}} {{.Message}}"`
	LoggerName     string `default:"octolog"`
	DefaultOutputs []string
	Outputs        []Output
	Loggers        []Logger
}

// Output is a helper for loading output configuration.
type Output struct {
	URL    string `required:"true"`
	Wants  []string
	Format string
}

// Logger is a helper for loading logger configuration.
type Logger struct {
	Name    string `required:"true"`
	Wants   []string
	Outputs []string
}

// Load returns the loaded configuration.
func Load(paths ...string) *Config {
	config := &Config{}
	configor.Load(config, paths...)
	return config
}
