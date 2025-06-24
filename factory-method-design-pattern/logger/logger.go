package logger

import "log"

type ILogger interface {
	Log(stmt string) 
}

type InfoLogger struct{
}

type DebugLogger struct{
}

type ErrorLogger struct{
}

func (l *InfoLogger) Log(stmt string) {
	log.Printf("Info: %s\n", stmt)
}

func (l *DebugLogger) Log(stmt string) {
	log.Printf("Debug: %s\n", stmt)
}

func (l *ErrorLogger) Log(stmt string) {
	log.Printf("Error: %s\n", stmt)
}

func newInfoLogger() *InfoLogger {
	return new(InfoLogger)
}

func newDebugLogger() *DebugLogger {
	return new(DebugLogger)
}

func newErrorLogger() *ErrorLogger {
	return new(ErrorLogger)
}