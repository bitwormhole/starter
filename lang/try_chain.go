package lang

type TryChain struct {
	err error
}

func (inst *TryChain) Try(task func() error) *TryChain {
	if inst.err == nil {
		inst.err = task()
	}
	return inst
}

func (inst *TryChain) Result() error {
	return inst.err
}
