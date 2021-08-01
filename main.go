package main

import (
	"embed"

	"github.com/bitwormhole/starter/starter"
)

//go:embed src/main/resources
var resources embed.FS

func main() {
	starter.Init().EmbedResources(&resources, "src/main/resources").Run()
}
