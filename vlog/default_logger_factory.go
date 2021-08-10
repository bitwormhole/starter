package vlog

////////////////////////////////////////////////////////////////////////////////

// func Default() Logger {
// 	return nil
// }

////////////////////////////////////////////////////////////////////////////////

type DefaultLoggerFactory struct {
	impl LoggerFactory
}

func (inst *DefaultLoggerFactory) _Impl() LoggerFactory {
	return inst
}

func (inst *DefaultLoggerFactory) getImpl() LoggerFactory {
	impl := inst.impl
	if impl == nil {
		simple := &SimpleLoggerFactory{}
		impl = simple.init()
		inst.impl = impl
	}
	return impl
}

func (inst *DefaultLoggerFactory) CreateLogger(src interface{}) Logger {
	return inst.getImpl().CreateLogger(src)
}

////////////////////////////////////////////////////////////////////////////////
