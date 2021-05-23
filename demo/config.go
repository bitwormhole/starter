package demo

import (
	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/demo/elements"
	lang "github.com/bitwormhole/starter/lang"
)

////////////////////////////////////////////////////////////////////////////////

func exampleComponent1(com *elements.ComExample1, context application.RuntimeContext) error {

	// [component]
	//	id=com1
	//	class=com1

	in := context.Injector()
	com.Name = in.GetPropertyString("com1.name", "cc1")
	return in.Done()
}

func exampleComponent2(com *elements.ComExample2, context application.RuntimeContext) error {

	// [component]
	//	id=com2
	//	class=com2  Looper
	//  initMethod=Open
	//  destroyMethod=Close

	in := context.Injector()

	in.Inject("#com1").Accept(func(key string, h application.ComponentHolder) bool {
		pt := h.GetPrototype()
		_, ok := pt.(*elements.ComExample1)
		return ok
	}).To(func(o lang.Object) bool {
		t, ok := o.(*elements.ComExample1)
		if ok {
			com.Com1ref = t
		}
		return ok
	})

	return in.Done()
}

func loop1(com *elements.Looper1, context application.RuntimeContext) error {

	// [component]
	//	class= looper

	return nil
}

func loop2(com *elements.Looper2, context application.RuntimeContext) error {

	// [component]
	//	class= looper

	return nil
}

func exit1(com *elements.ExitCoder1, context application.RuntimeContext) error {

	// [component]
	//	class= exit-code-generator

	return nil
}

func exit2(com *elements.ExitCoder2, context application.RuntimeContext) error {

	// [component]
	//	class= exit-code-generator

	return nil
}
