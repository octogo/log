package log

import (
	"os"

	"github.com/octogo/log/pkg/config"
	"github.com/octogo/log/pkg/level"
	"github.com/octogo/log/pkg/log"
)

// ConfigFile defines the file-name of the configuration file.
var ConfigFile = "logging.yml"

// LoadConfig returns the loaded configuration.
func LoadConfig(paths ...string) *config.Config {
	if paths == nil || len(paths) == 0 {
		paths = []string{ConfigFile}
	}
	return config.Load(paths...)
}

// Init initialized octolog.
func Init() {
	if _, err := os.Stat(ConfigFile); err != nil {
		log.Init()
	} else {
		log.Configure(config.Load(ConfigFile))
	}
}

// InitWithConfig intializes octolog with a custom configuration.
func InitWithConfig(c *config.Config) {
	log.Configure(c)
}

// New returns an initialized Logger.
func New(name string, wants []level.Level, outputs ...string) *log.Logger {
	return log.NewLogger(name, wants, outputs...)
}

// Println logs the given value with log-level INFO.
func Println(v interface{}) {
	log.Println(v)
}

// Printf wraps Println and supports string formatting.
func Printf(f string, args ...interface{}) {
	log.Printf(f, args...)
}

// Fatal logs the given value with log-level ERROR and exits with RC-1.
func Fatal(v interface{}) {
	log.Fatal(v)
}

// Fatalf wraps Fatal() and supports string formatting.
func Fatalf(f string, args ...interface{}) {
	log.Fatalf(f, args...)
}
