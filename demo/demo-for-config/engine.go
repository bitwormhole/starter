package config

import "log"

// Engine class
type Engine struct {
	name  string
	owner *Car
}

func (inst *Engine) start() error {

	log.Output(0, "Engine.start: "+inst.name)

	return nil
}

func (inst *Engine) stop() error {

	log.Output(0, "Engine.stop: "+inst.name)

	return nil
}
