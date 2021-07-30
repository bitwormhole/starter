// 这个配置文件是由 starter-configen 工具自动生成的。
// 任何时候，都不要手工修改这里面的内容！！！

package etc

import(
	errors "errors"
	application "github.com/bitwormhole/starter/application"
	config "github.com/bitwormhole/starter/application/config"
	lang "github.com/bitwormhole/starter/lang"
	tester_684224 "github.com/bitwormhole/starter/src/test/tester"
)


func autoGenConfig(configbuilder application.ConfigBuilder) error {

	cominfobuilder := &config.ComInfoBuilder{}
	err := errors.New("OK")

    
	// theContextPropertiesTester
	cominfobuilder.Reset()
	cominfobuilder.ID("theContextPropertiesTester").Class("").Scope("").Aliases("")
	cominfobuilder.OnNew(func() lang.Object {
		return &tester_684224.ContextPropertiesTester{}
	})
	cominfobuilder.OnInit(func(o lang.Object) error {
		return o.(*tester_684224.ContextPropertiesTester).Run()
	})
	cominfobuilder.OnDestroy(func(o lang.Object) error {
		return nil
	})
	cominfobuilder.OnInject(func(o lang.Object, context application.Context) error {
		adapter := &theContextPropertiesTester{}
		adapter.instance = o.(*tester_684224.ContextPropertiesTester)
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

	// theContextResourcesTester
	cominfobuilder.Reset()
	cominfobuilder.ID("theContextResourcesTester").Class("").Scope("").Aliases("")
	cominfobuilder.OnNew(func() lang.Object {
		return &tester_684224.ContextResourcesTester{}
	})
	cominfobuilder.OnInit(func(o lang.Object) error {
		return o.(*tester_684224.ContextResourcesTester).Run()
	})
	cominfobuilder.OnDestroy(func(o lang.Object) error {
		return nil
	})
	cominfobuilder.OnInject(func(o lang.Object, context application.Context) error {
		adapter := &theContextResourcesTester{}
		adapter.instance = o.(*tester_684224.ContextResourcesTester)
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
// type theContextPropertiesTester struct

func (inst *theContextPropertiesTester) __inject__(context application.Context) error {

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
	inst.AppContext=inst.__get_AppContext__(injection, "context")


	// to instance
	instance.AppContext=inst.AppContext


	// invoke custom inject method


	return injection.Close()
}

func (inst * theContextPropertiesTester) __get_AppContext__(injection application.Injection,selector string) application.Context {
	return injection.Context()
}

////////////////////////////////////////////////////////////////////////////////
// type theContextResourcesTester struct

func (inst *theContextResourcesTester) __inject__(context application.Context) error {

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
	inst.AppContext=inst.__get_AppContext__(injection, "context")


	// to instance
	instance.AppContext=inst.AppContext


	// invoke custom inject method


	return injection.Close()
}

func (inst * theContextResourcesTester) __get_AppContext__(injection application.Injection,selector string) application.Context {
	return injection.Context()
}

