package runtime

import (
	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/lang"
)

type innerContextGetter struct {
	context application.RuntimeContext
	errors  lang.ErrorCollector
}

func (inst *innerContextGetter) init(context application.RuntimeContext, ec lang.ErrorCollector) {

	if ec == nil {
		ec = lang.NewErrorCollector()
	}

	inst.errors = ec
	inst.context = context
}

func (inst *innerContextGetter) ErrorCollector() lang.ErrorCollector {
	return inst.errors
}

func (inst *innerContextGetter) Result() error {
	return inst.errors.Result()
}

func (inst *innerContextGetter) GetProperty(name string) string {
	text, err := inst.context.GetProperties().GetPropertyRequired(name)
	if err == nil {
		return text
	} else {
		inst.errors.AddError(err)
		return ""
	}
}

func (inst *innerContextGetter) GetPropertySafely(name string, _default string) string {
	return inst.context.GetProperties().GetProperty(name, _default)
}

func (inst *innerContextGetter) GetComponent(name string) lang.Object {
	com, err := inst.context.GetComponents().GetComponent(name)
	if err == nil {
		return com
	} else {
		inst.errors.AddError(err)
		return nil
	}
}

func (inst *innerContextGetter) GetComponentByClass(classSelector string) lang.Object {
	com, err := inst.context.GetComponents().GetComponentByClass(classSelector)
	if err == nil {
		return com
	} else {
		inst.errors.AddError(err)
		return nil
	}
}

func (inst *innerContextGetter) GetComponentsByClass(classSelector string) []lang.Object {
	return inst.context.GetComponents().GetComponentsByClass(classSelector)
}
