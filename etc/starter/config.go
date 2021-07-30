package starter

import (
	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/starter/util/configenchecker"
)

func Config(cb application.ConfigBuilder) error {

	dp := cb.DefaultProperties()
	dp.SetProperty("configen.checker.enable", "false")
	dp.SetProperty("test.enable", "false")
	dp.SetProperty("debug.enable", "false")

	return autoGenConfig(cb)
}

type theConfigenChecker struct {
	markup.Component
	instance *configenchecker.ConfigenChecker `initMethod:"Check"`
	Context  application.Context              `inject:"context"`
	Enable   bool                             `inject:"${configen.checker.enable}"`
}
