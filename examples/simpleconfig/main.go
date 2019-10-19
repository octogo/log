package main

import (
	"github.com/octogo/log"
	"github.com/octogo/log/pkg/config"
	octolog "github.com/octogo/log/pkg/log"
)

func main() {
	c := &config.Config{
		DefaultFormat: octolog.DefaultDebugFormat,
		LoggerName:    "main",
	}
	log.InitWithConfig(c)
	l := log.New("", nil)
	l.Debug("This is a DEBUG log...")
	l.Info("This is an INFO log...")
	l.Notice("This is a NOTICE log...")
	l.Warning("This is a WARNING log...")
	l.Error("This is an ERROR log...")
}
