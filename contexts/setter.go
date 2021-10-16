package contexts

import (
	"context"
	"errors"
	"fmt"
	"time"
)

// 【已废弃，用“lang.ContextSetter”代替】 Context 是对 context.Context 的扩展，是个可编辑的上下文
// type Context interface {
// 	context.Context
// 	GetValue(key string) interface{}
// 	SetValue(key string, value interface{})
// }

// ContextSetter 是 context.Context 独享的设置入口
type ContextSetter interface {
	GetContext() context.Context
	SetValue(key interface{}, value interface{})
}

// GetContextSetter 获取已绑定的可编辑上下文
func GetContextSetter(cc context.Context) (ContextSetter, error) {
	binding, err := getContextBinding(cc)
	if err != nil {
		return nil, err
	}
	setter := binding.setter
	if setter == nil {
		return nil, errors.New("no context is bound(2), need lang.SetupContextSetter(ctx)")
	}
	return setter, nil
}

// SetupContextSetter 绑定当前上下文
func SetupContextSetter(setter ContextSetter) error {
	_, err := openContextBinding(setter)
	return err
}

////////////////////////////////////////////////////////////////////////////////

// SimpleContext 一个简单的 lang.Context 实现
type SimpleContext struct {
	atts     map[string]interface{}
	deadline time.Time
	err      error
}

func (inst *SimpleContext) _Impl() (ContextSetter, context.Context) {
	return inst, inst
}

func (inst *SimpleContext) getAtts() map[string]interface{} {
	table := inst.atts
	if table == nil {
		table = make(map[string]interface{})
		inst.atts = table
	}
	return table
}

// GetContext 取setter的context
func (inst *SimpleContext) GetContext() context.Context {
	return inst
}

// SetValue 设置属性值
func (inst *SimpleContext) SetValue(key interface{}, value interface{}) {
	name := stringifyKey(key)
	table := inst.getAtts()
	table[name] = value
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
	name := stringifyKey(key)
	table := inst.getAtts()
	return table[name]
}

////////////////////////////////////////////////////////////////////////////////
// implementation

const contextBindingKey = "github.com/bitwormhole/starter/lang/ContextSetter#binding"

type contextBinding struct {
	setter ContextSetter
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
	return nil, errors.New("no context is bound(1), need lang.SetupContextSetter(ctx)")
}

func openContextBinding(setter ContextSetter) (*contextBinding, error) {
	const key = contextBindingKey
	if setter == nil {
		return nil, errors.New("contextSetter==nil")
	}
	o1 := setter.GetContext().Value(key)
	if o1 != nil {
		o2, ok := o1.(*contextBinding)
		if ok {
			return o2, nil
		}
	}
	binding := &contextBinding{}
	binding.setter = setter
	setter.SetValue(key, binding)
	return binding, nil
}

////////////////////////////////////////////////////////////////////////////////

func stringifyKey(key interface{}) string {
	str, ok := key.(string)
	if !ok {
		str = fmt.Sprint(key)
	}
	return str
}
