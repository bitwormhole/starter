package tests

import (
	"testing"

	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/io/fs"
)

////////////////////////////////////////////////////////////////////////////////

type innerRuntimeWrapperBuilder struct {
	inner              application.Runtime
	t                  *testing.T
	testingDataSrcName string // the name for resources
}

func (inst *innerRuntimeWrapperBuilder) Create() *innerRuntimeWrapper {
	wrapper := &innerRuntimeWrapper{
		t:     inst.t,
		inner: inst.inner,
	}
	wrapper.testingDataDir = inst.loadTestingDataDir(wrapper)
	return wrapper
}

func (inst *innerRuntimeWrapperBuilder) Wrap() TestingRuntime {
	return inst.Create()
}

func (inst *innerRuntimeWrapperBuilder) loadTestingDataDir(rt TestingRuntime) fs.Path {
	name := inst.testingDataSrcName
	if name == "" {
		return nil
	}
	loader := &defaultTestingDataLoader{}
	dir, err := loader.Load(rt, name, nil)
	if err != nil {
		panic(err)
	}
	return dir
}

////////////////////////////////////////////////////////////////////////////////

type innerRuntimeWrapper struct {
	inner          application.Runtime
	t              *testing.T
	testingDataDir fs.Path // cached
}

func (inst *innerRuntimeWrapper) _Impl() TestingRuntime {
	return inst
}

/// delegate

func (inst *innerRuntimeWrapper) Context() application.Context {
	return inst.inner.Context()
}

func (inst *innerRuntimeWrapper) Loop() error {
	return inst.inner.Loop()
}

func (inst *innerRuntimeWrapper) Exit() error {
	return inst.inner.Exit()
}

/// extends

func (inst *innerRuntimeWrapper) T() *testing.T {
	return inst.t
}

func (inst *innerRuntimeWrapper) TestingDataDir() fs.Path {
	dir := inst.testingDataDir
	if dir == nil {
		panic("no testing data")
	}
	return dir
}
