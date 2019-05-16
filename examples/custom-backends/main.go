package main

import (
	"os"

	"github.com/octogo/log"
)

type a struct {
	Name string
	Age  int
}

type b struct {
	*a
}

func main() {
	defer log.Close()

	logFormat := log.DefaultLogFormat
	debugFormat := log.DebugLogFormat

	var fileName = "out.log"
	var logFile *os.File

	logFile, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		log.Fatal(err)
	}

	log.SetBackends(
		log.NewStdoutBackend(
			logFormat,
			[]log.Level{log.DEBUG, log.INFO, log.NOTICE},
		),
		log.NewStderrBackend(
			logFormat,
			[]log.Level{log.ALERT, log.WARNING, log.ERROR},
		),
		log.NewFileBackend(
			debugFormat,
			log.AllLevelSlice(),
			logFile,
			false,
		),
	)

	logger := log.NewLogger("example")
	logger.Debug("debug")
	logger.Info("info")
	logger.Notice("notice")
	logger.Alert("alert")
	logger.Warning("warning")
	logger.Error("error")

	childLogger := logger.NewLogger("child")
	childLogger.Debug("debug")
	childLogger.Info("info")
	childLogger.Notice("notice")
	childLogger.Alert("alert")
	childLogger.Warning("warning")
	childLogger.Error("error")
}
