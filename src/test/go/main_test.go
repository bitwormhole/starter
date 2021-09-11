package main

import (
	"testing"

	"github.com/bitwormhole/starter/src/test"
	"github.com/bitwormhole/starter/tests"
	"github.com/bitwormhole/starter/vlog"
)

func TestMain(t *testing.T) {

	//	appinit := starter.InitApp()
	vlog.Debug("src/test/go#main_test.go")

	i := tests.Starter(t)
	i.UsePanic()
	i.UseResources(test.ExportResources())
	rt, _ := i.RunEx()

	ctx := rt.Context()
	nihao, err := ctx.GetComponent("#nihao")
	if err != nil {
		//t.Error(err)
	}
	vlog.Info(nihao)

	rt.Loop()
	rt.Exit()
}
