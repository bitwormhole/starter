package test

import (
	"embed"

	"github.com/bitwormhole/starter/collection"
)

//go:embed resources
var theResFS embed.FS

// ExportResources 导出资源
func ExportResources() collection.Resources {
	return collection.LoadEmbedResources(&theResFS, "resources")
}
