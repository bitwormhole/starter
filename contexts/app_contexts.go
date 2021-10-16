package contexts

import (
	"context"
	"errors"

	"github.com/bitwormhole/starter/application"
)

// GetApplicationContext 取 application.Context
func GetApplicationContext(cc context.Context) (application.Context, error) {
	binding, err := getApplicationContextBinding(cc)
	if err != nil {
		return nil, err
	}
	if !binding.isReady() {
		return nil, errors.New("need invoke contexts.SetupApplicationContext()")
	}
	return binding.ac, nil
}

// SetupApplicationContext 设置 application.Context
func SetupApplicationContext(ac application.Context) error {
	if ac == nil {
		return errors.New("ac==nil")
	}
	binding, err := openApplicationContextBinding(ac)
	if err != nil {
		return err
	}
	if binding.isReady() {
		return nil
	}
	binding.init(ac)
	return nil
}

////////////////////////////////////////////////////////////////////////////////

type applicationContextSetter struct {
	ac application.Context
}

func (inst *applicationContextSetter) GetContext() context.Context {
	return inst.ac
}

func (inst *applicationContextSetter) SetValue(key interface{}, value interface{}) {
	name := stringifyKey(key)
	inst.ac.GetAttributes().SetAttribute(name, value)
}

////////////////////////////////////////////////////////////////////////////////

const applicationContextBindingName = "contexts.applicationContext#Binding"

type applicationContextBinding struct {
	ac     application.Context
	setter ContextSetter
}

func (inst *applicationContextBinding) isReady() bool {
	return (inst.ac != nil) && (inst.setter != nil)
}

func (inst *applicationContextBinding) init(ac application.Context) {
	inst.ac = ac
	inst.setter = &applicationContextSetter{ac: ac}
	SetupContextSetter(inst.setter)
}

func getApplicationContextBinding(cc context.Context) (*applicationContextBinding, error) {
	const key = applicationContextBindingName
	o1 := cc.Value(key)
	o2, ok := o1.(*applicationContextBinding)
	if !ok {
		return nil, errors.New("need invoke contexts.SetupApplicationContext()")
	}
	return o2, nil
}

func openApplicationContextBinding(ac application.Context) (*applicationContextBinding, error) {
	const key = applicationContextBindingName
	o1 := ac.Value(key)
	o2, ok := o1.(*applicationContextBinding)
	if ok {
		return o2, nil
	}
	o2 = &applicationContextBinding{}
	ac.GetAttributes().SetAttribute(key, o2)
	return o2, nil
}
