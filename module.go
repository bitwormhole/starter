package starter

import (
	"github.com/bitwormhole/starter/application"
	etcstarter "github.com/bitwormhole/starter/etc/starter"
	srcmain "github.com/bitwormhole/starter/src/main"
)

const (
	myVersion  = "v0.0.31"
	myRevision = 31
)

// Module 函数用于导出本模块
func Module() application.Module {

	mod := &application.DefineModule{
		Name:     "github.com/bitwormhole/starter",
		Version:  myVersion,
		Revision: myRevision,
	}

	mod.OnMount = func(cb application.ConfigBuilder) error { return etcstarter.ExportConfig(cb, mod) }
	mod.Resources = srcmain.ExportResources()

	return mod
}
