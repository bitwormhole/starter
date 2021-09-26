package handlers

import "github.com/bitwormhole/starter/cli"

type Help struct {
}

func (inst *Help) _Impl() cli.Handler {
	return inst
}

func (inst *Help) Init(service cli.Service) error {
	return service.RegisterHandler("help", inst)
}

func (inst *Help) Handle(t *cli.TaskContext) error {
	return nil
}
