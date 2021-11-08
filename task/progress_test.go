package task

import (
	"testing"

	"github.com/bitwormhole/starter/contexts"
	"github.com/bitwormhole/starter/tests"
)

func TestGetReporter(t *testing.T) {
	err := doTestGetReporter(t)
	if err != nil {
		t.Error(err)
	}
}

func doTestGetReporter(t *testing.T) error {

	i := tests.Starter(t)
	rt, err := i.RunEx()
	if err != nil {
		return err
	}

	ctx := rt.Context()
	err = contexts.SetupApplicationContext(ctx)
	if err != nil {
		return err
	}

	h, err := GetProgressReporterHolder(ctx)
	if err != nil {
		return err
	}

	h.SetFactory(&DefaultProgressReporterFactory{})

	reporter, err := GetProgressReporter(ctx)
	if err != nil {
		return err
	}

	p := &Progress{}
	p.Name = "demo"
	p.Title = "Demo-1"
	p.TaskID = "demo1"
	p.ValueMax = 10
	p.ValueMin = 0
	p.Unit = "step"
	for cnt := p.ValueMax; cnt > p.ValueMin; cnt-- {
		p.Value = cnt
		reporter.Report(p)
	}
	p.State = StateStopped
	p.Status = StatusOK
	reporter.Update(p)

	return nil
}
