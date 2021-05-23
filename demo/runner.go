package demo

import (
	"log"
	"os"

	"github.com/bitwormhole/starter/application"
)

func Run(cb application.ConfigBuilder) error {

	Config(cb)

	configuration := cb.Create()
	context, err := application.Run(configuration, os.Args)
	if err != nil {
		return err
	}

	err = application.Loop(context)
	if err != nil {
		return err
	}

	code, err := application.Exit(context)
	if err != nil {
		return err
	}

	log.Println("Exit with code: ", code)
	return nil
}
