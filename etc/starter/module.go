package starter

import (
	"github.com/bitwormhole/starter/application"
	starterconf "github.com/bitwormhole/starter/etc/starter/starter.conf"
)

const (
	myVersion  = "v0.0.26"
	myRevision = 26
)

// Module 函数用于导出本模块
func Module() application.Module {
	return &application.DefineModule{
		Name:     "github.com/bitwormhole/starter",
		Version:  myVersion,
		Revision: myRevision,
		OnMount:  func(cb application.ConfigBuilder) error { return starterconf.ExportConfig(cb, myVersion, myRevision) },
	}
}
