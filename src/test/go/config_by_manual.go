package main

import (
	"errors"

	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/application/config"
	"github.com/bitwormhole/starter/lang"
)

////////////////////////////////////////////////////////////////////////////////

type myFooFactory struct {
}

func (inst *myFooFactory) _Impl() (application.ComponentFactory, application.ComponentAfterService) {
	return inst, inst
}

func (inst *myFooFactory) New() lang.Object {
	return &Foo{}
}

func (inst *myFooFactory) Cast(i application.ComponentInstance) (*Foo, error) {
	o2, ok := i.Get().(*Foo)
	if ok {
		return o2, nil
	}
	return nil, errors.New("cannot cast object to *Foo")
}

func (inst *myFooFactory) GetPrototype() lang.Object {
	return inst.New()
}

func (inst *myFooFactory) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.New())
}

func (inst *myFooFactory) AfterService() application.ComponentAfterService {
	return inst
}

func (inst *myFooFactory) Inject(instance application.ComponentInstance, ic application.InstanceContext) error {

	o2, err := inst.Cast(instance)
	if err != nil {
		return err
	}

	o2.Items, err = inst.getterForItems(".bar", ic)
	if err != nil {
		return err
	}

	o2.Value, err = inst._getterForValue("${abc}", ic)
	if err != nil {
		return err
	}

	return nil
}

func (inst *myFooFactory) _getterForValue(sel string, ic application.InstanceContext) (int, error) {
	return ic.GetInt(sel)
}

func (inst *myFooFactory) getterForItems(sel string, ic application.InstanceContext) ([]*Bar, error) {

	dst := make([]*Bar, 0)
	src, err := ic.GetComponents(sel)
	if err != nil {
		return nil, err
	}

	for _, o1 := range src {
		o2, ok := o1.(*Bar)
		if ok {
			dst = append(dst, o2)
		}
	}

	return dst, nil
}

func (inst *myFooFactory) Init(instance application.ComponentInstance) error {
	o2, err := inst.Cast(instance)
	if err != nil {
		return err
	}
	return o2.Begin()
}

func (inst *myFooFactory) Destroy(instance application.ComponentInstance) error {
	o2, err := inst.Cast(instance)
	if err != nil {
		return err
	}
	return o2.End()
}

////////////////////////////////////////////////////////////////////////////////

type myBarFactory struct {
}

func (inst *myBarFactory) _Impl() (application.ComponentFactory, application.ComponentAfterService) {
	return inst, inst
}

func (inst *myBarFactory) New() lang.Object {
	return &Bar{}
}

func (inst *myBarFactory) Cast(i application.ComponentInstance) (*Bar, error) {
	o2, ok := i.Get().(*Bar)
	if ok {
		return o2, nil
	}
	return nil, errors.New("cannot cast object to *Bar")
}

func (inst *myBarFactory) GetPrototype() lang.Object {
	return inst.New()
}

func (inst *myBarFactory) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.New())
}

func (inst *myBarFactory) AfterService() application.ComponentAfterService {
	return inst
}

func (inst *myBarFactory) Inject(instance application.ComponentInstance, context application.InstanceContext) error {

	o2, err := inst.Cast(instance)
	if err != nil {
		return err
	}

	o2.Owner, err = inst.getterForOwner("#foo1", context)
	if err != nil {
		return err
	}

	o2.Name, err = inst.getterForName("${name}", context)
	if err != nil {
		return err
	}

	return nil
}

func (inst *myBarFactory) getterForOwner(selector string, ic application.InstanceContext) (*Foo, error) {

	o1, err := ic.GetComponent(selector)
	if err != nil {
		return nil, err
	}

	o2, ok := o1.(*Foo)
	if ok {
		return o2, nil
	}

	return nil, errors.New("cannot cast ... etc")
}

func (inst *myBarFactory) getterForName(selector string, ic application.InstanceContext) (string, error) {
	return ic.GetString(selector)
}

func (inst *myBarFactory) Init(instance application.ComponentInstance) error {
	o2, err := inst.Cast(instance)
	if err != nil {
		return err
	}
	return o2.Start()
}

func (inst *myBarFactory) Destroy(instance application.ComponentInstance) error {
	o2, err := inst.Cast(instance)
	if err != nil {
		return err
	}
	return o2.Stop()
}

////////////////////////////////////////////////////////////////////////////////

func manualConfig(cb application.ConfigBuilder) error {

	dp := cb.DefaultProperties()
	dp.SetProperty("abc", "123")
	dp.SetProperty("name", "starter/src/test/go")

	//////////////////////////////

	cibuilder := config.ComInfo()
	var err error = nil

	// com1
	cibuilder = cibuilder.Next()
	cibuilder.ID("foo1").Class("foo").Aliases("").Scope("")
	cibuilder.Factory(&myFooFactory{})
	err = cibuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// com2
	cibuilder = cibuilder.Next()
	cibuilder.ID("com2").Class("bar").Aliases("").Scope("")
	cibuilder.Factory(&myBarFactory{})
	err = cibuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	return nil
}
