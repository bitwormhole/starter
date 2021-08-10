package vlog

import (
	"fmt"
	"strings"
	"time"
)

type SimpleLoggerFactory struct {
	level Level
	chain FilterChain
}

func (inst *SimpleLoggerFactory) init() LoggerFactory {
	inst.level = INFO
	inst.chain = (&simpleChain{}).init()
	return inst
}

func (inst *SimpleLoggerFactory) getChain() FilterChain {
	ch := inst.chain
	if ch == nil {
		ch = &simpleChain{}
		inst.chain = ch
	}
	return ch
}

func (inst *SimpleLoggerFactory) CreateLogger(source interface{}) Logger {
	sl := &simpleLogger{}
	sl.factory = inst
	sl.chain = inst.getChain()
	sl.source = source
	return sl
}

////////////////////////////////////////////////////////////////////////////////

type simpleChain struct {
}

func (inst *simpleChain) init() FilterChain {
	return inst
}

func (inst *simpleChain) formatLevel(l Level) string {
	switch l {
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

func (inst *simpleChain) format(r *Record) string {

	tt := time.Unix(r.Timestamp/1000, 0)
	args := r.Arguments

	builder := strings.Builder{}
	builder.WriteString(tt.String())
	builder.WriteString(" [" + inst.formatLevel(r.Level) + "] ")
	builder.WriteString(fmt.Sprint(args...))
	return builder.String()
}

func (inst *simpleChain) Append(r *Record) {
	text := inst.format(r)
	fmt.Println(text)
}

////////////////////////////////////////////////////////////////////////////////

type simpleLogger struct {
	factory *SimpleLoggerFactory
	chain   FilterChain
	source  interface{}
}

func (inst *simpleLogger) _Impl() Logger {
	return inst
}

func (inst *simpleLogger) isLevelEnabled(level Level) bool {
	return inst.factory.level <= level
}

func (inst *simpleLogger) makeRecord(level Level, a []interface{}) *Record {
	r := &Record{}
	r.Level = level
	r.Arguments = a
	r.Timestamp = inst.now()
	r.Source = inst.source
	return r
}

func (inst *simpleLogger) now() int64 {
	sec := time.Now().Unix()
	return sec * 1000
}

func (inst *simpleLogger) append(level Level, a []interface{}) Logger {
	r := inst.makeRecord(level, a)
	inst.chain.Append(r)
	return inst
}

func (inst *simpleLogger) SetSource(s interface{}) {
	if s == nil {
		return
	}
	inst.source = s
}

func (inst *simpleLogger) Debug(a ...interface{}) Logger {
	return inst.append(DEBUG, a)
}

func (inst *simpleLogger) Error(a ...interface{}) Logger {
	return inst.append(ERROR, a)
}

func (inst *simpleLogger) Fatal(a ...interface{}) Logger {
	return inst.append(FATAL, a)
}

func (inst *simpleLogger) Info(a ...interface{}) Logger {
	return inst.append(INFO, a)
}

func (inst *simpleLogger) Trace(a ...interface{}) Logger {
	return inst.append(TRACE, a)
}

func (inst *simpleLogger) Warn(a ...interface{}) Logger {
	return inst.append(WARN, a)
}

func (inst *simpleLogger) IsDebugEnabled() bool {
	return inst.isLevelEnabled(DEBUG)
}

func (inst *simpleLogger) IsErrorEnabled() bool {
	return inst.isLevelEnabled(ERROR)
}

func (inst *simpleLogger) IsFatalEnabled() bool {
	return inst.isLevelEnabled(FATAL)
}

func (inst *simpleLogger) IsInfoEnabled() bool {
	return inst.isLevelEnabled(INFO)
}

func (inst *simpleLogger) IsTraceEnabled() bool {
	return inst.isLevelEnabled(TRACE)
}

func (inst *simpleLogger) IsWarnEnabled() bool {
	return inst.isLevelEnabled(WARN)
}
