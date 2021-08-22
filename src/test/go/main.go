package main

import (
	"github.com/bitwormhole/starter"
	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/vlog"
)

func myMod() application.Module {
	mb := &application.ModuleBuilder{}
	mb.Name("starter/src/test/go").Version("v0.0.0").Revision(0)
	mb.OnMount(func(cb application.ConfigBuilder) error { return manualConfig(cb) })
	return mb.Create()
}

func main() {

	vlog.UseSimpleLogger(vlog.INFO)

	vlog.Debug("src/test/go#main.go")
	starter.InitApp().Use(myMod()).Run()
}
