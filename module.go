package starter

import (
	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/gen"
	srcmain "github.com/bitwormhole/starter/src/main"
)

const (
	myName     = "github.com/bitwormhole/starter"
	myVersion  = "v0.0.76"
	myRevision = 76
)

// Module 函数用于导出本模块
func Module() application.Module {

	builder := &application.ModuleBuilder{}
	builder.Name(myName).Version(myVersion).Revision(myRevision)
	builder.Resources(srcmain.ExportResources())
	builder.Dependency(nil)
	mod := builder.Create()

	builder.OnMount(func(cb application.ConfigBuilder) error { return gen.ExportConfigForStarter(cb, mod) })
	return builder.Create()
}
