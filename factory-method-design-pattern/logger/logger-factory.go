package logger

type ILoggerFactory interface {
	CreateLogger() *ILogger
}

type InfoLoggerFactory struct{
}

type DebugLoggerFactory struct{
}

type ErrorLoggerFactory struct{
}

func (l *InfoLoggerFactory) CreateLogger() ILogger {
	return newInfoLogger()
}

func (l *DebugLoggerFactory) CreateLogger() ILogger {
	return newDebugLogger()
}

func (l *ErrorLoggerFactory) CreateLogger() ILogger {
	return newErrorLogger()
}

func NewInfoLoggerFactory() *InfoLoggerFactory {
	return new(InfoLoggerFactory)
}

func NewDebugLoggerFactory() *DebugLoggerFactory {
	return new(DebugLoggerFactory)
}

func NewErrorLoggerFactory() *ErrorLoggerFactory {
	return new(ErrorLoggerFactory)
}