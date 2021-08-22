package vlog

var defaultLoggerInstance Logger
var defaultLoggerFactory LoggerFactory
var defaultLoggerLocked bool // 锁定当前的日志系统

// SetDefaultFactory 设置默认 Logger 工厂
func SetDefaultFactory(f LoggerFactory) {
	if f == nil {
		return
	}
	if defaultLoggerLocked {
		if defaultLoggerFactory != nil || defaultLoggerInstance != nil {
			return
		}
	}
	defaultLoggerFactory = f
	defaultLoggerInstance = nil
}

// LockDefaultFactory 锁定当前的默认 Logger 工厂
func LockDefaultFactory() {
	defaultLoggerLocked = true
}

// 获取默认的日志工厂
func getDefaultLoggerFactory() LoggerFactory {
	f := defaultLoggerFactory
	if f == nil {
		f = &SimpleLoggerFactory{}
		defaultLoggerFactory = f
	}
	return f
}

// Default 获取默认 Logger
func Default() Logger {
	logger := defaultLoggerInstance
	if logger == nil {
		logger = getDefaultLoggerFactory().CreateLogger(nil)
		defaultLoggerInstance = logger
	}
	return logger
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
