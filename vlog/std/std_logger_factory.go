package std

import "github.com/bitwormhole/starter/vlog"

// StandardLoggerFactory 标准的日志工厂
type StandardLoggerFactory struct {

	// public
	Context Context

	// private
	manager *stdLogManager
}

func (inst *StandardLoggerFactory) _Impl() vlog.LoggerFactory {
	return inst
}

func (inst *StandardLoggerFactory) getManager() *stdLogManager {
	m := inst.manager
	if m == nil {
		m = getManager()
		inst.manager = m
	}
	return m
}

func (inst *StandardLoggerFactory) CreateLogger(source interface{}) vlog.Logger {

	manager := inst.getManager()
	logger := &stdLogger{manager: manager}
	return logger._Impl()
}

func (inst *StandardLoggerFactory) Start() error {
	ctx := inst.Context
	manager := inst.getManager()
	chl := ctx.GetMainChannel()
	manager.SetOutput(chl)
	return nil
}

func (inst *StandardLoggerFactory) Stop() error { return nil }
