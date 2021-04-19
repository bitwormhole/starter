package runtime

import (
	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/collection"
	"github.com/bitwormhole/starter/lang"
)

////////////////////////////////////////////////////////////////////////////////
// struct

type runtimeContextCore struct {

	// facade

	context    application.RuntimeContext
	components application.Components

	// collection

	releasePool collection.ReleasePool

	arguments   collection.Arguments
	attributes  collection.Attributes
	environment collection.Environment
	properties  collection.Properties
	parameters  collection.Parameters
	resources   collection.Resources

	// helper
	finder       *componentFinder
	loader       *runtimeComponentLoader
	errorHandler lang.ErrorHandler

	// data

	componentTable map[string]application.ComponentHolder

	time1      int64
	time2      int64
	appName    string
	appVersion string
	uri        string
}

type creationContextCore struct {

	// parent
	parent *runtimeContextCore

	// facade
	context    application.RuntimeContext
	components application.Components
	facade     application.CreationContext
	proxy      *contextProxy

	// helper
	finder *componentFinder
	loader *creationComponentLoader

	// data
	pool  collection.ReleasePool
	scope application.ComponentScope
	cache map[string]*componentLoading
}

////////////////////////////////////////////////////////////////////////////////
// create

func createRuntimeContextCore() *runtimeContextCore {

	inst := &runtimeContextCore{}
	comTable := make(map[string]application.ComponentHolder)

	inst.arguments = collection.CreateArguments()
	inst.attributes = collection.CreateAttributes()
	inst.environment = collection.CreateEnvironment()
	inst.parameters = collection.CreateParameters()
	inst.properties = collection.CreateProperties()
	inst.releasePool = collection.CreateReleasePool()
	inst.resources = nil

	inst.context = &runtimeContextFacade{core: inst}
	inst.components = &runtimeComponentsFacade{core: inst}
	inst.componentTable = comTable

	inst.finder = &componentFinder{table: comTable}
	inst.loader = &runtimeComponentLoader{core: inst}

	inst.time1 = 0 // now

	return inst
}

func createCreationContextCore(parent *runtimeContextCore) *creationContextCore {

	inst := &creationContextCore{}

	facade := &creationContextFacade{core: inst}
	proxy := &contextProxy{}
	components := &creationComponentsFacade{core: inst}
	loader := &creationComponentLoader{core: inst}

	inst.parent = parent
	inst.context = facade
	inst.components = components
	inst.facade = facade
	inst.proxy = proxy
	inst.loader = loader
	inst.finder = parent.finder
	inst.cache = make(map[string]*componentLoading)

	proxy.runtime = parent.context
	proxy.creation = inst.context
	proxy.current = inst.context

	return inst
}

func CreateRuntimeContext(parent application.RuntimeContext) (application.RuntimeContext, error) {

	// create new
	core := createRuntimeContextCore()
	child := core.context

	if parent == nil {
		return child, nil
	}

	// export & import
	core.resources = parent.GetResources()
	child.GetArguments().Import(parent.GetArguments().Export())
	child.GetAttributes().Import(parent.GetAttributes().Export(nil))
	child.GetEnvironment().Import(parent.GetEnvironment().Export(nil))
	child.GetParameters().Import(parent.GetParameters().Export(nil))
	child.GetProperties().Import(parent.GetProperties().Export(nil))
	child.GetComponents().Import(parent.GetComponents().Export(nil))

	return child, nil
}

////////////////////////////////////////////////////////////////////////////////
// EOF
