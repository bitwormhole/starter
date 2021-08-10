package vlog

var defaultLoggerInstance Logger

// SetDefaultLogger 设置默认 Logger
func SetDefaultLogger(logger Logger) {
	if logger == nil {
		return
	}
	defaultLoggerInstance = logger
}

// Default 获取默认 Logger
func Default() Logger {
	logger := defaultLoggerInstance
	if logger == nil {
		logger = initSimpleDefaultLogger()
		defaultLoggerInstance = logger
	}
	return logger
}

func initSimpleDefaultLogger() Logger {
	factory := &SimpleLoggerFactory{}
	return factory.CreateLogger(factory)
}

// Debug 输出日志
func Debug(a ...interface{}) Logger {
	return Default().Debug(a...)
}

// Error 输出日志
func Error(a ...interface{}) Logger {
	return Default().Error(a...)
}

// Fatal 输出日志
func Fatal(a ...interface{}) Logger {
	return Default().Fatal(a...)
}

// Info 输出日志
func Info(a ...interface{}) Logger {
	return Default().Info(a...)
}

// Trace 输出日志
func Trace(a ...interface{}) Logger {
	return Default().Trace(a...)
}

// Warn 输出日志
func Warn(a ...interface{}) Logger {
	return Default().Warn(a...)
}
