package main

import (
	"embed"

	"github.com/bitwormhole/starter"
	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/collection"
	"github.com/bitwormhole/starter/vlog"
)

//go:embed res
var theResources embed.FS

func myMod() application.Module {
	mb := &application.ModuleBuilder{}
	mb.Name("starter/src/test/go").Version("v0.0.0").Revision(0)
	mb.OnMount(func(cb application.ConfigBuilder) error { return manualConfig(cb) })
	mb.Resources(collection.LoadEmbedResources(&theResources, "res"))
	mb.Dependency(starter.Module())
	return mb.Create()
}

func main() {

	vlog.UseSimpleLogger(vlog.ALL)
	vlog.Debug("src/test/go#main.go")

	starter.InitApp().Use(myMod()).Run()
}
