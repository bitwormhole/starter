package application

import "github.com/bitwormhole/starter/lang"

// ContextGetter 接口向 Context 的使用者提供简易的 getter 方法
type ContextGetter interface {
	ErrorCollector() lang.ErrorCollector
	Result() error

	GetProperty(name string) string
	GetPropertySafely(name string, _default string) string

	GetComponent(name string) lang.Object
	GetComponentByClass(classSelector string) lang.Object
	GetComponentsByClass(classSelector string) []lang.Object
}

func NewGetter(context Context) ContextGetter {
	return &innerContextGetter{context: context}
}

// innerContextGetter impl ContextGetter
type innerContextGetter struct {
	ContextGetter

	context Context
}

func (inst *innerContextGetter) Result() error {
	return nil

}

func (inst *innerContextGetter) Feedback(success bool, message string) {

}

func (inst *innerContextGetter) GetProperty(name string) string {
	return ""

}

func (inst *innerContextGetter) GetPropertySafely(name string, _default string) string {
	return ""

}

func (inst *innerContextGetter) GetComponent(name string) lang.Object {
	return nil

}

func (inst *innerContextGetter) GetComponentByClass(classSelector string) lang.Object {
	return nil

}

func (inst *innerContextGetter) GetComponentsByClass(classSelector string) []lang.Object {
	return nil
}
