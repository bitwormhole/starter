package main

import (
	"testing"

	"github.com/bitwormhole/starter/tests"
	"github.com/bitwormhole/starter/vlog"
)

func TestMain(t *testing.T) {

	//	appinit := starter.InitApp()
	vlog.Debug("src/test/go#main_test.go")

	appinit := tests.TestingStarter(t)
	ch, err := appinit.RunEx()
	if err != nil {
		t.Error(err)
	}

	ctx := ch.Context()
	nihao, err := ctx.GetComponent("#nihao")
	if err != nil {
		//t.Error(err)
	}

	vlog.Info(nihao)

	// ch.Loop()
	// ch.Exit()
}
