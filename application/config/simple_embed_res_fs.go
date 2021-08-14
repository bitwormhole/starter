package config

import (
	"embed"

	"github.com/bitwormhole/starter/collection"
)

// LoadResourcesFromEmbedFS 从嵌入的FS加载资源组
func LoadResourcesFromEmbedFS(fs *embed.FS, path string) collection.Resources {
	return collection.LoadEmbedResources(fs, path)
}
