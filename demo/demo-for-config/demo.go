package config

import (
	"os"

	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/application/config"
)

func Demo() int {

	config := &config.AppConfig{}
	Config(config)

	context, err := application.Run(config, os.Args)
	if err != nil {
		panic(err)
	}

	code := application.Exit(context)
	return code
}
