package log

import (
	"os"

	"github.com/octogo/log/pkg/config"
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
