package conf

import (
	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/demo/demo-for-config/components"
)

// func DefaultConfig(cfg application.ConfigBuilder) {}

/****
word
*/

func injectCar1(t *components.Car, context application.RuntimeContext) error {

	/****
	  hi
	*/

	// [component]
	// id=
	// aliases=
	// class=
	// initMethod=start
	// destroyMethod=stop

	engine, _ := context.GetComponents().GetComponent("engine")
	t.Engine = engine.(*components.Engine)
	t.Id = "car1"

	return nil
}

func injectCar2(t *components.Car, context application.RuntimeContext) error {

	engine, _ := context.GetComponents().GetComponent("engine")
	t.Engine = engine.(*components.Engine)
	t.Id = "car2"

	return nil
}

type injectCar3Dep interface {
	getEngine(sel string) *components.Engine
}

func injectCar3(t *components.Car, dep injectCar3Dep) error {

	// [component]
	// id=car3
	// class=car
	// scope=singleton

	t.Engine = dep.getEngine(".engine")
	t.Id = "car2"
	return nil
}

func injectEngine(t *components.Engine, context application.RuntimeContext) error {

	return nil
}

func injectDriver(t *components.Driver, context application.RuntimeContext) error {

	car1, _ := context.GetComponents().GetComponent("car1")
	car2, _ := context.GetComponents().GetComponent("car2")

	list := []*components.Car{car1.(*components.Car), car2.(*components.Car)}
	t.MyCars = list
	t.Name = "seby"

	return nil
}

////////////////////////////////////////////////////////////////////////////////

type injectDriver2 struct {
	t    *components.Driver `tag:"target" class:"driver"`
	car1 *components.Car    `tag:"var" value:"#car1"`
	car2 *components.Car    `tag:"var" value:"#car1"`
}

func (inst *injectDriver2) inject() error {

	t := inst.t
	list := []*components.Car{inst.car1, inst.car2}

	t.MyCars = list
	t.Name = "seby"

	return nil
}

////////////////////////////////////////////////////////////////////////////////

type injectDriver3 struct {
	target *components.Driver `tag:"target" class:"driver"`
	Car    *components.Car    `value:"#car4"`
}

////////////////////////////////////////////////////////////////////////////////
// style like a func

func styleLikeFunc(t *components.Door, context application.RuntimeContext) error {

	// [component]
	// id=style-like-a-func

	getter := context.NewGetter()
	ok := true

	t.Position = getter.GetPropertyString(ok, "position", "front-left")
	t.Owner, ok = getter.GetComponent(ok, "#car1").(*components.Car)

	return getter.Result(ok)
}

////////////////////////////////////////////////////////////////////////////////
// style like a struct

type xyzInjector struct {

	// [component]
	// id=style-like-a-struct

	target   *components.Door           `tag:"target" id:"1"  class:"xxx aaa b-b-b c_c_c" scope:"singleton"  initMethod:"start" destroyMethod:"stop" `
	context  application.RuntimeContext `tag:"context"`
	Owner    *components.Car            `value:"#abc"`
	Position string                     `value:"${a.b.c}"`
}

func (inst *xyzInjector) inject() error {

	t := inst.target

	t.Owner = inst.Owner
	t.Position = inst.Position

	return nil
}
