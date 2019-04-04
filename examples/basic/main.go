package main

import (
	"github.com/octogo/log"
)

func main() {
	router := octolog.New()
	defer router.Drain()

	logger := router.NewLogger("example")
	logger.Debug("debug")
	logger.Info("info")
	logger.Notice("notice")
	logger.Alert("alert")
	logger.Warning("warning")
	logger.Error("error")
}
