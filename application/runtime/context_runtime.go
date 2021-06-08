package runtime

import (
	"errors"

	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/collection"
	"github.com/bitwormhole/starter/lang"
)

////////////////////////////////////////////////////////////////////////////////
// struct  RuntimeContextBuilder

type RuntimeContextBuilder struct {
	Parent application.Context

	Time1      int64
	Time2      int64
	AppName    string
	AppVersion string
	URL        string

	Injector  application.Injector
	ComLoader application.ComponentLoader
	Resources collection.Resources
	Pool      lang.ReleasePool
}

func (inst *RuntimeContextBuilder) initContextPre(ctx *contextRuntime) error {

	// create new
	components := &componentTable{}
	components.context = ctx
	components.table = make(map[string]application.ComponentHolder)
	ctx.components = components

	ctx.arguments = collection.CreateArguments()
	ctx.attributes = collection.CreateAttributes()
	ctx.environment = collection.CreateEnvironment()
	ctx.parameters = collection.CreateParameters()
	ctx.properties = collection.CreateProperties()

	ctx.comLoader = inst.ComLoader
	ctx.injector = inst.Injector
	ctx.resources = inst.Resources
	ctx.pool = inst.Pool

	return nil
}

func (inst *RuntimeContextBuilder) initContextWithParent(parent application.Context, child *contextRuntime) error {

	// export & import
	child.GetArguments().Import(parent.GetArguments().Export())
	child.GetAttributes().Import(parent.GetAttributes().Export(nil))
	child.GetEnvironment().Import(parent.GetEnvironment().Export(nil))
	child.GetParameters().Import(parent.GetParameters().Export(nil))
	child.GetProperties().Import(parent.GetProperties().Export(nil))
	child.GetComponents().Import(parent.GetComponents().Export(nil))

	if child.resources == nil {
		child.resources = parent.GetResources()
	}

	if child.comLoader == nil {
		child.comLoader = parent.ComponentLoader()
	}

	if child.injector == nil {
		child.injector = parent.Injector()
	}

	return nil
}

func (inst *RuntimeContextBuilder) initContextWithoutParent(ctx *contextRuntime) error {
	return nil
}

func (inst *RuntimeContextBuilder) initContextPost(ctx *contextRuntime) error {

	ctx.time1 = 0 // now

	if ctx.pool == nil {
		ctx.pool = lang.CreateReleasePool()
	}

	if ctx.resources == nil {
		return errors.New("no Resources")
	}

	if ctx.injector == nil {
		ctx.injector = &innerInjector{}
	}

	if ctx.comLoader == nil {
		return errors.New("no ComponentLoader")
	}

	if ctx.errorHandler == nil {
		ctx.errorHandler = lang.DefaultErrorHandler()
	}

	return nil
}

func (inst *RuntimeContextBuilder) Create() (application.Context, error) {

	context := &contextRuntime{}
	err := inst.initContextPre(context)
	if err != nil {
		return nil, err
	}

	parent := inst.Parent
	if parent == nil {
		err = inst.initContextWithoutParent(context)
	} else {
		err = inst.initContextWithParent(parent, context)
	}
	if err != nil {
		return nil, err
	}

	err = inst.initContextPost(context)
	if err != nil {
		return nil, err
	}

	return context, nil
}

////////////////////////////////////////////////////////////////////////////////
// struct contextRuntime

type contextRuntime struct {

	// core
	pool         lang.ReleasePool
	components   application.Components
	errorHandler lang.ErrorHandler
	comLoader    application.ComponentLoader
	injector     application.Injector

	// collection
	arguments   collection.Arguments
	attributes  collection.Attributes
	environment collection.Environment
	properties  collection.Properties
	parameters  collection.Parameters
	resources   collection.Resources

	// info
	time1      int64
	time2      int64
	appName    string
	appVersion string
	uri        string
}

func (inst *contextRuntime) GetComponents() application.Components {
	return inst.components
}

func (inst *contextRuntime) GetReleasePool() lang.ReleasePool {
	return inst.pool
}

func (inst *contextRuntime) GetArguments() collection.Arguments {
	args := inst.arguments
	return args
}

func (inst *contextRuntime) GetAttributes() collection.Attributes {
	return inst.attributes
}

func (inst *contextRuntime) GetEnvironment() collection.Environment {
	return inst.environment
}

func (inst *contextRuntime) GetProperties() collection.Properties {
	return inst.properties
}

func (inst *contextRuntime) GetParameters() collection.Parameters {
	return inst.parameters
}

func (inst *contextRuntime) GetResources() collection.Resources {
	return inst.resources
}

func (inst *contextRuntime) GetApplicationName() string {
	return inst.appName
}

func (inst *contextRuntime) GetApplicationVersion() string {
	return inst.appVersion
}

func (inst *contextRuntime) GetStartupTimestamp() int64 {
	return inst.time1
}

func (inst *contextRuntime) GetShutdownTimestamp() int64 {
	return inst.time2
}

func (inst *contextRuntime) GetURI() string {
	return inst.uri
}

func (inst *contextRuntime) GetErrorHandler() lang.ErrorHandler {
	h := inst.errorHandler
	if h == nil {
		h = lang.DefaultErrorHandler()
	}
	return h
}

func (inst *contextRuntime) SetErrorHandler(h lang.ErrorHandler) {
	inst.errorHandler = h
}

func (inst *contextRuntime) GetComponent(selector string) (lang.Object, error) {
	getter := &ComGetter{context: inst}
	return getter.GetOne(selector)
}

func (inst *contextRuntime) GetComponentList(selector string) ([]lang.Object, error) {
	getter := &ComGetter{context: inst}
	return getter.GetList(selector)
}

func (inst *contextRuntime) Injector() application.Injector {
	return inst.injector
}

func (inst *contextRuntime) ComponentLoader() application.ComponentLoader {
	return inst.comLoader
}

func (inst *contextRuntime) NewChild() application.Context {
	builder := &RuntimeContextBuilder{}
	builder.Parent = inst
	child, err := builder.Create()
	if err != nil {
		panic(err)
	}
	return child
}

////////////////////////////////////////////////////////////////////////////////
