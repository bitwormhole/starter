// 这个配置文件是由 starter-configen 工具自动生成的。
// 任何时候，都不要手工修改这里面的内容！！！

package etcstarter

import(
	errors "errors"
	application "github.com/bitwormhole/starter/application"
	config "github.com/bitwormhole/starter/application/config"
	lang "github.com/bitwormhole/starter/lang"
	configenchecker_e7a472 "github.com/bitwormhole/starter/util/configenchecker"
)


func autoGenConfig(configbuilder application.ConfigBuilder) error {

	cominfobuilder := &config.ComInfoBuilder{}
	err := errors.New("OK")

    
	// theConfigenChecker
	cominfobuilder.Reset()
	cominfobuilder.ID("theConfigenChecker").Class("").Scope("").Aliases("")
	cominfobuilder.OnNew(func() lang.Object {
		return &configenchecker_e7a472.ConfigenChecker{}
	})
	cominfobuilder.OnInit(func(o lang.Object) error {
		return o.(*configenchecker_e7a472.ConfigenChecker).Check()
	})
	cominfobuilder.OnDestroy(func(o lang.Object) error {
		return nil
	})
	cominfobuilder.OnInject(func(o lang.Object, context application.Context) error {
		adapter := &theConfigenChecker{}
		adapter.instance = o.(*configenchecker_e7a472.ConfigenChecker)
		// adapter.context = context
		err := adapter.__inject__(context)
		if err != nil {
			return err
		}
		return nil
	})
	err = cominfobuilder.CreateTo(configbuilder)
    if err !=nil{
        return err
    }


	return nil
}


////////////////////////////////////////////////////////////////////////////////
// type theConfigenChecker struct

func (inst *theConfigenChecker) __inject__(context application.Context) error {

	// prepare
	instance := inst.instance
	injection, err := context.Injector().OpenInjection(context)
	if err != nil {
		return err
	}
	defer injection.Close()
	if instance == nil {
		return nil
	}

	// from getters
	inst.Context=inst.__get_Context__(injection, "context")
	inst.Enable=inst.__get_Enable__(injection, "${configen.checker.enable}")


	// to instance
	instance.Context=inst.Context
	instance.Enable=inst.Enable


	// invoke custom inject method


	return injection.Close()
}

func (inst * theConfigenChecker) __get_Context__(injection application.Injection,selector string) application.Context {
	return injection.Context()
}

func (inst * theConfigenChecker) __get_Enable__(injection application.Injection,selector string) bool {
	reader := injection.Select(selector)
	defer reader.Close()
	value, err := reader.ReadBool()
	if err != nil {
		injection.OnError(err)
	}
	return value
}

