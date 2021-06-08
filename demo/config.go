package demo

import (
	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/application/config"
	"github.com/bitwormhole/starter/demo/elements"
	"github.com/bitwormhole/starter/lang"
)

////////////////////////////////////////////////////////////////////////////////

func Config(cb application.ConfigBuilder) error {

	cb.AddComponent(&config.ComInfo{
		ID: "com1",

		OnNew: func() lang.Object {
			return &elements.ComExample2{}
		},

		OnInit: func(o lang.Object) error {
			o2 := o.(*elements.ComExample2)
			return o2.Open()
		},

		OnDestroy: func(o lang.Object) error {
			o2 := o.(*elements.ComExample2)
			return o2.Close()
		},
	})

	return nil
}
