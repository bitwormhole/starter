package vlog

import (
	"errors"
	"strings"
)

const (
	// ALL 启用所有信息
	ALL = 1

	// TRACE  跟踪信息
	TRACE = 2

	// DEBUG 调试信息
	DEBUG = 3

	// INFO 常规信息
	INFO = 4

	// WARN 警告信息
	WARN = 5

	// ERROR  错误信息
	ERROR = 6

	// FATAL  致命错误信息
	FATAL = 7

	// NONE  禁用所有信息
	NONE = 8
)

// Level 日志信息等级
type Level int

// Logger 是一个统一的抽象日志接口
type Logger interface {
	Fatal(a ...interface{}) Logger
	Error(a ...interface{}) Logger
	Warn(a ...interface{}) Logger
	Info(a ...interface{}) Logger
	Debug(a ...interface{}) Logger
	Trace(a ...interface{}) Logger

	// 废弃： SetSource(s interface{})

	IsFatalEnabled() bool
	IsErrorEnabled() bool
	IsWarnEnabled() bool
	IsInfoEnabled() bool
	IsDebugEnabled() bool
	IsTraceEnabled() bool
}

// LoggerFactory 是Logger的工厂
type LoggerFactory interface {

	// CreateLogger 新建Logger
	CreateLogger(source interface{}) Logger
}

// Formatter 用于格式化日志
type Formatter interface {
	// Format 方法用于格式化
	Format(rec *Record) string
}

// Filter 是日志记录的过滤器
type Filter interface {
	// DoFilter 方法过滤一条记录
	DoFilter(rec *Record, chain FilterChain)
}

// Writer 是日志记录的写入器
type Writer interface {
	// Write 写入一条记录
	Write(rec *Record)
}

// Channel 是日志记录的通道
type Channel interface {
	Writer
	IsLevelEnabled(level Level) bool
}

// FilterChain 是过滤器的链条
type FilterChain interface {
	// Append 方法向链条追加一条记录
	Append(rec *Record)
}

// Record 表示一条日志记录
type Record struct {
	// 时间戳
	Timestamp int64

	// 记录等级
	Level Level

	// 参数
	Arguments []interface{}

	// Source 表示本条记录的来源
	Source interface{}

	// Message 是由 Formatter 生成的字符串
	Message string
}

////////////////////////////////////////////////////////////////////////////////

// ParseLevel 解析一个字符串为日志等级
func ParseLevel(str string) (Level, error) {
	str = strings.TrimSpace(str)
	str = strings.ToUpper(str)
	switch str {
	case "FATAL":
		return FATAL, nil
	case "ERROR":
		return ERROR, nil
	case "WARN":
		return WARN, nil
	case "INFO":
		return INFO, nil
	case "DEBUG":
		return DEBUG, nil
	case "TRACE":
		return TRACE, nil
	case "ALL":
		return ALL, nil
	default:
		return NONE, errors.New("bad level string:" + str)
	}
}

func (lv Level) String() string {
	switch lv {
	case TRACE:
		return "TRACE"
	case DEBUG:
		return "DEBUG"
	case INFO:
		return "INFO "
	case WARN:
		return "WARN "
	case ERROR:
		return "ERROR"
	case FATAL:
		return "FATAL"
	default:
		return "UNDEF"
	}
}
