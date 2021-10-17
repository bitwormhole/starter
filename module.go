package starter

import (
	"github.com/bitwormhole/starter/application"
	etcstarter "github.com/bitwormhole/starter/etc/starter"
	srcmain "github.com/bitwormhole/starter/src/main"
)

const (
	myName     = "github.com/bitwormhole/starter"
	myVersion  = "v0.0.62"
	myRevision = 62
)

// Module 函数用于导出本模块
func Module() application.Module {

	builder := &application.ModuleBuilder{}
	builder.Name(myName).Version(myVersion).Revision(myRevision)
	builder.Resources(srcmain.ExportResources())
	builder.Dependency(nil)

	mod := builder.Create()
	builder.OnMount(func(cb application.ConfigBuilder) error { return etcstarter.ExportConfig(cb, mod) })
	return builder.Create()
}
