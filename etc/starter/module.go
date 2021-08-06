package starter

import (
	"github.com/bitwormhole/starter/application"
)

// Module 函数用于导出本模块
func Module() application.Module {
	return &application.DefineModule{
		Name:     "github.com/bitwormhole/starter",
		Version:  StarterVersion,
		Revision: StarterRevision,
		OnMount:  func(cb application.ConfigBuilder) error { return configure(cb, StarterVersion, StarterRevision) },
	}
}
