package gen

import (
	"strconv"

	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/starter/util/configenchecker"
)

// ExportConfigForStarter 对外导出配置
func ExportConfigForStarter(cb application.ConfigBuilder, module application.Module) error {

	dp := cb.DefaultProperties()

	//  dp.SetProperty("starter.version", starterVersion)
	//  dp.SetProperty("starter.revision", strconv.Itoa(starterRevision))
	dp.SetProperty("module.starter.name", module.GetName())
	dp.SetProperty("module.starter.version", module.GetVersion())
	dp.SetProperty("module.starter.revision", strconv.Itoa(module.GetRevision()))

	return autoGenConfig(cb)
}

type theConfigenChecker struct {
	markup.Component
	instance *configenchecker.ConfigenChecker `initMethod:"Check"`
	Context  application.Context              `inject:"context"`
	Enable   bool                             `inject:"${configen.checker.enable}"`
}
