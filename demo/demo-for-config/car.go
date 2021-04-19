package config

import (
	"log"

	"github.com/bitwormhole/starter/application"
)

// Car class
type Car struct {
	id      string
	model   string
	context application.Context

	driver *Driver
	engine *Engine

	wheelFrontLeft  *Wheel
	wheelFrontRight *Wheel
	wheelBackLeft   *Wheel
	wheelBackRight  *Wheel
}

func (inst *Car) start() error {

	log.Output(0, "Car.start: "+inst.id)

	return nil
}

func (inst *Car) stop() error {

	log.Output(0, "Car.stop: "+inst.id)

	return nil
}
