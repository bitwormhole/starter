package vlog

type LoggerAdapter struct {

	// public
	LoggerFactory LoggerFactory

	// private
	next Logger
}

func (inst *LoggerAdapter) _Impl() Logger {
	return inst
}

func (inst *LoggerAdapter) loadNext() Logger {
	return inst.LoggerFactory.CreateLogger(inst)
}

func (inst *LoggerAdapter) getNext() Logger {
	next := inst.next
	if next == nil {
		next = inst.loadNext()
		inst.next = next
	}
	return next
}

// func (inst *LoggerAdapter) SetSource(s interface{}) {
// 	inst.getNext().SetSource(s)
// }

func (inst *LoggerAdapter) Info(a ...interface{}) Logger {
	inst.getNext().Info(a)
	return inst
}

func (inst *LoggerAdapter) Debug(a ...interface{}) Logger {
	inst.getNext().Debug(a)
	return inst
}

func (inst *LoggerAdapter) Warn(a ...interface{}) Logger {
	inst.getNext().Warn(a)
	return inst
}

func (inst *LoggerAdapter) Error(a ...interface{}) Logger {
	inst.getNext().Error(a)
	return inst
}

func (inst *LoggerAdapter) Fatal(a ...interface{}) Logger {
	inst.getNext().Fatal(a)
	return inst
}

func (inst *LoggerAdapter) Trace(a ...interface{}) Logger {
	inst.getNext().Trace(a)
	return inst
}

func (inst *LoggerAdapter) IsInfoEnabled() bool {
	return inst.getNext().IsInfoEnabled()
}

func (inst *LoggerAdapter) IsDebugEnabled() bool {
	return inst.getNext().IsDebugEnabled()
}

func (inst *LoggerAdapter) IsWarnEnabled() bool {
	return inst.getNext().IsWarnEnabled()
}

func (inst *LoggerAdapter) IsErrorEnabled() bool {
	return inst.getNext().IsErrorEnabled()
}

func (inst *LoggerAdapter) IsFatalEnabled() bool {
	return inst.getNext().IsFatalEnabled()
}

func (inst *LoggerAdapter) IsTraceEnabled() bool {
	return inst.getNext().IsTraceEnabled()
}
