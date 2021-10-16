package mod

import (
	"github.com/bitwormhole/starter"
	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/modules/cli/mod/gen"
)

// ModuleCLI 导出CLI的核心模块
func ModuleCLI() application.Module {

	parent := starter.Module()

	mb := application.ModuleBuilder{}
	mb.Name(parent.GetName() + "#cli")
	mb.Version(parent.GetVersion())
	mb.Revision(parent.GetRevision())
	mb.OnMount(gen.ExportConfigCLI)

	mb.Dependency(parent)

	return mb.Create()
}
