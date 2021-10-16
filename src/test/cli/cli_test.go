package main

import (
	"testing"

	"github.com/bitwormhole/starter/cli"
	"github.com/bitwormhole/starter/src/test/cli/mod"
	"github.com/bitwormhole/starter/tests"
)

func TestCLI(t *testing.T) {
	err := doTestCLI(t)
	if err != nil {
		t.Error(err)
	}
}

func doTestCLI(t *testing.T) error {

	i := tests.Starter(t)
	i.Use(mod.ExportCLITestModule())

	rt, err := i.RunEx()
	if err != nil {
		return err
	}

	ctx := rt.Context()
	o1, err := ctx.GetComponent("#cli-service")
	if err != nil {
		return err
	}

	factory := o1.(cli.Service).GetClientFactory()
	client := factory.CreateClient(ctx)
	err = client.ExecuteScript("demo")
	if err != nil {
		return err
	}

	return rt.Exit()
}
