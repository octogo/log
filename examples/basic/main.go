package main

import (
	"github.com/octogo/log"
)

func main() {
	defer log.Close()

	log.Println("output")
	log.Printf("formatted %s", "output")

	logger := log.NewLogger("main")
	logger.Debug("debug")
	logger.Info("info")
	logger.Notice("notice")
	logger.Alert("alert")
	logger.Warning("warning")
	logger.Error("error")

	log.Fatalf("formatted %s", "error")
}
