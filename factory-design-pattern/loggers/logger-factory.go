package loggers

type ILogger interface {
	Log(statement string)
}

type LoggerFactory struct {
}

func (l *LoggerFactory) CreateLogger(kind LogType) ILogger {
	switch kind {
		case DEBUG_LOG :
			return NewDebugLogger()
		case INFO_LOG :
			return NewInfoLogger()
		case ERROR_LOG :
			return NewErrorLogger()
		default :
			return nil
	}
}

func NewLoggerFactory() *LoggerFactory {
	return new(LoggerFactory)
}