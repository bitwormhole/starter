package main

import (
	"embed"
	"log"
	"os"

	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/application/config"

	"github.com/bitwormhole/starter/demo/index"
)

//go:embed src/main/resources
var resources embed.FS

func runMainRunner(context application.RuntimeContext) error {

	// get  '${starter.main.runner}:string'
	const key = "starter.main.runner"
	runnerId, err := context.GetProperties().GetPropertyRequired(key)
	if err != nil {
		return err
	}

	com, err := context.GetComponents().GetComponent(runnerId)
	if err != nil {
		return err
	}

	return com.(application.Runnable).Run(context)
}

func tryMain1() error {

	cfg := config.NewBuilderFS(&resources, "src/main/resources")
	index.Config(cfg)

	context, err := application.Run(cfg.Create(), os.Args)
	if err != nil {
		return err
	}

	err = runMainRunner(context)
	if err != nil {
		return err
	}

	code, err := application.Exit(context)
	log.Println("Exit.code:", code)
	return err
}

func main() {

	err := tryMain1()
	if err != nil {
		panic(err)
	}

	// fmt.Println("hello, starter")
}
