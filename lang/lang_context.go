package lang

import (
	"context"
	"errors"
	"fmt"
	"time"
)

// Context 是对 context.Context 的扩展，是个可编辑的上下文
type Context interface {
	context.Context
	GetValue(key string) interface{}
	SetValue(key string, value interface{})
}

// GetContext 获取已绑定的可编辑上下文
func GetContext(cc context.Context) (Context, error) {
	binding, err := getContextBinding(cc)
	if err != nil {
		return nil, err
	}
	result := binding.context
	if result == nil {
		return nil, errors.New("no context is bound(2), need lang.SetupContext(ctx)")
	}
	return result, nil
}

// SetupContext 绑定当前上下文
func SetupContext(lc Context) error {
	binding, err := openContextBinding(lc)
	if err != nil {
		return err
	}
	binding.context = lc
	return nil
}

////////////////////////////////////////////////////////////////////////////////

// SimpleContext 一个简单的 lang.Context 实现
type SimpleContext struct {
	atts     map[string]interface{}
	deadline time.Time
	err      error
}

func (inst *SimpleContext) _Impl() Context {
	return inst
}

func (inst *SimpleContext) getAtts() map[string]interface{} {
	table := inst.atts
	if table == nil {
		table = make(map[string]interface{})
		inst.atts = table
	}
	return table
}

// GetValue 取属性值
func (inst *SimpleContext) GetValue(key string) interface{} {
	table := inst.getAtts()
	return table[key]
}

// SetValue 设置属性值
func (inst *SimpleContext) SetValue(key string, value interface{}) {
	table := inst.getAtts()
	table[key] = value
}

// Deadline 取上下文的截止日期
func (inst *SimpleContext) Deadline() (deadline time.Time, ok bool) {
	return inst.deadline, false
}

// Done ... of context.Context
func (inst *SimpleContext) Done() <-chan struct{} {
	return nil
}

// Err ... of context.Context
func (inst *SimpleContext) Err() error {
	return inst.err
}

// Value 取属性值
func (inst *SimpleContext) Value(key interface{}) interface{} {
	name, ok := key.(string)
	if !ok {
		name = fmt.Sprint(key)
	}
	return inst.GetValue(name)
}

////////////////////////////////////////////////////////////////////////////////
// implementation

const contextBindingKey = "github.com/bitwormhole/starter/lang/Context#binding"

type contextBinding struct {
	context Context
}

func getContextBinding(cc context.Context) (*contextBinding, error) {
	const key = contextBindingKey
	if cc == nil {
		return nil, errors.New("context==nil")
	}
	o1 := cc.Value(key)
	if o1 != nil {
		o2, ok := o1.(*contextBinding)
		if ok {
			return o2, nil
		}
	}
	return nil, errors.New("no context is bound(1), need lang.BindContext(ctx)")
}

func openContextBinding(lc Context) (*contextBinding, error) {
	const key = contextBindingKey
	if lc == nil {
		return nil, errors.New("context==nil")
	}
	o1 := lc.GetValue(key)
	if o1 != nil {
		o2, ok := o1.(*contextBinding)
		if ok {
			return o2, nil
		}
	}
	binding := &contextBinding{}
	binding.context = lc
	lc.SetValue(key, binding)
	return binding, nil
}

////////////////////////////////////////////////////////////////////////////////
