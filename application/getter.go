package application

import (
	"errors"

	"github.com/bitwormhole/starter/lang"
)

// ContextGetter 接口向 Context 的使用者提供简易的 getter 方法
type ContextGetter interface {
	ErrorCollector() lang.ErrorCollector
	Result(last_ok bool) error

	// for property
	GetProperty(last_ok bool, name string) string
	GetPropertySafely(last_ok bool, name string, _default string) string
	GetPropertyString(last_ok bool, name string, _default string) string
	GetPropertyInt(last_ok bool, name string, _default int) int

	//	for component

	//  last_ok 指出上一步操作是否成功
	//  selector 类似css中的selector，用“#xxx”表示id选择器，用“.xxx”表示类选择器
	GetComponent(last_ok bool, selector string) lang.Object

	//  last_ok 指出上一步操作是否成功
	//  selector 类似css中的selector，用“#xxx”表示id选择器，用“.xxx”表示类选择器
	GetComponents(last_ok bool, selector string) []lang.Object
}

func NewGetter(context Context) ContextGetter {
	return &innerContextGetter{context: context}
}

// innerContextGetter impl ContextGetter
type innerContextGetter struct {
	// impl  ContextGetter
	last_opt       string
	context        Context
	errorCollector lang.ErrorCollector
}

func (inst *innerContextGetter) Result(last_ok bool) error {
	inst.feedbackLastOpt(last_ok)
	return inst.errorCollector.Result()
}

func (inst *innerContextGetter) feedbackLastOpt(success bool) {
	opt := inst.last_opt
	inst.last_opt = ""
	inst.feedback(success, opt)
}

func (inst *innerContextGetter) feedback(success bool, message string) {
	if !success {
		err := errors.New("[LastOperationError msg:" + message + "]")
		inst.errorCollector.AddError(err)
	}
}

func (inst *innerContextGetter) GetProperty(last_ok bool, name string) string {
	return ""

}

func (inst *innerContextGetter) GetPropertySafely(last_ok bool, name string, _default string) string {
	return ""

}

func (inst *innerContextGetter) GetPropertyInt(last_ok bool, name string, _default int) int {
	return 0

}

func (inst *innerContextGetter) GetPropertyString(last_ok bool, name string, _default string) string {
	return ""

}

func (inst *innerContextGetter) GetComponent(last_ok bool, selector string) lang.Object {
	return nil

}

func (inst *innerContextGetter) GetComponents(last_ok bool, selector string) []lang.Object {

	// 	inst.context.GetComponents().get

	return nil
}

func (inst *innerContextGetter) ErrorCollector() lang.ErrorCollector {
	return inst.errorCollector
}
