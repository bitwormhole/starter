package handlers

import "github.com/bitwormhole/starter/cli"

type LS struct {
}

func (inst *LS) _Impl() cli.Handler {
	return inst
}

func (inst *LS) Init(service cli.Service) error {
	return service.RegisterHandler("ls", inst)
}

func (inst *LS) Handle(t *cli.TaskContext) error {
	return nil
}
