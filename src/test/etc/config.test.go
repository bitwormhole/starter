package etc

import (
	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/starter/src/test/tester"
)

type theContextPropertiesTester struct {
	markup.Component
	instance *tester.ContextPropertiesTester `initMethod:"Run"`

	AppContext application.Context `inject:"context"`
}

type theContextResourcesTester struct {
	markup.Component
	instance *tester.ContextResourcesTester `initMethod:"Run"`

	AppContext application.Context `inject:"context"`
}
