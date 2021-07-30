package main

import (
	"embed"

	"github.com/bitwormhole/starter/application/config"
	"github.com/bitwormhole/starter/demo"
)

const resourcesBasePath = "src/main/resources"

//go:embed src/main/resources
var resources embed.FS

func main() {
	cb := config.NewBuilderFS(&resources, resourcesBasePath)
	cb.SetEnableLoadPropertiesFromArguments(false)

	err := demo.Run(cb)
	if err != nil {
		panic(err)
	}
}
