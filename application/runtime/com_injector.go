package runtime

import (
	"errors"
	"strconv"

	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/lang"
)

////////////////////////////////////////////////////////////////////////////////

type innerInjector struct {
	creation CreationContext
	_is_root bool
}

func (inst *innerInjector) init(creation CreationContext, root bool) application.Injector {
	inst.creation = creation
	inst._is_root = root
	return inst
}

func (inst *innerInjector) Inject(selector string) application.Injection {
	injection := &innerInjection{}
	injection.parent = inst
	injection.selector = selector
	return injection
}

func (inst *innerInjector) GetComponent(selector string) (lang.Object, error) {
	context := inst.creation.GetContext()
	return context.FindComponent(selector)
}

func (inst *innerInjector) Done() error {
	if inst._is_root {
		err := inst.creation.Close()
		inst.creation.HandleError(err)
	}
	return inst.creation.LastError()
}

////////////////////////////////////////////////////////////////////////////////

type innerInjection struct {
	parent          *innerInjector
	selector        string
	bAsList         bool
	bIncludeAliases bool
	fnAccept        application.ComponentHolderFilter
}

func (inst *innerInjection) AsList() application.Injection {
	inst.bAsList = true
	return inst
}

func (inst *innerInjection) IncludeAliases() application.Injection {
	inst.bIncludeAliases = true
	return inst
}

func (inst *innerInjection) Accept(f application.ComponentHolderFilter) application.Injection {
	inst.fnAccept = f
	return inst
}

func (inst *innerInjection) To(fn func(lang.Object) bool) {

	creation := inst.parent.creation
	context := creation.GetContext()
	coms := context.GetComponents()
	errSet := creation.ErrorCollector()
	fnAccept := inst.fnAccept

	if fnAccept == nil {
		fnAccept = func(name string, holder application.ComponentHolder) bool {
			return true
		}
	}

	if inst.bAsList {
		holders := coms.GetComponents(inst.selector)
		for index := range holders {
			holder := holders[index]
			name := holder.GetInfo().GetID()
			if !fnAccept(name, holder) {
				continue
			}
			inst._try_load_com(holder, creation, fn, errSet)
		}
	} else {
		holder, err := coms.GetComponent(inst.selector)
		if err != nil {
			errSet.Append(err)
			return
		}
		name := holder.GetInfo().GetID()
		if !fnAccept(name, holder) {
			return
		}
		inst._try_load_com(holder, creation, fn, errSet)
	}
}

func (inst *innerInjection) _try_load_com(holder application.ComponentHolder, cc CreationContext, fn func(lang.Object) bool, errs lang.ErrorCollector) {

	obj, err := cc.LoadComponent(holder)
	if err != nil {
		errs.Append(err)
		return
	}

	ok := fn(obj)
	if !ok {
		err := errors.New("inject fail: " + holder.GetInfo().GetID())
		errs.Append(err)
	}
}

////////////////////////////////////////////////////////////////////////////////
// impl Getter

func (inst *innerInjector) _impl_getter() application.ContextGetter {
	return inst
}

func (inst *innerInjector) GetProperty(name string) (string, error) {
	return inst.creation.GetContext().GetProperties().GetPropertyRequired(name)
}

func (inst *innerInjector) GetPropertySafely(name string, _default string) string {
	return inst.creation.GetContext().GetProperties().GetProperty(name, _default)
}

func (inst *innerInjector) GetPropertyString(name string, _default string) string {
	text, err := inst.GetProperty(name)
	if err != nil {
		return _default
	}
	return text
}

func (inst *innerInjector) GetPropertyInt(name string, _default int) int {
	text, err := inst.GetProperty(name)
	if err != nil {
		return _default
	}
	n, err := strconv.Atoi(text)
	if err != nil {
		return _default
	}
	return n
}
