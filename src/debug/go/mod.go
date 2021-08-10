package main

import (
	"github.com/bitwormhole/starter"
	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/src/debug/go/conf"
)

func innerModule() application.Module {

	mod := &application.DefineModule{
		Name:     "github.com/bitwormhole/starter/+debug",
		Version:  "0.1",
		Revision: 1,
	}

	mod.OnMount = func(cb application.ConfigBuilder) error { return conf.ExportConfig(cb) }
	mod.AddDependency(starter.Module())

	return mod
}
