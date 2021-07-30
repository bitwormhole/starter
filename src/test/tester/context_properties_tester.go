package tester

import (
	"fmt"
	"sort"

	"github.com/bitwormhole/starter/application"
)

type ContextPropertiesTester struct {
	AppContext application.Context
	Enable     bool
}

func (inst *ContextPropertiesTester) Run() error {

	if !inst.Enable {
		return nil
	}

	app := inst.AppContext
	props := app.GetProperties()
	table := props.Export(nil)

	keys := make([]string, 0)
	for key := range table {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	for index := range keys {
		key := keys[index]
		val := table[key]
		fmt.Println("app.property[", key, "] = [", val, "]")
	}

	return nil
}
