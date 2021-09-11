// (todo:gen2.template)
// 这个配置文件是由 starter-configen 工具自动生成的。
// 任何时候，都不要手工修改这里面的内容！！！

package etc

import (
	application "github.com/bitwormhole/starter/application"
	config "github.com/bitwormhole/starter/application/config"
	lang "github.com/bitwormhole/starter/lang"
	tester0x684224 "github.com/bitwormhole/starter/src/test/tester"
)

func autoGenConfig(cb application.ConfigBuilder) error {

	var err error = nil
	cominfobuilder := config.ComInfo()

	// component: com0-tester.ContextPropertiesTester
	cominfobuilder.Next()
	cominfobuilder.ID("com0-tester.ContextPropertiesTester").Class("").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4theContextPropertiesTester{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com1-tester.ContextResourcesTester
	cominfobuilder.Next()
	cominfobuilder.ID("com1-tester.ContextResourcesTester").Class("").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4theContextResourcesTester{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	return nil
}

////////////////////////////////////////////////////////////////////////////////

// comFactory4theContextPropertiesTester : the factory of component: com0-tester.ContextPropertiesTester
type comFactory4theContextPropertiesTester struct {
	mPrototype *tester0x684224.ContextPropertiesTester

	mEnableSelector     config.InjectionSelector
	mAppContextSelector config.InjectionSelector
}

func (inst *comFactory4theContextPropertiesTester) init() application.ComponentFactory {

	inst.mEnableSelector = config.NewInjectionSelector("${test.enable}", nil)
	inst.mAppContextSelector = config.NewInjectionSelector("context", nil)

	inst.mPrototype = inst.newObject()
	return inst
}

func (inst *comFactory4theContextPropertiesTester) newObject() *tester0x684224.ContextPropertiesTester {
	return &tester0x684224.ContextPropertiesTester{}
}

func (inst *comFactory4theContextPropertiesTester) castObject(instance application.ComponentInstance) *tester0x684224.ContextPropertiesTester {
	return instance.Get().(*tester0x684224.ContextPropertiesTester)
}

func (inst *comFactory4theContextPropertiesTester) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst *comFactory4theContextPropertiesTester) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst *comFactory4theContextPropertiesTester) AfterService() application.ComponentAfterService {
	return inst
}

func (inst *comFactory4theContextPropertiesTester) Init(instance application.ComponentInstance) error {
	return inst.castObject(instance).Run()
}

func (inst *comFactory4theContextPropertiesTester) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst *comFactory4theContextPropertiesTester) Inject(instance application.ComponentInstance, context application.InstanceContext) error {

	obj := inst.castObject(instance)
	obj.Enable = inst.getterForFieldEnableSelector(context)
	obj.AppContext = inst.getterForFieldAppContextSelector(context)
	return context.LastError()
}

//getterForFieldEnableSelector
func (inst *comFactory4theContextPropertiesTester) getterForFieldEnableSelector(context application.InstanceContext) bool {
	return inst.mEnableSelector.GetBool(context)
}

//getterForFieldAppContextSelector
func (inst *comFactory4theContextPropertiesTester) getterForFieldAppContextSelector(context application.InstanceContext) application.Context {
	return context.Context()
}

////////////////////////////////////////////////////////////////////////////////

// comFactory4theContextResourcesTester : the factory of component: com1-tester.ContextResourcesTester
type comFactory4theContextResourcesTester struct {
	mPrototype *tester0x684224.ContextResourcesTester

	mEnableSelector     config.InjectionSelector
	mAppContextSelector config.InjectionSelector
}

func (inst *comFactory4theContextResourcesTester) init() application.ComponentFactory {

	inst.mEnableSelector = config.NewInjectionSelector("${test.enable}", nil)
	inst.mAppContextSelector = config.NewInjectionSelector("context", nil)

	inst.mPrototype = inst.newObject()
	return inst
}

func (inst *comFactory4theContextResourcesTester) newObject() *tester0x684224.ContextResourcesTester {
	return &tester0x684224.ContextResourcesTester{}
}

func (inst *comFactory4theContextResourcesTester) castObject(instance application.ComponentInstance) *tester0x684224.ContextResourcesTester {
	return instance.Get().(*tester0x684224.ContextResourcesTester)
}

func (inst *comFactory4theContextResourcesTester) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst *comFactory4theContextResourcesTester) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst *comFactory4theContextResourcesTester) AfterService() application.ComponentAfterService {
	return inst
}

func (inst *comFactory4theContextResourcesTester) Init(instance application.ComponentInstance) error {
	return inst.castObject(instance).Run()
}

func (inst *comFactory4theContextResourcesTester) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst *comFactory4theContextResourcesTester) Inject(instance application.ComponentInstance, context application.InstanceContext) error {

	obj := inst.castObject(instance)
	obj.Enable = inst.getterForFieldEnableSelector(context)
	obj.AppContext = inst.getterForFieldAppContextSelector(context)
	return context.LastError()
}

//getterForFieldEnableSelector
func (inst *comFactory4theContextResourcesTester) getterForFieldEnableSelector(context application.InstanceContext) bool {
	return inst.mEnableSelector.GetBool(context)
}

//getterForFieldAppContextSelector
func (inst *comFactory4theContextResourcesTester) getterForFieldAppContextSelector(context application.InstanceContext) application.Context {
	return context.Context()
}
