package demo

import (
	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/demo/elements"
	"github.com/bitwormhole/starter/markup"
)

////////////////////////////////////////////////////////////////////////////////

type exampleComponent1 struct {
	markup.Component `id:"com1" class:"com1"`

	instance *elements.ComExample1
	context  application.Context

	Name string `inject:"${com1.name}" default:"cc1"`
}

type exampleComponent2 struct {
	markup.Component `id:"com2" class:"com2  Looper"`

	instance *elements.ComExample2 `initMethod:"Open" destroyMethod:"Close"`
	context  application.Context

	Com1ref *elements.ComExample1 `inject:"#com1"`
}

type loop1 struct {
	markup.Component `class:"looper"`

	instance *elements.Looper1
	context  application.Context
}

type loop2 struct {
	markup.Component `class:"looper"`

	instance *elements.Looper2
	context  application.Context
}

type exit1 struct {
	markup.Component `class:"exit-code-generator"`

	instance *elements.ExitCoder1
	context  application.Context
}

type exit2 struct {
	markup.Component `class:"exit-code-generator"`

	instance *elements.ExitCoder2
	context  application.Context
}
