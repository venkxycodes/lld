package main

import (
	"logger/logger"
	out_writer "logger/out-writer"
)

func main() {
	// Console logger
	consoleSink := out_writer.NewConsoleSink()
	consoleLogger, _ := logger.NewLogger(logger.Info, consoleSink)
	defer consoleLogger.Close()

	consoleLogger.Debug("this debug message will be filtered")
	consoleLogger.Info("application started successfully")
	consoleLogger.Error("failed to connect to DB")

	// File logger
	fileSink, _ := out_writer.NewFileSink("app.log")
	fileLogger, _ := logger.NewLogger(logger.Debug, fileSink)
	defer fileLogger.Close()

	fileLogger.Debug("writing debug logs to file")
	fileLogger.Error("something went wrong")
}
