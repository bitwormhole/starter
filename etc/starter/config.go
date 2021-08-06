package starter

import (
	"strconv"

	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/starter/util/configenchecker"
)

func configure(cb application.ConfigBuilder, starterVersion string, starterRevision int) error {

	dp := cb.DefaultProperties()
	dp.SetProperty("configen.checker.enable", "false")
	dp.SetProperty("test.enable", "false")
	dp.SetProperty("debug.enable", "false")
	dp.SetProperty("starter.version", starterVersion)
	dp.SetProperty("starter.revision", strconv.Itoa(starterRevision))

	return autoGenConfig(cb)
}

type theConfigenChecker struct {
	markup.Component
	instance *configenchecker.ConfigenChecker `initMethod:"Check"`
	Context  application.Context              `inject:"context"`
	Enable   bool                             `inject:"${configen.checker.enable}"`
}
