package lang

import (
	"context"
	"errors"
)

// Context 是对 context.Context 的扩展，是个可编辑的上下文
type Context interface {
	context.Context

	GetValue(key string) interface{}
	SetValue(key string, value interface{})
}

////////////////////////////////////////////////////////////////////////////////

const contextBindingKey = "github.com/bitwormhole/starter/lang/Context#binding"

type contextBinding struct {
	context Context
}

func getContextBinding(readonly context.Context) (*contextBinding, error) {
	const key = contextBindingKey
	if readonly == nil {
		return nil, errors.New("context==nil")
	}
	o1 := readonly.Value(key)
	if o1 != nil {
		o2, ok := o1.(*contextBinding)
		if ok {
			return o2, nil
		}
	}
	return nil, errors.New("no context is bound(1), need lang.BindContext(ctx)")
}

func openContextBinding(ctx Context) (*contextBinding, error) {
	const key = contextBindingKey
	if ctx == nil {
		return nil, errors.New("context==nil")
	}
	o1 := ctx.GetValue(key)
	if o1 != nil {
		o2, ok := o1.(*contextBinding)
		if ok {
			return o2, nil
		}
	}
	binding := &contextBinding{}
	binding.context = ctx
	ctx.SetValue(key, binding)
	return binding, nil
}

////////////////////////////////////////////////////////////////////////////////

// EditContext 获取已绑定的可编辑上下文
func EditContext(readonly context.Context) (Context, error) {
	binding, err := getContextBinding(readonly)
	if err != nil {
		return nil, err
	}
	result := binding.context
	if result == nil {
		return nil, errors.New("no context is bound(2), need lang.BindContext(ctx)")
	}
	return result, nil
}

// BindContext 绑定当前上下文
func BindContext(ctx Context) error {
	binding, err := openContextBinding(ctx)
	if err != nil {
		return err
	}
	binding.context = ctx
	return nil
}
