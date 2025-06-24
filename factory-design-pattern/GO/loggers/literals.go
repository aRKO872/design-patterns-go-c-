package loggers

type LogType int

const (
	INFO_LOG LogType = iota 
	DEBUG_LOG
	ERROR_LOG
)

var (
	LogStatement = "Hello, Arko here"
)

