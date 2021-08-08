package starter

import (
	"github.com/bitwormhole/starter/application"
	starterconf "github.com/bitwormhole/starter/etc/starter/starter.conf"
	srcmain "github.com/bitwormhole/starter/src/main"
)

const (
	myVersion  = "v0.0.27"
	myRevision = 27
)

// ExportModule 函数用于导出本模块
func ExportModule() application.Module {

	res := srcmain.ExportResources()

	return &application.DefineModule{
		Name:      "github.com/bitwormhole/starter",
		Version:   myVersion,
		Revision:  myRevision,
		Resources: res,
		OnMount:   func(cb application.ConfigBuilder) error { return starterconf.ExportConfig(cb, myVersion, myRevision) },
	}
}
