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
