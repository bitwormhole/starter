package vlog

const (
	// ALL 启用所有信息
	ALL = 0

	// TRACE  跟踪信息
	TRACE = 1

	// DEBUG 调试信息
	DEBUG = 2

	// INFO 常规信息
	INFO = 3

	// WARN 警告信息
	WARN = 4

	// ERROR  错误信息
	ERROR = 5

	// FATAL  致命错误信息
	FATAL = 6

	// NONE  禁用所有信息
	NONE = 7
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

	SetSource(s interface{})

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
	Format(rec Record) string
}

// Filter 是日志记录的过滤器
type Filter interface {
	// DoFilter 方法过滤一条记录
	DoFilter(rec *Record, chain FilterChain)
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
