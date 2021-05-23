package elements

import (
	"fmt"

	"github.com/bitwormhole/starter/application"
)

////////////////////////////////////////////////////////////////////////////////

type Looper1 struct {
}

func (inst *Looper1) _api_to_looper() application.Looper {
	return inst
}

func (inst *Looper1) Loop() error {
	fmt.Println("run loop1()")
	return nil
}

func (inst *Looper1) Priority() int {
	return 111
}

////////////////////////////////////////////////////////////////////////////////

type Looper2 struct {
}

func (inst *Looper2) _api_to_looper() application.Looper {
	return inst
}

func (inst *Looper2) Loop() error {
	fmt.Println("run loop2()")
	return nil
}

func (inst *Looper2) Priority() int {
	return 22
}
