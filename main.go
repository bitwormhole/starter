package main

import (
	"github.com/bitwormhole/starter/starter"
)

// //go:embed src/main/resources
// var resources embed.FS

func main() {
	starter.InitApp().Use(starter.Module()).Run()
}
