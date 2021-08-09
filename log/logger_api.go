package log

// Logger 是一个统一的抽象日志接口
type Logger interface {
	Fatal(a ...interface{})
	Error(a ...interface{})
	Warn(a ...interface{})
	Info(a ...interface{})
	Debug(a ...interface{})
	Trace(a ...interface{})

	IsFatalEnabled() bool
	IsErrorEnabled() bool
	IsWarnEnabled() bool
	IsInfoEnabled() bool
	IsDebugEnabled() bool
	IsTraceEnabled() bool
}

// LoggerFactory 是Logger的工厂
type LoggerFactory interface {
	CreateLogger() Logger
}
