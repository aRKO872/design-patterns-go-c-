package main

import "github.com/design-pattern-factory/loggers"

func main () {
	loggerFactory := loggers.NewLoggerFactory()
	debugLog := loggerFactory.CreateLogger(loggers.DEBUG_LOG)
	infoLog := loggerFactory.CreateLogger(loggers.INFO_LOG)
	errorLog := loggerFactory.CreateLogger(loggers.ERROR_LOG)

	debugLog.Log(loggers.LogStatement)
	infoLog.Log(loggers.LogStatement)
	errorLog.Log(loggers.LogStatement)
}