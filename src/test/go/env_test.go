package main

import (
	"fmt"
	"testing"

	"github.com/bitwormhole/starter/tests"
	"github.com/bitwormhole/starter/vlog"
)

func TestEnv(t *testing.T) {

	ts := tests.TestingStarter(t)
	ts.UsePanic()
	rt, _ := ts.RunEx()
	env := rt.Context().GetEnvironment()
	all := env.Export(nil)

	fmt.Println("os.Environ():")

	str, err := env.GetEnv("ni-hao")
	if err == nil {
		vlog.Debug(str)
	} else {
		rt.Context().GetErrorHandler().HandleError(err)
	}

	for k, v := range all {
		fmt.Println("\t", k, " = ", v)
	}
}
