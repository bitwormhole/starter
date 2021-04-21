package demo4gin

import (
	"fmt"
	"os"

	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/application/config"
	"github.com/bitwormhole/starter/web/gin_starter"
)

type Runner struct {
}

func (inst *Runner) Run(context application.RuntimeContext) error {

	// 这里新建了一个内嵌的context

	cfg := config.NewBuilder()
	cfg.SetResources(context.GetResources())
	gin_starter.Config(cfg)

	context, err := application.Run(cfg.Create(), os.Args)
	if err != nil {
		return err
	}

	err = gin_starter.Run(context)
	if err != nil {
		return err
	}

	code, err := application.Exit(context)
	fmt.Println("exit.code=", code)
	return err
}
