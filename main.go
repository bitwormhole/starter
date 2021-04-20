package main

import (
	"embed"
	"os"

	"github.com/bitwormhole/starter/docs/help"
	"github.com/bitwormhole/starter/tools/configen"

	demo "github.com/bitwormhole/starter/demo/demo-for-config"
)

//go:embed src/main/resources
var resources embed.FS

func tryMain() error {

	action := ""
	args := os.Args
	if len(args) > 1 {
		action = args[1]
	}

	if action == "demo" {
		return demo.Demo(&resources, "src/main/resources")
	} else if action == "configen" {
		return configen.Main(args)
	} else {
		return help.PrintHelpInfo()
	}
}

func main() {
	err := tryMain()
	if err != nil {
		panic(err)
	}
}
