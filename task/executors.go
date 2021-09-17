package task

// DefaultExecutor 创建一个默认的任务的执行者
func DefaultExecutor() Executor {
	return &defaultExecutor{}
}

////////////////////////////////////////////////////////////////////////////////

type defaultExecutor struct {
}

func (inst *defaultExecutor) _Impl() Executor {
	return inst
}

func (inst *defaultExecutor) Execute(r Runnable) {
	if r == nil {
		return
	}
	r.Run()
}
