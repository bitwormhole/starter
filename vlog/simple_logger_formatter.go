package vlog

import (
	"fmt"
	"strings"
	"time"
)

const (
	FormatStyleFatal  = "\033[97;40m"
	FormatStyleError  = "\033[97;41m"
	FormatStyleWarn   = "\033[30;103m"
	FormatStyleInfo   = "\033[0m"
	FormatStyleDebug  = "\033[97;42m"
	FormatStyleTrace  = "\033[97;46m"
	FormatStyleNormal = "\033[0m"
)

type SimpleFormatter struct {
}

func (inst *SimpleFormatter) _Impl() Formatter {
	return inst
}

func (inst *SimpleFormatter) Format(r *Record) string {

	tt := time.Unix(r.Timestamp/1000, 0)
	args := r.Arguments

	builder := strings.Builder{}
	builder.WriteString(tt.String())

	builder.WriteString(" ")
	builder.WriteString(inst.getStyle(r))
	builder.WriteString("[")
	builder.WriteString(r.Level.String())
	builder.WriteString("]")
	builder.WriteString(FormatStyleNormal)
	builder.WriteString(" ")

	builder.WriteString(fmt.Sprint(args...))
	return builder.String()
}

func (inst *SimpleFormatter) getStyle(r *Record) string {
	switch r.Level {
	case INFO:
		return FormatStyleInfo

	case TRACE:
		return FormatStyleTrace

	case DEBUG:
		return FormatStyleDebug

	case WARN:
		return FormatStyleWarn

	case ERROR:
		return FormatStyleError

	case FATAL:
		return FormatStyleFatal
	}
	return FormatStyleNormal
}
