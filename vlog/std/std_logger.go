package std

import (
	"github.com/bitwormhole/starter/util"
	"github.com/bitwormhole/starter/vlog"
)

type stdLogger struct {
	manager *stdLogManager
	output  vlog.Channel
}

func (inst *stdLogger) _Impl() vlog.Logger {
	return inst
}

func (inst *stdLogger) getManager() *stdLogManager {
	m := inst.manager
	if m == nil {
		m = getManager()
		inst.manager = m
	}
	return m
}

func (inst *stdLogger) isLevelEnabled(level vlog.Level) bool {
	return inst.getOutputChannel().IsLevelEnabled(level)
}

func (inst *stdLogger) getOutputChannel() vlog.Channel {
	chl := inst.output
	if chl == nil {
		chl = inst.getManager().GetOutput()
		inst.output = chl
	}
	return chl
}

func (inst *stdLogger) log(level vlog.Level, a ...interface{}) vlog.Logger {

	if !inst.isLevelEnabled(level) {
		return inst
	}

	rec := &vlog.Record{}
	rec.Level = level
	rec.Arguments = a
	rec.Source = nil
	rec.Timestamp = util.CurrentTimestamp()

	inst.getOutputChannel().Write(rec)

	return inst
}

func (inst *stdLogger) Fatal(a ...interface{}) vlog.Logger {
	return inst.log(vlog.FATAL, a...)
}

func (inst *stdLogger) Error(a ...interface{}) vlog.Logger {
	return inst.log(vlog.ERROR, a...)
}

func (inst *stdLogger) Warn(a ...interface{}) vlog.Logger {
	return inst.log(vlog.WARN, a...)
}

func (inst *stdLogger) Info(a ...interface{}) vlog.Logger {
	return inst.log(vlog.INFO, a...)
}

func (inst *stdLogger) Debug(a ...interface{}) vlog.Logger {
	return inst.log(vlog.DEBUG, a...)
}

func (inst *stdLogger) Trace(a ...interface{}) vlog.Logger {
	return inst.log(vlog.TRACE, a...)
}

func (inst *stdLogger) IsFatalEnabled() bool {
	return inst.isLevelEnabled(vlog.FATAL)
}

func (inst *stdLogger) IsErrorEnabled() bool {
	return inst.isLevelEnabled(vlog.ERROR)
}

func (inst *stdLogger) IsWarnEnabled() bool {
	return inst.isLevelEnabled(vlog.WARN)
}

func (inst *stdLogger) IsInfoEnabled() bool {
	return inst.isLevelEnabled(vlog.INFO)
}

func (inst *stdLogger) IsDebugEnabled() bool {
	return inst.isLevelEnabled(vlog.DEBUG)
}

func (inst *stdLogger) IsTraceEnabled() bool {
	return inst.isLevelEnabled(vlog.TRACE)
}
