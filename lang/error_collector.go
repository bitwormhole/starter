package lang

import "errors"

// ErrorCollector 表示一个错误收集器
type ErrorCollector interface {
	AddError(err error)
	AddErrorIfFalse(ok bool, msg string)
	AddErrorIfNil(target Object, msg string)
	Result() error
}

// NewErrorCollector 创建一个错误收集器
func NewErrorCollector() ErrorCollector {
	return &innerErrorCollector{}
}

type innerErrorCollector struct {
	all []error
}

func (inst *innerErrorCollector) Result() error {
	all := inst.all
	if all == nil {
		return nil
	}
	if len(all) < 1 {
		return nil
	}
	return all[0]
}

func (inst *innerErrorCollector) AddError(err error) {
	if err == nil {
		return
	}
	all := inst.all
	if all == nil {
		all = []error{err}
	} else {
		all = append(all, err)
	}
	inst.all = all
}

func (inst *innerErrorCollector) AddErrorIfNil(value Object, msg string) {
	if value == nil {
		inst.AddError(errors.New(msg))
	}
}

func (inst *innerErrorCollector) AddErrorIfFalse(value bool, msg string) {
	if !value {
		inst.AddError(errors.New(msg))
	}
}
