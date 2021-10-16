package config

import (
	"fmt"

	"github.com/bitwormhole/starter/cli"
	"github.com/bitwormhole/starter/markup"
)

type DemoHandler struct {
	markup.Component `class:"cli-handler"`
}

func (inst *DemoHandler) _Impl() cli.Handler {
	return inst
}

func (inst *DemoHandler) Init(service cli.Service) error {
	service.RegisterHandler("demo", inst)
	return nil
}

func (inst *DemoHandler) Handle(c *cli.TaskContext) error {
	fmt.Println("demo:hello!")
	return nil
}
