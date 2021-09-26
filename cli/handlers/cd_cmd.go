package handlers

import "github.com/bitwormhole/starter/cli"

type CD struct {
}

func (inst *CD) _Impl() cli.Handler {
	return inst
}

func (inst *CD) Init(service cli.Service) error {
	return service.RegisterHandler("cd", inst)
}

func (inst *CD) Handle(t *cli.TaskContext) error {
	return nil
}
