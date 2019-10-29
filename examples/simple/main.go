package main

import (
	"github.com/octogo/log"
)

func main() {
	log.Init()
	log.Log("Hello world!")
	logger := log.New("myapp", nil)
	logger.Debug("...")
	logger.Info("...")
	logger.Notice("...")
	logger.Warning("...")
	logger.Error("...")
}
