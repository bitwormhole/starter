package starter

import (
	"embed"

	"github.com/bitwormhole/starter/application/config"
	"github.com/bitwormhole/starter/collection"
)

//go:embed src/main/resources
var resources embed.FS

func GetResources() collection.Resources {
	const basePath = "src/main/resources"
	return config.CreateEmbedFsResources(&resources, basePath)
}
