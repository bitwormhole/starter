package main

import (
	"embed"

	"github.com/bitwormhole/starter/application/config"
	"github.com/bitwormhole/starter/demo"
)

//go:embed src/main/resources
var resources embed.FS

func main() {
	cb := config.NewBuilderFS(&resources, "src/main/resources")
	err := demo.Run(cb)
	if err != nil {
		panic(err)
	}
}
