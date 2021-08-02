package starter

import (
	"github.com/bitwormhole/starter/application"
)

// Module 函数用于导出本模块
func Module() application.Module {
	return &application.DefineModule{
		Name:     "github.com/bitwormhole/starter",
		Version:  "1.0",
		Revision: 1,
		OnMount:  func(cb application.ConfigBuilder) error { return configure(cb) },
	}
}
