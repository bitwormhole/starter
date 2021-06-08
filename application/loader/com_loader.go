package loader

import (
	"errors"
	"sort"

	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/application/runtime"
	"github.com/bitwormhole/starter/collection"
	"github.com/bitwormhole/starter/lang"
)

////////////////////////////////////////////////////////////////////////////////
// StandardComponentLoader

type StandardComponentLoader struct {
}

func (inst *StandardComponentLoader) init() application.ComponentLoader {
	return inst
}

func (inst *StandardComponentLoader) OpenLoading(context application.Context) (application.ComponentLoading, error) {
	loading := &standardComponentsLoading{}
	_, err := loading.open(context)
	if err != nil {
		return nil, err
	}
	return loading, nil
}

////////////////////////////////////////////////////////////////////////////////
// standardComponentsLoading

type standardComponentsLoading struct {
	pool          lang.ReleasePool
	runtime       application.Context
	contextProxy  *runtime.ContextProxy
	comTable      map[string]*comHolderProxy
	comIndexCount int
	closed        bool
	lastError     error
}

func (inst *standardComponentsLoading) open(contextRT application.Context) (application.ComponentLoading, error) {

	comTable := make(map[string]*comHolderProxy)
	pool := &lang.SimpleReleasePool{}
	proxy := &runtime.ContextProxy{}
	loadingCtx := &standardComponentsLoadingContext{}

	proxy.Runtime = contextRT
	proxy.Creation = loadingCtx
	proxy.Current = loadingCtx

	inst.comTable = comTable
	inst.pool = pool
	inst.runtime = contextRT
	inst.contextProxy = proxy

	loadingCtx.init(inst)
	return inst, nil
}

func (inst *standardComponentsLoading) OnError(err error) {
	if err == nil {
		return
	}
	inst.lastError = err
}

func (inst *standardComponentsLoading) Pool() lang.ReleasePool {
	return inst.pool
}

func (inst *standardComponentsLoading) Context() application.Context {
	return inst.contextProxy
}

func (inst *standardComponentsLoading) Load(holder application.ComponentHolder) (lang.Object, error) {
	chp, err := inst.getComHolderProxy(holder, true)
	if err != nil {
		return nil, err
	}
	err = chp.doInject(inst.contextProxy)
	if err != nil {
		return nil, err
	}
	com := chp.instance.Get()
	return com, nil
}

func (inst *standardComponentsLoading) getComHolderProxy(holder application.ComponentHolder, create bool) (*comHolderProxy, error) {
	if holder == nil {
		return nil, errors.New("ComponentHolder==nil")
	}
	id := holder.GetInfo().GetID()
	proxy := inst.comTable[id]
	if proxy == nil {
		if create {
			proxy = &comHolderProxy{}
			proxy.initial(holder)
			proxy.index = inst.comIndexCount
			inst.comTable[id] = proxy
			inst.comIndexCount++
		} else {
			return nil, errors.New("no comHolderProxy with id: " + id)
		}
	}
	return proxy, nil
}

func (inst *standardComponentsLoading) LoadAll(src []application.ComponentHolder) ([]lang.Object, error) {
	if src == nil {
		return nil, errors.New("holderList==nil")
	}
	dst := make([]lang.Object, 0)
	for index := range src {
		holder := src[index]
		obj, err := inst.Load(holder)
		if err != nil {
			return nil, err
		}
		dst = append(dst, obj)
	}
	return dst, nil
}

func (inst *standardComponentsLoading) Close() error {

	lastErr := inst.lastError
	if lastErr != nil {
		return lastErr
	}

	if inst.closed {
		return nil
	} else {
		inst.closed = true
	}

	for {
		cnt, err := inst.tryInjectToAllComponents()
		if err != nil {
			return err
		}
		if cnt < 1 {
			break
		}
	}

	list := make([]*comHolderProxy, 0)
	table := inst.comTable

	for key := range table {
		item := table[key]
		if item == nil {
			continue
		}
		list = append(list, item)
	}

	sorter := &comHolderProxyListSorter{}
	sorter.list = list
	sorter.sort()

	err := inst.doInitForAllComponents(list)
	if err != nil {
		return err
	}

	inst.contextProxy.Close()
	return nil
}

func (inst *standardComponentsLoading) tryInjectToAllComponents() (int, error) {
	count := 0
	ctx := inst.contextProxy
	table := inst.comTable
	for key := range table {
		item := table[key]
		if item.injected {
			continue
		}
		err := item.doInject(ctx)
		if err != nil {
			return 0, err
		}
		count++
	}
	return count, nil
}

func (inst *standardComponentsLoading) doInitForAllComponents(list []*comHolderProxy) error {
	ctx := inst.contextProxy
	for index := range list {
		item := list[index]
		err := item.doInit(ctx)
		if err != nil {
			return err
		}
	}
	return nil
}

////////////////////////////////////////////////////////////////////////////////
// EOF

type standardComponentsLoadingContext struct {
	loading       *standardComponentsLoading
	loadingFacade application.ComponentLoading
}

func (inst *standardComponentsLoadingContext) init(loading *standardComponentsLoading) application.Context {
	inst.loading = loading
	inst.loadingFacade = &standardComponentsLoadingFacade{inner: loading}
	return inst
}

func (inst *standardComponentsLoadingContext) GetArguments() collection.Arguments {
	return inst.loading.runtime.GetArguments()
}

func (inst *standardComponentsLoadingContext) GetAttributes() collection.Attributes {
	return inst.loading.runtime.GetAttributes()
}

func (inst *standardComponentsLoadingContext) GetParameters() collection.Parameters {
	return inst.loading.runtime.GetParameters()
}

func (inst *standardComponentsLoadingContext) GetProperties() collection.Properties {
	return inst.loading.runtime.GetProperties()
}

func (inst *standardComponentsLoadingContext) GetEnvironment() collection.Environment {
	return inst.loading.runtime.GetEnvironment()
}

func (inst *standardComponentsLoadingContext) GetResources() collection.Resources {
	return inst.loading.runtime.GetResources()
}

func (inst *standardComponentsLoadingContext) GetComponents() application.Components {
	return inst.loading.runtime.GetComponents()
}

func (inst *standardComponentsLoadingContext) GetReleasePool() lang.ReleasePool {
	return inst.loading.pool
}

func (inst *standardComponentsLoadingContext) GetErrorHandler() lang.ErrorHandler {
	return inst.loading.runtime.GetErrorHandler()
}

func (inst *standardComponentsLoadingContext) SetErrorHandler(eh lang.ErrorHandler) {
	inst.loading.runtime.SetErrorHandler(eh)
}

func (inst *standardComponentsLoadingContext) NewChild() application.Context {
	return inst
}

func (inst *standardComponentsLoadingContext) Injector() application.Injector {
	return inst.loading.runtime.Injector()
}

func (inst *standardComponentsLoadingContext) ComponentLoader() application.ComponentLoader {
	return inst
}

func (inst *standardComponentsLoadingContext) OpenLoading(ctx application.Context) (application.ComponentLoading, error) {
	return inst.loadingFacade, nil
}

func (inst *standardComponentsLoadingContext) GetComponent(selector string) (lang.Object, error) {
	getter := &runtime.ComGetter{}
	getter.Init(inst)
	return getter.GetOne(selector)
}

func (inst *standardComponentsLoadingContext) GetComponentList(selector string) ([]lang.Object, error) {
	getter := &runtime.ComGetter{}
	getter.Init(inst)
	return getter.GetList(selector)
}

func (inst *standardComponentsLoadingContext) GetApplicationName() string {
	return inst.loading.runtime.GetApplicationName()
}

func (inst *standardComponentsLoadingContext) GetApplicationVersion() string {
	return inst.loading.runtime.GetApplicationVersion()
}

func (inst *standardComponentsLoadingContext) GetStartupTimestamp() int64 {
	return inst.loading.runtime.GetStartupTimestamp()
}

func (inst *standardComponentsLoadingContext) GetShutdownTimestamp() int64 {
	return inst.loading.runtime.GetShutdownTimestamp()
}

func (inst *standardComponentsLoadingContext) GetURI() string {
	return inst.loading.runtime.GetURI()
}

////////////////////////////////////////////////////////////////////////////////
// standardComponentsLoadingFacade

type standardComponentsLoadingFacade struct {
	inner application.ComponentLoading
}

func (inst *standardComponentsLoadingFacade) init() application.ComponentLoading {
	return inst
}

func (inst *standardComponentsLoadingFacade) OnError(err error) {
	inst.inner.OnError(err)
}

func (inst *standardComponentsLoadingFacade) Pool() lang.ReleasePool {
	return inst.inner.Pool()
}

func (inst *standardComponentsLoadingFacade) Context() application.Context {
	return inst.inner.Context()
}

func (inst *standardComponentsLoadingFacade) Load(h application.ComponentHolder) (lang.Object, error) {
	return inst.inner.Load(h)
}

func (inst *standardComponentsLoadingFacade) LoadAll(hhh []application.ComponentHolder) ([]lang.Object, error) {
	return inst.inner.LoadAll(hhh)
}

func (inst *standardComponentsLoadingFacade) Close() error {
	// do NOP
	return nil
}

////////////////////////////////////////////////////////////////////////////////
// comHolderProxyListSorter

type comHolderProxyListSorter struct {
	list []*comHolderProxy
}

func (inst *comHolderProxyListSorter) sort() {
	sort.Sort(inst)
}

func (inst *comHolderProxyListSorter) Len() int {
	return len(inst.list)
}

func (inst *comHolderProxyListSorter) Less(i1 int, i2 int) bool {
	list := inst.list
	item1 := list[i1]
	item2 := list[i2]
	return item1.index < item2.index
}

func (inst *comHolderProxyListSorter) Swap(i1 int, i2 int) {
	list := inst.list
	item1 := list[i1]
	item2 := list[i2]
	list[i1] = item2
	list[i2] = item1
}

////////////////////////////////////////////////////////////////////////////////
// comHolderProxy

type comHolderProxy struct {
	holder   application.ComponentHolder
	instance application.ComponentInstance

	index    int // for sort
	injected bool
	started  bool
}

func (inst *comHolderProxy) initial(holder application.ComponentHolder) error {
	if holder == nil {
		return errors.New("com.holder==nil")
	}
	inst.holder = holder
	inst.instance = holder.GetInstance()
	return nil
}

func (inst *comHolderProxy) doInject(ctx application.Context) error {

	if inst.injected {
		return nil
	} else {
		inst.injected = true
	}

	if inst.instance.IsLoaded() {
		return nil
	}

	return inst.instance.Inject(ctx)
}

func (inst *comHolderProxy) doInit(ctx application.Context) error {

	if inst.started {
		return nil
	} else {
		inst.started = true
	}

	if inst.instance.IsLoaded() {
		return nil
	}

	err := inst.instance.Init()
	if err != nil {
		return err
	}

	ctx.GetReleasePool().Push(inst)
	return nil
}

func (inst *comHolderProxy) Dispose() error {
	instance := inst.instance
	return instance.Destroy()
}

////////////////////////////////////////////////////////////////////////////////
// EOF
