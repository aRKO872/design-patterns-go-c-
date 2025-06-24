package loggers

import "log"

type ErrorLogger struct{

}

func (l *ErrorLogger) Log(stmt string) {
	log.Printf("Error: %s\n", stmt)
}

func NewErrorLogger() *ErrorLogger {
	return new(ErrorLogger)
}