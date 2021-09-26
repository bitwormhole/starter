package handlers

import "github.com/bitwormhole/starter/cli"

type PWD struct {
}

func (inst *PWD) _Impl() cli.Handler {
	return inst
}

func (inst *PWD) Init(service cli.Service) error {
	return service.RegisterHandler("pwd", inst)
}

func (inst *PWD) Handle(t *cli.TaskContext) error {
	return nil
}
