package config

import (
	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/application/config"
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
			inst := obj.(*Driver)

			inst.name = "seby"
			inst.car = helper.getCar("car-y")
			inst.birthday = "1999-09-10"
			inst.sex = "female"

			inst.car.driver = inst

			return helper.err
		},

		OnNew: func() lang.Object {
			return &Driver{}
		},
	})

	cfg.AddComponent(&config.ComInfo{
		ID:    "car-x",
		Class: "car car-model-x",
		Scope: application.ScopePrototype,

		OnNew: func() lang.Object {
			return &Car{}
		},

		OnInit: func(obj lang.Object) error {
			car := obj.(*Car)
			return car.start()
		},

		OnDestroy: func(obj lang.Object) error {
			car := obj.(*Car)
			return car.stop()
		},

		OnInject: func(obj lang.Object, context application.RuntimeContext) error {

			car := obj.(*Car)
			helper := &injectHelper{context: context}

			car.context = context
			car.engine = nil
			car.model = "X"

			return helper.err
		},
	})

	cfg.AddComponent(&config.ComInfo{
		ID:    "car-y",
		Class: "car car-model-y",
		Scope: application.ScopePrototype,

		OnNew: func() lang.Object {
			return &Car{}
		},

		OnInit: func(obj lang.Object) error {
			car := obj.(*Car)
			return car.start()
		},

		OnDestroy: func(obj lang.Object) error {
			car := obj.(*Car)
			return car.stop()
		},

		OnInject: func(obj lang.Object, context application.RuntimeContext) error {

			helper := &injectHelper{context: context}
			car := obj.(*Car)

			car.context = context
			car.engine = helper.getEngine("car-y-engine")
			car.model = "Y"
			car.id = "GC17258"

			return helper.err
		},
	})

	cfg.AddComponent(&config.ComInfo{
		ID:    "car-y-engine",
		Class: "engine engine-y",
		Scope: application.ScopePrototype,

		OnNew: func() lang.Object {
			return &Engine{}
		},

		OnInit: func(obj lang.Object) error {
			engine := obj.(*Engine)
			return engine.start()
		},

		OnDestroy: func(obj lang.Object) error {
			engine := obj.(*Engine)
			return engine.stop()
		},

		OnInject: func(obj lang.Object, context application.RuntimeContext) error {

			helper := &injectHelper{context: context}
			engine := obj.(*Engine)

			engine.owner = helper.getCar("car-y")
			engine.name = "car-y-engine"

			return helper.err
		},
	})

}
