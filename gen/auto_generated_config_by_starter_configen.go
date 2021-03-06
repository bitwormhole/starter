// (todo:gen2.template) 
// 这个配置文件是由 starter-configen 工具自动生成的。
// 任何时候，都不要手工修改这里面的内容！！！

package gen

import (
	application "github.com/bitwormhole/starter/application"
	config "github.com/bitwormhole/starter/application/config"
	bootstrap0x1b594d "github.com/bitwormhole/starter/bootstrap"
	lang "github.com/bitwormhole/starter/lang"
	util "github.com/bitwormhole/starter/util"
	configenchecker0xe7a472 "github.com/bitwormhole/starter/util/configenchecker"
	vlog0x6d1dd2 "github.com/bitwormhole/starter/vlog"
	std0x639069 "github.com/bitwormhole/starter/vlog/std"
    
)


func nop(x ... interface{}){
	util.Int64ToTime(0)
	lang.CreateReleasePool()
}


func autoGenConfig(cb application.ConfigBuilder) error {

	var err error = nil
	cominfobuilder := config.ComInfo()
	nop(err,cominfobuilder)

	// component: main-looper
	cominfobuilder.Next()
	cominfobuilder.ID("main-looper").Class("").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComBoot{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com1-configenchecker.ConfigenChecker
	cominfobuilder.Next()
	cominfobuilder.ID("com1-configenchecker.ConfigenChecker").Class("").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4theConfigenChecker{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: vlog-std-context
	cominfobuilder.Next()
	cominfobuilder.ID("vlog-std-context").Class("").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4theVlogDefaultContext{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: vlog-std-logger-factory
	cominfobuilder.Next()
	cominfobuilder.ID("vlog-std-logger-factory").Class("").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4theVlogLoggerFactory{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: vlog-default-formatter
	cominfobuilder.Next()
	cominfobuilder.ID("vlog-default-formatter").Class("").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4theVlogDefaultFormatter{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: vlog-std-main-channel
	cominfobuilder.Next()
	cominfobuilder.ID("vlog-std-main-channel").Class("vlog-std-channel").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4theVlogMainChannel{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: vlog-std-distributor
	cominfobuilder.Next()
	cominfobuilder.ID("vlog-std-distributor").Class("").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4theVlogDistributor{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: vlog-std-console-channel
	cominfobuilder.Next()
	cominfobuilder.ID("vlog-std-console-channel").Class("vlog-std-channel vlog-std-sub-channel").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4theVlogConsoleChannel{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: vlog-std-console-writer
	cominfobuilder.Next()
	cominfobuilder.ID("vlog-std-console-writer").Class("").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4theVlogConsoleWriter{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: vlog-std-file-channel
	cominfobuilder.Next()
	cominfobuilder.ID("vlog-std-file-channel").Class("vlog-std-channel vlog-std-sub-channel").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4theVlogFileChannel{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: vlog-std-file-writer
	cominfobuilder.Next()
	cominfobuilder.ID("vlog-std-file-writer").Class("").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4theVlogFileWriter{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}



    return nil
}

////////////////////////////////////////////////////////////////////////////////

// comFactory4pComBoot : the factory of component: main-looper
type comFactory4pComBoot struct {

    mPrototype * bootstrap0x1b594d.Boot

	
	mLivesSelector config.InjectionSelector
	mConcurrentSelector config.InjectionSelector

}

func (inst * comFactory4pComBoot) init() application.ComponentFactory {

	
	inst.mLivesSelector = config.NewInjectionSelector(".life",nil)
	inst.mConcurrentSelector = config.NewInjectionSelector("${application.loopers.concurrent}",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComBoot) newObject() * bootstrap0x1b594d.Boot {
	return & bootstrap0x1b594d.Boot {}
}

func (inst * comFactory4pComBoot) castObject(instance application.ComponentInstance) * bootstrap0x1b594d.Boot {
	return instance.Get().(*bootstrap0x1b594d.Boot)
}

func (inst * comFactory4pComBoot) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComBoot) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComBoot) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComBoot) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComBoot) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComBoot) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.Lives = inst.getterForFieldLivesSelector(context)
	obj.Concurrent = inst.getterForFieldConcurrentSelector(context)
	return context.LastError()
}

//getterForFieldLivesSelector
func (inst * comFactory4pComBoot) getterForFieldLivesSelector (context application.InstanceContext) []application.LifeRegistry {
	list1 := inst.mLivesSelector.GetList(context)
	list2 := make([]application.LifeRegistry, 0, len(list1))
	for _, item1 := range list1 {
		item2, ok := item1.(application.LifeRegistry)
		if ok {
			list2 = append(list2, item2)
		}
	}
	return list2
}

//getterForFieldConcurrentSelector
func (inst * comFactory4pComBoot) getterForFieldConcurrentSelector (context application.InstanceContext) bool {
    return inst.mConcurrentSelector.GetBool(context)
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4theConfigenChecker : the factory of component: com1-configenchecker.ConfigenChecker
type comFactory4theConfigenChecker struct {

    mPrototype * configenchecker0xe7a472.ConfigenChecker

	
	mContextSelector config.InjectionSelector
	mEnableSelector config.InjectionSelector

}

func (inst * comFactory4theConfigenChecker) init() application.ComponentFactory {

	
	inst.mContextSelector = config.NewInjectionSelector("context",nil)
	inst.mEnableSelector = config.NewInjectionSelector("${configen.checker.enable}",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4theConfigenChecker) newObject() * configenchecker0xe7a472.ConfigenChecker {
	return & configenchecker0xe7a472.ConfigenChecker {}
}

func (inst * comFactory4theConfigenChecker) castObject(instance application.ComponentInstance) * configenchecker0xe7a472.ConfigenChecker {
	return instance.Get().(*configenchecker0xe7a472.ConfigenChecker)
}

func (inst * comFactory4theConfigenChecker) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4theConfigenChecker) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4theConfigenChecker) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4theConfigenChecker) Init(instance application.ComponentInstance) error {
	return inst.castObject(instance).Check()
}

func (inst * comFactory4theConfigenChecker) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4theConfigenChecker) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.Context = inst.getterForFieldContextSelector(context)
	obj.Enable = inst.getterForFieldEnableSelector(context)
	return context.LastError()
}

//getterForFieldContextSelector
func (inst * comFactory4theConfigenChecker) getterForFieldContextSelector (context application.InstanceContext) application.Context {
    return context.Context()
}

//getterForFieldEnableSelector
func (inst * comFactory4theConfigenChecker) getterForFieldEnableSelector (context application.InstanceContext) bool {
    return inst.mEnableSelector.GetBool(context)
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4theVlogDefaultContext : the factory of component: vlog-std-context
type comFactory4theVlogDefaultContext struct {

    mPrototype * std0x639069.DefaultContext

	
	mDefaultLevelSelector config.InjectionSelector
	mDefaultFormatterSelector config.InjectionSelector
	mChannelsSelector config.InjectionSelector
	mMainChannelSelector config.InjectionSelector

}

func (inst * comFactory4theVlogDefaultContext) init() application.ComponentFactory {

	
	inst.mDefaultLevelSelector = config.NewInjectionSelector("${vlog.default.level}",nil)
	inst.mDefaultFormatterSelector = config.NewInjectionSelector("#vlog-default-formatter",nil)
	inst.mChannelsSelector = config.NewInjectionSelector(".vlog-std-channel",nil)
	inst.mMainChannelSelector = config.NewInjectionSelector("#vlog-std-main-channel",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4theVlogDefaultContext) newObject() * std0x639069.DefaultContext {
	return & std0x639069.DefaultContext {}
}

func (inst * comFactory4theVlogDefaultContext) castObject(instance application.ComponentInstance) * std0x639069.DefaultContext {
	return instance.Get().(*std0x639069.DefaultContext)
}

func (inst * comFactory4theVlogDefaultContext) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4theVlogDefaultContext) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4theVlogDefaultContext) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4theVlogDefaultContext) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4theVlogDefaultContext) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4theVlogDefaultContext) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.DefaultLevel = inst.getterForFieldDefaultLevelSelector(context)
	obj.DefaultFormatter = inst.getterForFieldDefaultFormatterSelector(context)
	obj.Channels = inst.getterForFieldChannelsSelector(context)
	obj.MainChannel = inst.getterForFieldMainChannelSelector(context)
	return context.LastError()
}

//getterForFieldDefaultLevelSelector
func (inst * comFactory4theVlogDefaultContext) getterForFieldDefaultLevelSelector (context application.InstanceContext) string {
    return inst.mDefaultLevelSelector.GetString(context)
}

//getterForFieldDefaultFormatterSelector
func (inst * comFactory4theVlogDefaultContext) getterForFieldDefaultFormatterSelector (context application.InstanceContext) vlog0x6d1dd2.Formatter {

	o1 := inst.mDefaultFormatterSelector.GetOne(context)
	o2, ok := o1.(vlog0x6d1dd2.Formatter)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "vlog-std-context")
		eb.Set("field", "DefaultFormatter")
		eb.Set("type1", "?")
		eb.Set("type2", "vlog0x6d1dd2.Formatter")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldChannelsSelector
func (inst * comFactory4theVlogDefaultContext) getterForFieldChannelsSelector (context application.InstanceContext) []vlog0x6d1dd2.Channel {
	list1 := inst.mChannelsSelector.GetList(context)
	list2 := make([]vlog0x6d1dd2.Channel, 0, len(list1))
	for _, item1 := range list1 {
		item2, ok := item1.(vlog0x6d1dd2.Channel)
		if ok {
			list2 = append(list2, item2)
		}
	}
	return list2
}

//getterForFieldMainChannelSelector
func (inst * comFactory4theVlogDefaultContext) getterForFieldMainChannelSelector (context application.InstanceContext) vlog0x6d1dd2.Channel {

	o1 := inst.mMainChannelSelector.GetOne(context)
	o2, ok := o1.(vlog0x6d1dd2.Channel)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "vlog-std-context")
		eb.Set("field", "MainChannel")
		eb.Set("type1", "?")
		eb.Set("type2", "vlog0x6d1dd2.Channel")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4theVlogLoggerFactory : the factory of component: vlog-std-logger-factory
type comFactory4theVlogLoggerFactory struct {

    mPrototype * std0x639069.StandardLoggerFactory

	
	mContextSelector config.InjectionSelector

}

func (inst * comFactory4theVlogLoggerFactory) init() application.ComponentFactory {

	
	inst.mContextSelector = config.NewInjectionSelector("#vlog-std-context",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4theVlogLoggerFactory) newObject() * std0x639069.StandardLoggerFactory {
	return & std0x639069.StandardLoggerFactory {}
}

func (inst * comFactory4theVlogLoggerFactory) castObject(instance application.ComponentInstance) * std0x639069.StandardLoggerFactory {
	return instance.Get().(*std0x639069.StandardLoggerFactory)
}

func (inst * comFactory4theVlogLoggerFactory) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4theVlogLoggerFactory) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4theVlogLoggerFactory) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4theVlogLoggerFactory) Init(instance application.ComponentInstance) error {
	return inst.castObject(instance).Start()
}

func (inst * comFactory4theVlogLoggerFactory) Destroy(instance application.ComponentInstance) error {
	return inst.castObject(instance).Stop()
}

func (inst * comFactory4theVlogLoggerFactory) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.Context = inst.getterForFieldContextSelector(context)
	return context.LastError()
}

//getterForFieldContextSelector
func (inst * comFactory4theVlogLoggerFactory) getterForFieldContextSelector (context application.InstanceContext) std0x639069.Context {

	o1 := inst.mContextSelector.GetOne(context)
	o2, ok := o1.(std0x639069.Context)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "vlog-std-logger-factory")
		eb.Set("field", "Context")
		eb.Set("type1", "?")
		eb.Set("type2", "std0x639069.Context")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4theVlogDefaultFormatter : the factory of component: vlog-default-formatter
type comFactory4theVlogDefaultFormatter struct {

    mPrototype * std0x639069.DefaultFormatter

	

}

func (inst * comFactory4theVlogDefaultFormatter) init() application.ComponentFactory {

	


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4theVlogDefaultFormatter) newObject() * std0x639069.DefaultFormatter {
	return & std0x639069.DefaultFormatter {}
}

func (inst * comFactory4theVlogDefaultFormatter) castObject(instance application.ComponentInstance) * std0x639069.DefaultFormatter {
	return instance.Get().(*std0x639069.DefaultFormatter)
}

func (inst * comFactory4theVlogDefaultFormatter) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4theVlogDefaultFormatter) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4theVlogDefaultFormatter) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4theVlogDefaultFormatter) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4theVlogDefaultFormatter) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4theVlogDefaultFormatter) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	return nil
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4theVlogMainChannel : the factory of component: vlog-std-main-channel
type comFactory4theVlogMainChannel struct {

    mPrototype * std0x639069.LogChannel

	
	mContextSelector config.InjectionSelector
	mNameSelector config.InjectionSelector
	mEnableSelector config.InjectionSelector
	mWriterSelector config.InjectionSelector
	mLevelSelector config.InjectionSelector

}

func (inst * comFactory4theVlogMainChannel) init() application.ComponentFactory {

	
	inst.mContextSelector = config.NewInjectionSelector("#vlog-std-context",nil)
	inst.mNameSelector = config.NewInjectionSelector("vlog-main",nil)
	inst.mEnableSelector = config.NewInjectionSelector("${vlog.main.enable}",nil)
	inst.mWriterSelector = config.NewInjectionSelector("#vlog-std-distributor",nil)
	inst.mLevelSelector = config.NewInjectionSelector("${vlog.main.level}",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4theVlogMainChannel) newObject() * std0x639069.LogChannel {
	return & std0x639069.LogChannel {}
}

func (inst * comFactory4theVlogMainChannel) castObject(instance application.ComponentInstance) * std0x639069.LogChannel {
	return instance.Get().(*std0x639069.LogChannel)
}

func (inst * comFactory4theVlogMainChannel) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4theVlogMainChannel) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4theVlogMainChannel) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4theVlogMainChannel) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4theVlogMainChannel) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4theVlogMainChannel) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.Context = inst.getterForFieldContextSelector(context)
	obj.Name = inst.getterForFieldNameSelector(context)
	obj.Enable = inst.getterForFieldEnableSelector(context)
	obj.Writer = inst.getterForFieldWriterSelector(context)
	obj.Level = inst.getterForFieldLevelSelector(context)
	return context.LastError()
}

//getterForFieldContextSelector
func (inst * comFactory4theVlogMainChannel) getterForFieldContextSelector (context application.InstanceContext) std0x639069.Context {

	o1 := inst.mContextSelector.GetOne(context)
	o2, ok := o1.(std0x639069.Context)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "vlog-std-main-channel")
		eb.Set("field", "Context")
		eb.Set("type1", "?")
		eb.Set("type2", "std0x639069.Context")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldNameSelector
func (inst * comFactory4theVlogMainChannel) getterForFieldNameSelector (context application.InstanceContext) string {
    return inst.mNameSelector.GetString(context)
}

//getterForFieldEnableSelector
func (inst * comFactory4theVlogMainChannel) getterForFieldEnableSelector (context application.InstanceContext) bool {
    return inst.mEnableSelector.GetBool(context)
}

//getterForFieldWriterSelector
func (inst * comFactory4theVlogMainChannel) getterForFieldWriterSelector (context application.InstanceContext) vlog0x6d1dd2.Writer {

	o1 := inst.mWriterSelector.GetOne(context)
	o2, ok := o1.(vlog0x6d1dd2.Writer)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "vlog-std-main-channel")
		eb.Set("field", "Writer")
		eb.Set("type1", "?")
		eb.Set("type2", "vlog0x6d1dd2.Writer")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldLevelSelector
func (inst * comFactory4theVlogMainChannel) getterForFieldLevelSelector (context application.InstanceContext) string {
    return inst.mLevelSelector.GetString(context)
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4theVlogDistributor : the factory of component: vlog-std-distributor
type comFactory4theVlogDistributor struct {

    mPrototype * std0x639069.Distributor

	
	mChannelsSelector config.InjectionSelector

}

func (inst * comFactory4theVlogDistributor) init() application.ComponentFactory {

	
	inst.mChannelsSelector = config.NewInjectionSelector(".vlog-std-sub-channel",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4theVlogDistributor) newObject() * std0x639069.Distributor {
	return & std0x639069.Distributor {}
}

func (inst * comFactory4theVlogDistributor) castObject(instance application.ComponentInstance) * std0x639069.Distributor {
	return instance.Get().(*std0x639069.Distributor)
}

func (inst * comFactory4theVlogDistributor) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4theVlogDistributor) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4theVlogDistributor) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4theVlogDistributor) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4theVlogDistributor) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4theVlogDistributor) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.Channels = inst.getterForFieldChannelsSelector(context)
	return context.LastError()
}

//getterForFieldChannelsSelector
func (inst * comFactory4theVlogDistributor) getterForFieldChannelsSelector (context application.InstanceContext) []vlog0x6d1dd2.Channel {
	list1 := inst.mChannelsSelector.GetList(context)
	list2 := make([]vlog0x6d1dd2.Channel, 0, len(list1))
	for _, item1 := range list1 {
		item2, ok := item1.(vlog0x6d1dd2.Channel)
		if ok {
			list2 = append(list2, item2)
		}
	}
	return list2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4theVlogConsoleChannel : the factory of component: vlog-std-console-channel
type comFactory4theVlogConsoleChannel struct {

    mPrototype * std0x639069.LogChannel

	
	mContextSelector config.InjectionSelector
	mNameSelector config.InjectionSelector
	mEnableSelector config.InjectionSelector
	mWriterSelector config.InjectionSelector
	mLevelSelector config.InjectionSelector

}

func (inst * comFactory4theVlogConsoleChannel) init() application.ComponentFactory {

	
	inst.mContextSelector = config.NewInjectionSelector("#vlog-std-context",nil)
	inst.mNameSelector = config.NewInjectionSelector("vlog-console",nil)
	inst.mEnableSelector = config.NewInjectionSelector("${vlog.console.enable}",nil)
	inst.mWriterSelector = config.NewInjectionSelector("#vlog-std-console-writer",nil)
	inst.mLevelSelector = config.NewInjectionSelector("${vlog.console.level}",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4theVlogConsoleChannel) newObject() * std0x639069.LogChannel {
	return & std0x639069.LogChannel {}
}

func (inst * comFactory4theVlogConsoleChannel) castObject(instance application.ComponentInstance) * std0x639069.LogChannel {
	return instance.Get().(*std0x639069.LogChannel)
}

func (inst * comFactory4theVlogConsoleChannel) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4theVlogConsoleChannel) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4theVlogConsoleChannel) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4theVlogConsoleChannel) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4theVlogConsoleChannel) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4theVlogConsoleChannel) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.Context = inst.getterForFieldContextSelector(context)
	obj.Name = inst.getterForFieldNameSelector(context)
	obj.Enable = inst.getterForFieldEnableSelector(context)
	obj.Writer = inst.getterForFieldWriterSelector(context)
	obj.Level = inst.getterForFieldLevelSelector(context)
	return context.LastError()
}

//getterForFieldContextSelector
func (inst * comFactory4theVlogConsoleChannel) getterForFieldContextSelector (context application.InstanceContext) std0x639069.Context {

	o1 := inst.mContextSelector.GetOne(context)
	o2, ok := o1.(std0x639069.Context)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "vlog-std-console-channel")
		eb.Set("field", "Context")
		eb.Set("type1", "?")
		eb.Set("type2", "std0x639069.Context")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldNameSelector
func (inst * comFactory4theVlogConsoleChannel) getterForFieldNameSelector (context application.InstanceContext) string {
    return inst.mNameSelector.GetString(context)
}

//getterForFieldEnableSelector
func (inst * comFactory4theVlogConsoleChannel) getterForFieldEnableSelector (context application.InstanceContext) bool {
    return inst.mEnableSelector.GetBool(context)
}

//getterForFieldWriterSelector
func (inst * comFactory4theVlogConsoleChannel) getterForFieldWriterSelector (context application.InstanceContext) vlog0x6d1dd2.Writer {

	o1 := inst.mWriterSelector.GetOne(context)
	o2, ok := o1.(vlog0x6d1dd2.Writer)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "vlog-std-console-channel")
		eb.Set("field", "Writer")
		eb.Set("type1", "?")
		eb.Set("type2", "vlog0x6d1dd2.Writer")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldLevelSelector
func (inst * comFactory4theVlogConsoleChannel) getterForFieldLevelSelector (context application.InstanceContext) string {
    return inst.mLevelSelector.GetString(context)
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4theVlogConsoleWriter : the factory of component: vlog-std-console-writer
type comFactory4theVlogConsoleWriter struct {

    mPrototype * std0x639069.ConsoleWriter

	

}

func (inst * comFactory4theVlogConsoleWriter) init() application.ComponentFactory {

	


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4theVlogConsoleWriter) newObject() * std0x639069.ConsoleWriter {
	return & std0x639069.ConsoleWriter {}
}

func (inst * comFactory4theVlogConsoleWriter) castObject(instance application.ComponentInstance) * std0x639069.ConsoleWriter {
	return instance.Get().(*std0x639069.ConsoleWriter)
}

func (inst * comFactory4theVlogConsoleWriter) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4theVlogConsoleWriter) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4theVlogConsoleWriter) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4theVlogConsoleWriter) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4theVlogConsoleWriter) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4theVlogConsoleWriter) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	return nil
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4theVlogFileChannel : the factory of component: vlog-std-file-channel
type comFactory4theVlogFileChannel struct {

    mPrototype * std0x639069.LogChannel

	
	mContextSelector config.InjectionSelector
	mNameSelector config.InjectionSelector
	mEnableSelector config.InjectionSelector
	mWriterSelector config.InjectionSelector
	mLevelSelector config.InjectionSelector

}

func (inst * comFactory4theVlogFileChannel) init() application.ComponentFactory {

	
	inst.mContextSelector = config.NewInjectionSelector("#vlog-std-context",nil)
	inst.mNameSelector = config.NewInjectionSelector("vlog-file",nil)
	inst.mEnableSelector = config.NewInjectionSelector("${vlog.file.enable}",nil)
	inst.mWriterSelector = config.NewInjectionSelector("#vlog-std-file-writer",nil)
	inst.mLevelSelector = config.NewInjectionSelector("${vlog.file.level}",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4theVlogFileChannel) newObject() * std0x639069.LogChannel {
	return & std0x639069.LogChannel {}
}

func (inst * comFactory4theVlogFileChannel) castObject(instance application.ComponentInstance) * std0x639069.LogChannel {
	return instance.Get().(*std0x639069.LogChannel)
}

func (inst * comFactory4theVlogFileChannel) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4theVlogFileChannel) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4theVlogFileChannel) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4theVlogFileChannel) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4theVlogFileChannel) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4theVlogFileChannel) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.Context = inst.getterForFieldContextSelector(context)
	obj.Name = inst.getterForFieldNameSelector(context)
	obj.Enable = inst.getterForFieldEnableSelector(context)
	obj.Writer = inst.getterForFieldWriterSelector(context)
	obj.Level = inst.getterForFieldLevelSelector(context)
	return context.LastError()
}

//getterForFieldContextSelector
func (inst * comFactory4theVlogFileChannel) getterForFieldContextSelector (context application.InstanceContext) std0x639069.Context {

	o1 := inst.mContextSelector.GetOne(context)
	o2, ok := o1.(std0x639069.Context)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "vlog-std-file-channel")
		eb.Set("field", "Context")
		eb.Set("type1", "?")
		eb.Set("type2", "std0x639069.Context")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldNameSelector
func (inst * comFactory4theVlogFileChannel) getterForFieldNameSelector (context application.InstanceContext) string {
    return inst.mNameSelector.GetString(context)
}

//getterForFieldEnableSelector
func (inst * comFactory4theVlogFileChannel) getterForFieldEnableSelector (context application.InstanceContext) bool {
    return inst.mEnableSelector.GetBool(context)
}

//getterForFieldWriterSelector
func (inst * comFactory4theVlogFileChannel) getterForFieldWriterSelector (context application.InstanceContext) vlog0x6d1dd2.Writer {

	o1 := inst.mWriterSelector.GetOne(context)
	o2, ok := o1.(vlog0x6d1dd2.Writer)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "vlog-std-file-channel")
		eb.Set("field", "Writer")
		eb.Set("type1", "?")
		eb.Set("type2", "vlog0x6d1dd2.Writer")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldLevelSelector
func (inst * comFactory4theVlogFileChannel) getterForFieldLevelSelector (context application.InstanceContext) string {
    return inst.mLevelSelector.GetString(context)
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4theVlogFileWriter : the factory of component: vlog-std-file-writer
type comFactory4theVlogFileWriter struct {

    mPrototype * std0x639069.FileWriter

	
	mEnableSelector config.InjectionSelector
	mPath1Selector config.InjectionSelector
	mPath2Selector config.InjectionSelector
	mContextSelector config.InjectionSelector

}

func (inst * comFactory4theVlogFileWriter) init() application.ComponentFactory {

	
	inst.mEnableSelector = config.NewInjectionSelector("${vlog.file.enable}",nil)
	inst.mPath1Selector = config.NewInjectionSelector("${vlog.file.path1}",nil)
	inst.mPath2Selector = config.NewInjectionSelector("${vlog.file.path2}",nil)
	inst.mContextSelector = config.NewInjectionSelector("context",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4theVlogFileWriter) newObject() * std0x639069.FileWriter {
	return & std0x639069.FileWriter {}
}

func (inst * comFactory4theVlogFileWriter) castObject(instance application.ComponentInstance) * std0x639069.FileWriter {
	return instance.Get().(*std0x639069.FileWriter)
}

func (inst * comFactory4theVlogFileWriter) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4theVlogFileWriter) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4theVlogFileWriter) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4theVlogFileWriter) Init(instance application.ComponentInstance) error {
	return inst.castObject(instance).Open()
}

func (inst * comFactory4theVlogFileWriter) Destroy(instance application.ComponentInstance) error {
	return inst.castObject(instance).Close()
}

func (inst * comFactory4theVlogFileWriter) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.Enable = inst.getterForFieldEnableSelector(context)
	obj.Path1 = inst.getterForFieldPath1Selector(context)
	obj.Path2 = inst.getterForFieldPath2Selector(context)
	obj.Context = inst.getterForFieldContextSelector(context)
	return context.LastError()
}

//getterForFieldEnableSelector
func (inst * comFactory4theVlogFileWriter) getterForFieldEnableSelector (context application.InstanceContext) bool {
    return inst.mEnableSelector.GetBool(context)
}

//getterForFieldPath1Selector
func (inst * comFactory4theVlogFileWriter) getterForFieldPath1Selector (context application.InstanceContext) string {
    return inst.mPath1Selector.GetString(context)
}

//getterForFieldPath2Selector
func (inst * comFactory4theVlogFileWriter) getterForFieldPath2Selector (context application.InstanceContext) string {
    return inst.mPath2Selector.GetString(context)
}

//getterForFieldContextSelector
func (inst * comFactory4theVlogFileWriter) getterForFieldContextSelector (context application.InstanceContext) application.Context {
    return context.Context()
}




