package lang

import "errors"

// ErrorCollector 表示一个错误收集器
type ErrorCollector interface {
	Append(err error)
	AppendIfFalse(ok bool, msg string)
	AppendIfNil(target Object, msg string)

	LastError() error
	Result() error
}

// NewErrorCollector 创建一个错误收集器
func NewErrorCollector() ErrorCollector {
	return &DefaultErrorCollector{}
}

type DefaultErrorCollector struct {
	all     []error
	lastErr error
}

func (inst *DefaultErrorCollector) Result() error {
	all := inst.all
	if all == nil {
		return nil
	}
	if len(all) < 1 {
		return nil
	}
	return all[0]
}

func (inst *DefaultErrorCollector) LastError() error {
	return inst.lastErr
}

func (inst *DefaultErrorCollector) Append(err error) {
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
	inst.lastErr = err
}

func (inst *DefaultErrorCollector) AppendIfNil(value Object, msg string) {
	if value == nil {
		inst.Append(errors.New(msg))
	}
}

func (inst *DefaultErrorCollector) AppendIfFalse(value bool, msg string) {
	if !value {
		inst.Append(errors.New(msg))
	}
}
