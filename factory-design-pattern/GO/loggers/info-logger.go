package loggers

import "log"

type InfoLogger struct{

}

func (l *InfoLogger) Log(stmt string) {
	log.Printf("Info: %s\n", stmt)
}

func NewInfoLogger() *InfoLogger {
	return new(InfoLogger)
}

