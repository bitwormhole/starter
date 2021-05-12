package index
// This file is auto-generate by configen, never edit it.

import (
	application "github.com/bitwormhole/starter/application"
	config "github.com/bitwormhole/starter/application/config"
	demo4gin "github.com/bitwormhole/starter/demo/demo4gin"
	configen "github.com/bitwormhole/starter/tools/configen"
	configen2 "github.com/bitwormhole/starter/tools/configen2"
	help "github.com/bitwormhole/starter/demo/index/help"
	lang "github.com/bitwormhole/starter/lang"
)

func configDemoIndex(cfg application.ConfigBuilder){

	// runner-for-demo4gorm
	cfg.AddComponent(&config.ComInfo{
		ID:"runner-for-demo4gorm",
		OnInject: func(obj lang.Object, context application.RuntimeContext) error {
			t := obj.(*demo4gin.Runner)
			return injectRunnerDemo1(t,context)
		},
		OnNew: func() lang.Object {
			return &demo4gin.Runner{}
		},
	})

	// runner-for-demo4gin
	cfg.AddComponent(&config.ComInfo{
		ID:"runner-for-demo4gin",
		OnInject: func(obj lang.Object, context application.RuntimeContext) error {
			t := obj.(*demo4gin.Runner)
			return injectRunnerDemo2(t,context)
		},
		OnNew: func() lang.Object {
			return &demo4gin.Runner{}
		},
	})

	// runner-for-configen
	cfg.AddComponent(&config.ComInfo{
		ID:"runner-for-configen",
		OnInject: func(obj lang.Object, context application.RuntimeContext) error {
			t := obj.(*configen.Runner)
			return injectRunnerConfigen(t,context)
		},
		OnNew: func() lang.Object {
			return &configen.Runner{}
		},
	})

	// runner-for-configen2
	cfg.AddComponent(&config.ComInfo{
		ID:"runner-for-configen2",
		OnInject: func(obj lang.Object, context application.RuntimeContext) error {
			t := obj.(*configen2.Runner)
			return injectRunnerConfigen2(t,context)
		},
		OnNew: func() lang.Object {
			return &configen2.Runner{}
		},
	})

	// runner-for-help
	cfg.AddComponent(&config.ComInfo{
		ID:"runner-for-help",
		OnInject: func(obj lang.Object, context application.RuntimeContext) error {
			t := obj.(*help.HelpRunner)
			return injectRunnerHelp(t,context)
		},
		OnNew: func() lang.Object {
			return &help.HelpRunner{}
		},
	})
}

