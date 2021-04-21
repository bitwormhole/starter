package conf

import (
	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/demo/demo-for-config/components"
)

// func DefaultConfig(cfg application.ConfigBuilder) {}

func injectCar1(t *components.Car, context application.RuntimeContext) error {

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

type injectDriver2dep interface {
	getCar(sel string) *components.Car
}

func injectDriver2(t *components.Driver, dep injectDriver2dep) error {

	// [component]
	// class=driver

	car1 := dep.getCar("#car1")
	car2 := dep.getCar("#car2")
	list := []*components.Car{car1, car2}

	t.MyCars = list
	t.Name = "seby"

	return nil
}

func injectDriver3(t *components.Driver, dep interface {
	getCar(sel string) *components.Car
}) error {

	// [component]
	// class=driver

	t.Car = dep.getCar("#car4")

	return nil
}
