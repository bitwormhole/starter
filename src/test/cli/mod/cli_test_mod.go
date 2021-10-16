package mod

import (
	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/modules/cli/mod"
	"github.com/bitwormhole/starter/src/test/cli/mod/gen"
)

func ExportCLITestModule() application.Module {
	mb := application.ModuleBuilder{}
	mb.Name("src/test/cli/mod").Version("v1").Revision(1)
	mb.OnMount(gen.ExportConfigCLITest)

	mb.Dependency(mod.ModuleCLI())

	return mb.Create()
}
