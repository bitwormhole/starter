package tester

import (
	"fmt"

	"github.com/bitwormhole/starter/application"
)

type ContextResourcesTester struct {
	AppContext application.Context
}

func (inst *ContextResourcesTester) Run() error {

	res := inst.AppContext.GetResources()
	all := res.All()

	for index := range all {
		name := all[index]
		text, err := res.GetText(name)
		if err != nil {
			return err
		}
		fmt.Println("load resource [", name, "]")
		fmt.Println("content: [", text, "]")
	}

	return nil
}
