package tests

import (
	"embed"
	"testing"

	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/collection"
	"github.com/bitwormhole/starter/lang"
	"github.com/bitwormhole/starter/vlog"
)

////////////////////////////////////////////////////////////////////////////////

// wrapInitializer 包装  application.Initializer, 为其添加测试的功能
func wrapInitializer(inner application.Initializer, t *testing.T) TestingInitializer {
	return &innerInitializerWrapper{inner: inner, t: t}
}

////////////////////////////////////////////////////////////////////////////////

type innerInitializerWrapper struct {
	inner              application.Initializer
	t                  *testing.T
	testingDataSrcName string
}

func (inst *innerInitializerWrapper) _Impl() TestingInitializer {
	return inst
}

//// delegate

func (inst *innerInitializerWrapper) SetErrorHandler(h lang.ErrorHandler) application.Initializer {
	return inst.inner.SetErrorHandler(h)
}

func (inst *innerInitializerWrapper) SetAttribute(name string, value interface{}) application.Initializer {
	return inst.inner.SetAttribute(name, value)
}

func (inst *innerInitializerWrapper) SetExitEnabled(en bool) application.Initializer {
	inst.inner.SetExitEnabled(en)
	return inst
}

func (inst *innerInitializerWrapper) SetPanicEnabled(enabled bool) application.Initializer {
	inst.inner.SetPanicEnabled(enabled)
	return inst
}

func (inst *innerInitializerWrapper) SetArguments(args []string) application.Initializer {
	inst.inner.SetArguments(args)
	return inst
}

func (inst *innerInitializerWrapper) Use(module application.Module) application.Initializer {
	return inst.inner.Use(module)
}

func (inst *innerInitializerWrapper) UsePanic() application.Initializer {
	return inst.inner.UsePanic()
}

func (inst *innerInitializerWrapper) Run() {
	inst.inner.Run()
}

func (inst *innerInitializerWrapper) RunEx() (application.Runtime, error) {
	return inst.inner.RunEx()
}

func (inst *innerInitializerWrapper) UseResources(res collection.Resources) application.Initializer {
	inst.inner.UseResources(res)
	return inst
}

func (inst *innerInitializerWrapper) UseProperties(p collection.Properties) application.Initializer {
	inst.inner.UseProperties(p)
	return inst
}

//// extends

func (inst *innerInitializerWrapper) LoadPropertisFromGitConfig(required bool) TestingInitializer {
	loader := &testPropertiesInGitLoader{}
	src, err := loader.load()
	if err != nil {
		if required {
			panic(err)
		} else {
			vlog.Warn(err)
		}
	}
	inst.UseProperties(src)
	return inst
}

func (inst *innerInitializerWrapper) T() *testing.T {
	return inst.t
}

func (inst *innerInitializerWrapper) UseResourcesFS(efs *embed.FS, path string) TestingInitializer {
	r := collection.LoadEmbedResources(efs, path)
	inst.UseResources(r)
	return inst
}

func (inst *innerInitializerWrapper) PrepareTestingDataFromResource(name string) TestingInitializer {
	inst.testingDataSrcName = name
	return inst
}

func (inst *innerInitializerWrapper) RunTest() TestingRuntime {
	rt, err := inst.RunEx()
	if err != nil {
		panic(err)
	}
	wb := &innerRuntimeWrapperBuilder{}
	wb.inner = rt
	wb.t = inst.t
	wb.testingDataSrcName = inst.testingDataSrcName
	return wb.Wrap()
}
