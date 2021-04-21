package demo4gorm

import (
	"embed"
	"fmt"
	"os"

	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/application/config"
)

func Run(fs *embed.FS, prefix string) error {

	cfg := config.Builder(fs, prefix)

	context, err := application.Run(cfg.Create(), os.Args)
	if err != nil {
		return err
	}

	code, err := application.Exit(context)
	fmt.Println("exit.code=", code)
	return err
}
