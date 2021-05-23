// 这个文件是由 starter-configen 工具生成的配置代码，千万不要手工修改里面的任何内容。
package demo

import(
	application "github.com/bitwormhole/starter/application"
	config "github.com/bitwormhole/starter/application/config"
	elements_a3e8536a "github.com/bitwormhole/starter/demo/elements"
	lang "github.com/bitwormhole/starter/lang"
)

func Config(cb application.ConfigBuilder) error {

    // exampleComponent1
    cb.AddComponent(&config.ComInfo{
		ID: "com1",
		Class: "com1",
		Scope: application.ScopeSingleton,
		Aliases: []string{},
		OnNew: func() lang.Object {
		    return &elements_a3e8536a.ComExample1{}
		},
		OnInject: func(obj lang.Object,context application.RuntimeContext) error {
		    target := obj.(*elements_a3e8536a.ComExample1)
		    return exampleComponent1(target,context)
		},
    })

    // exampleComponent2
    cb.AddComponent(&config.ComInfo{
		ID: "com2",
		Class: "com2  Looper",
		Scope: application.ScopeSingleton,
		Aliases: []string{},
		OnNew: func() lang.Object {
		    return &elements_a3e8536a.ComExample2{}
		},
		OnInit: func(obj lang.Object) error {
		    target := obj.(*elements_a3e8536a.ComExample2)
		    return target.Open()
		},
		OnDestroy: func(obj lang.Object) error {
		    target := obj.(*elements_a3e8536a.ComExample2)
		    return target.Close()
		},
		OnInject: func(obj lang.Object,context application.RuntimeContext) error {
		    target := obj.(*elements_a3e8536a.ComExample2)
		    return exampleComponent2(target,context)
		},
    })

    return nil
}

