package std

import (
	"fmt"

	"github.com/bitwormhole/starter/vlog"
)

type ConsoleWriter struct {
	mDefaultFormatter vlog.Formatter
}

func (inst *ConsoleWriter) _Impl() vlog.Writer {
	return inst
}

func (inst *ConsoleWriter) getDefaultFormatter() vlog.Formatter {
	f := inst.mDefaultFormatter
	if f == nil {
		f = &DefaultFormatter{}
		inst.mDefaultFormatter = f
	}
	return f
}

func (inst *ConsoleWriter) Write(rec *vlog.Record) {
	if rec == nil {
		return
	}
	text := rec.Message
	if text == "" {
		text = inst.getDefaultFormatter().Format(rec)
	}
	fmt.Println(text)
}
