package tester

import (
	"fmt"

	"github.com/bitwormhole/starter/application"
)

type ContextResourcesTester struct {
	AppContext application.Context
	Enable     bool
}

func (inst *ContextResourcesTester) Run() error {

	if !inst.Enable {
		return nil
	}

	res := inst.AppContext.GetResources()
	all := res.All()

	for index := range all {
		name := all[index].AbsolutePath
		text, err := res.GetText(name)
		if err != nil {
			return err
		}
		fmt.Println("load resource [", name, "]")
		fmt.Println("content: [", text, "]")
	}

	return nil
}
