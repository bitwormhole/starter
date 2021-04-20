package config

import (
	"embed"
	"fmt"
	"os"

	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/application/config"
	"github.com/bitwormhole/starter/demo/demo-for-config/conf"
)

func Demo(fs *embed.FS, prefix string) error {

	config := &config.AppConfig{}
	config.SetResources(fs, prefix)

	// Config(config)
	conf.DefaultConfig(config)

	context, err := application.Run(config, os.Args)
	if err != nil {
		panic(err)
	}

	code := application.Exit(context)
	fmt.Println("exit.code=", code)
	return nil
}
