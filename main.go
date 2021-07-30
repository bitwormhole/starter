package main

import (
	"embed"

	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/application/config"
	"github.com/bitwormhole/starter/demo"
)

const resourcesBasePath = "src/main/resources"

//go:embed src/main/resources
var resources embed.FS

func main() {
	cfg := config.NewBuilderFS(&resources, resourcesBasePath)
	demo.Config(cfg)
	_, err := application.RunAndLoop(cfg.Create())
	if err != nil {
		panic(err)
	}
}
