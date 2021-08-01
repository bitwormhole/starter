package starter

import (
	"embed"

	"github.com/bitwormhole/starter/application/config"
	"github.com/bitwormhole/starter/collection"
)

//go:embed src/main/resources
var resources embed.FS

func GetResources() collection.Resources {
	return config.CreateEmbedFsResources(&resources, "src/main/resources")
}
