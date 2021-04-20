package components

import (
	"log"

	"github.com/bitwormhole/starter/application"
)

// Car class
type Car struct {
	Id      string
	Model   string
	Context application.Context

	Driver *Driver
	Engine *Engine

	WheelFrontLeft  *Wheel
	WheelFrontRight *Wheel
	WheelBackLeft   *Wheel
	WheelBackRight  *Wheel
}

func (inst *Car) Start() error {

	log.Output(0, "Car.start: "+inst.Id)

	return nil
}

func (inst *Car) Stop() error {

	log.Output(0, "Car.stop: "+inst.Id)

	return nil
}
