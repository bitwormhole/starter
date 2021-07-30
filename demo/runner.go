package demo

import (
	"log"
	"os"

	"github.com/bitwormhole/starter/application"
)

func Run(cb application.ConfigBuilder) error {

	Config(cb)

	code, err := application.RunAndLoop(cb.Create())
	if err != nil {
		return err
	}

	log.Println("Exit with code: ", code)
	os.Exit(code)
	return nil
}
