// (todo:gen2.template)
// 这个配置文件是由 starter-configen 工具自动生成的。
// 任何时候，都不要手工修改这里面的内容！！！

package gen

import (
	application "github.com/bitwormhole/starter/application"
	config "github.com/bitwormhole/starter/application/config"
	lang "github.com/bitwormhole/starter/lang"
	config0x6718c1 "github.com/bitwormhole/starter/src/test/cli/mod/config"
)

func autoGenConfig(cb application.ConfigBuilder) error {

	var err error = nil
	cominfobuilder := config.ComInfo()

	// component: com0-config0x6718c1.DemoHandler
	cominfobuilder.Next()
	cominfobuilder.ID("com0-config0x6718c1.DemoHandler").Class("cli-handler").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComDemoHandler{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	return nil
}

////////////////////////////////////////////////////////////////////////////////

// comFactory4pComDemoHandler : the factory of component: com0-config0x6718c1.DemoHandler
type comFactory4pComDemoHandler struct {
	mPrototype *config0x6718c1.DemoHandler
}

func (inst *comFactory4pComDemoHandler) init() application.ComponentFactory {

	inst.mPrototype = inst.newObject()
	return inst
}

func (inst *comFactory4pComDemoHandler) newObject() *config0x6718c1.DemoHandler {
	return &config0x6718c1.DemoHandler{}
}

func (inst *comFactory4pComDemoHandler) castObject(instance application.ComponentInstance) *config0x6718c1.DemoHandler {
	return instance.Get().(*config0x6718c1.DemoHandler)
}

func (inst *comFactory4pComDemoHandler) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst *comFactory4pComDemoHandler) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst *comFactory4pComDemoHandler) AfterService() application.ComponentAfterService {
	return inst
}

func (inst *comFactory4pComDemoHandler) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst *comFactory4pComDemoHandler) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst *comFactory4pComDemoHandler) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	return nil
}
