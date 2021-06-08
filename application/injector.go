package application

import (
	"errors"
	"io"

	"github.com/bitwormhole/starter/lang"
)

type InjectionSource interface {
	io.Closer
	Count() int
	Selector() string
	HasMore() bool
	Read() (lang.Object, error)
}

type InjectionTarget interface {
	io.Closer
	Write(lang.Object) error
}

type Injection interface {
	io.Closer
	lang.ErrorHandler

	Context() Context
	Pool() lang.ReleasePool

	// 类选择器: ".class"
	// ID选择器: "#id"
	// 属性选择器: "${prop.name}"
	// value选择器: "foo"
	// context选择器: "context"
	Select(selector string) InjectionSource
}

type Injector interface {
	OpenInjection(context Context) (Injection, error)
}

////////////////////////////////////////////////////////////////////////////////

type FunctionInjectionTargetFn func(lang.Object) error

type FunctionInjectionTarget struct {
	fn FunctionInjectionTargetFn
}

func (inst *FunctionInjectionTarget) Init(fn FunctionInjectionTargetFn) InjectionTarget {
	inst.fn = fn
	return inst
}

func (inst *FunctionInjectionTarget) Write(o lang.Object) error {
	fn := inst.fn
	if fn == nil {
		return errors.New("FunctionInjectionTarget.fn==nil")
	}
	return fn(o)
}

func (inst *FunctionInjectionTarget) Close() error {
	return nil
}
