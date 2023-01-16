package starter

import (
	"embed"

	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/collection"
	"github.com/bitwormhole/starter/gen"
)

const (
	myName     = "github.com/bitwormhole/starter"
	myVersion  = "v0.1.6"
	myRevision = 89
)

//go:embed src/main/resources
var theMainRes embed.FS

// Module 函数用于导出本模块
func Module() application.Module {

	dm := &application.DefineModule{
		Name:     myName,
		Version:  myVersion,
		Revision: myRevision,
	}

	builder := &application.ModuleBuilder{}
	builder.Name(dm.Name).Version(dm.Version).Revision(dm.Revision)
	builder.Resources(collection.LoadEmbedResources(&theMainRes, "src/main/resources"))
	builder.Dependency(nil)
	builder.OnMount(
		func(cb application.ConfigBuilder) error {
			return gen.ExportConfigForStarter(cb, dm)
		})
	return builder.Create()
}
