package main

import "github.com/factory-method-design-pattern/logger"

func main() {
	errLoggerFactory := logger.NewErrorLoggerFactory()
	debugLoggerFactory := logger.NewDebugLoggerFactory()
	infoLoggerFactory := logger.NewInfoLoggerFactory()

	errorLogger := errLoggerFactory.CreateLogger()
	debugLogger := debugLoggerFactory.CreateLogger()
	infoLogger := infoLoggerFactory.CreateLogger()

	errorLogger.Log(logger.LogStatement)
	debugLogger.Log(logger.LogStatement)
	infoLogger.Log(logger.LogStatement)
}