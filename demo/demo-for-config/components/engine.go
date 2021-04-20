package components

import "log"

// Engine class
type Engine struct {
	Name  string
	Owner *Car
}

func (inst *Engine) Start() error {

	log.Output(0, "Engine.start: "+inst.Name)

	return nil
}

func (inst *Engine) Stop() error {

	log.Output(0, "Engine.stop: "+inst.Name)

	return nil
}
