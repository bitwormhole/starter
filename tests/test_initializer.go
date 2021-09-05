package tests

import (
	"embed"
	"testing"

	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/collection"
	"github.com/bitwormhole/starter/io/fs"
	"github.com/bitwormhole/starter/lang"
	"github.com/bitwormhole/starter/vlog"
)

// Initializer 是对 application.Initializer 的扩展，添加了几个用于测试的功能
type Initializer interface {
	application.Initializer

	T() *testing.T
	UseResourcesFS(efs *embed.FS, path string) Initializer
	PrepareTestingDir(res string) fs.Path
	PrepareTestingDirZip(zip string) fs.Path
	LoadPropertisFromGitConfig(required bool) Initializer
}

////////////////////////////////////////////////////////////////////////////////

// WrapInitializer 包装  application.Initializer, 为其添加测试的功能
func WrapInitializer(inner application.Initializer, t *testing.T) Initializer {
	return &innerInitializerWrapper{inner: inner, t: t}
}

////////////////////////////////////////////////////////////////////////////////

type innerInitializerWrapper struct {
	inner application.Initializer
	t     *testing.T
}

func (inst *innerInitializerWrapper) _Impl() Initializer {
	return inst
}

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

func (inst *innerInitializerWrapper) Use(module application.Module) application.Initializer {
	return inst.inner.Use(module)
}

func (inst *innerInitializerWrapper) UsePanic() application.Initializer {
	return inst.inner.UsePanic()
}

func (inst *innerInitializerWrapper) LoadPropertisFromGitConfig(required bool) Initializer {
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

func (inst *innerInitializerWrapper) Run() {
	inst.inner.Run()
}

func (inst *innerInitializerWrapper) RunEx() (application.Runtime, error) {
	return inst.inner.RunEx()
}

func (inst *innerInitializerWrapper) T() *testing.T {
	return inst.t
}

func (inst *innerInitializerWrapper) UseResourcesFS(efs *embed.FS, path string) Initializer {
	r := collection.LoadEmbedResources(efs, path)
	inst.UseResources(r)
	return inst
}

func (inst *innerInitializerWrapper) UseResources(res collection.Resources) application.Initializer {
	inst.inner.UseResources(res)
	return inst
}

func (inst *innerInitializerWrapper) UseProperties(p collection.Properties) application.Initializer {
	inst.inner.UseProperties(p)
	return inst
}

func (inst *innerInitializerWrapper) PrepareTestingDir(res string) fs.Path {
	panic("no impl")
	//	return nil
}

func (inst *innerInitializerWrapper) PrepareTestingDirZip(zip string) fs.Path {
	panic("no impl")
	// return nil
}
