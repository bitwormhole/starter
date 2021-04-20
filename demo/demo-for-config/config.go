package config

import (
	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/application/config"
	"github.com/bitwormhole/starter/demo/demo-for-config/components"
	"github.com/bitwormhole/starter/lang"
)

// Config config the app
func Config(cfg application.ConfigBuilder) {

	cfg.AddComponent(&config.ComInfo{
		ID:      "seby",
		Class:   "driver",
		Aliases: []string{"badgirl", "faster", "car-y-owner"},

		OnInject: func(obj lang.Object, context application.RuntimeContext) error {

			helper := &injectHelper{context: context}
			inst := obj.(*components.Driver)

			inst.Name = "seby"
			inst.Car = helper.getCar("car-y")
			inst.Birthday = "1999-09-10"
			inst.Sex = "female"

			inst.Car.Driver = inst

			return helper.err
		},

		OnNew: func() lang.Object {
			return &components.Driver{}
		},
	})

	cfg.AddComponent(&config.ComInfo{
		ID:    "car-x",
		Class: "car car-model-x",
		Scope: application.ScopePrototype,

		OnNew: func() lang.Object {
			return &components.Car{}
		},

		OnInit: func(obj lang.Object) error {
			car := obj.(*components.Car)
			return car.Start()
		},

		OnDestroy: func(obj lang.Object) error {
			car := obj.(*components.Car)
			return car.Stop()
		},

		OnInject: func(obj lang.Object, context application.RuntimeContext) error {

			car := obj.(*components.Car)
			helper := &injectHelper{context: context}

			car.Context = context
			car.Engine = nil
			car.Model = "X"

			return helper.err
		},
	})

	cfg.AddComponent(&config.ComInfo{
		ID:    "car-y",
		Class: "car car-model-y",
		Scope: application.ScopePrototype,

		OnNew: func() lang.Object {
			return &components.Car{}
		},

		OnInit: func(obj lang.Object) error {
			car := obj.(*components.Car)
			return car.Start()
		},

		OnDestroy: func(obj lang.Object) error {
			car := obj.(*components.Car)
			return car.Stop()
		},

		OnInject: func(obj lang.Object, context application.RuntimeContext) error {

			helper := &injectHelper{context: context}
			car := obj.(*components.Car)

			car.Context = context
			car.Engine = helper.getEngine("car-y-engine")
			car.Model = "Y"
			car.Id = "GC17258"

			return helper.err
		},
	})

	cfg.AddComponent(&config.ComInfo{
		ID:    "car-y-engine",
		Class: "engine engine-y",
		Scope: application.ScopePrototype,

		OnNew: func() lang.Object {
			return &components.Engine{}
		},

		OnInit: func(obj lang.Object) error {
			engine := obj.(*components.Engine)
			return engine.Start()
		},

		OnDestroy: func(obj lang.Object) error {
			engine := obj.(*components.Engine)
			return engine.Stop()
		},

		OnInject: func(obj lang.Object, context application.RuntimeContext) error {

			helper := &injectHelper{context: context}
			engine := obj.(*components.Engine)

			engine.Owner = helper.getCar("car-y")
			engine.Name = "car-y-engine"

			return helper.err
		},
	})

}
