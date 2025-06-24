package loggers

import "log"

type DebugLogger struct{

}

func (l *DebugLogger) Log(stmt string) {
	log.Printf("Debug: %s\n", stmt)
}

func NewDebugLogger() *DebugLogger {
	return new(DebugLogger)
}