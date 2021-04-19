package main

import (
	"embed"
	"fmt"
	"os"

	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/application/config"

	demo "github.com/bitwormhole/starter/demo/demo-for-config"
)

//go:embed src/main/resources
var resources embed.FS

func tryMain() error {

	config := &config.AppConfig{}
	// fsys := fs.Default()
	// roots := fsys.Roots()
	args := os.Args

	config.SetResources(&resources, "src/main/resources")
	demo.Config(config)

	context, err := application.Run(config, args)
	if err != nil {
		return err
	}

	context.GetComponents().GetComponent("seby")
	context.GetComponents().GetComponent("car-x")
	context.GetComponents().GetComponent("car-y")
	context.GetComponents().GetComponent("car-z")

	fmt.Println("components.names(include aliases):")
	namelist := context.GetComponents().GetComponentNameList(true)
	for index := range namelist {
		fmt.Println("    " + namelist[index])
	}

	fmt.Println("components.names(without aliases):")
	namelist = context.GetComponents().GetComponentNameList(false)
	for index := range namelist {
		fmt.Println("    " + namelist[index])
	}

	code := application.Exit(context)
	fmt.Println("exited, code=", code)
	// fmt.Println("  file.roots=", roots)

	return nil
}

func main() {
	err := tryMain()
	if err != nil {
		panic(err)
	}
}
