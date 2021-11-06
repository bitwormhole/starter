package main

import (
	"testing"
	"time"

	"github.com/bitwormhole/starter"
	"github.com/bitwormhole/starter/collection"
	"github.com/bitwormhole/starter/tests"
	"github.com/bitwormhole/starter/vlog"
)

func TestVlog(t *testing.T) {
	err := doTestVlog(t)
	if err != nil {
		t.Error(err)
	}
}

func doTestVlog(t *testing.T) error {

	props := collection.CreateProperties()
	props.SetProperty("vlog.file.enable", "true")

	i := tests.Starter(t)
	i.UseMain(starter.Module())
	i.UseProperties(props)

	rt, err := i.RunEx()
	if err != nil {
		return err
	}

	err = rt.Loop()
	if err != nil {
		return err
	}

	const total = 5

	for ttl := total; ttl > 0; ttl-- {
		vlog.Trace("test for vlog")
		vlog.Debug("test for vlog")
		vlog.Info("test for vlog")
		vlog.Warn("test for vlog")
		vlog.Error("test for vlog")
		vlog.Fatal("test for vlog")
		time.Sleep(time.Second)
	}

	return rt.Exit()
}
