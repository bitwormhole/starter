package index

import (
	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/demo/demo4gin"
	"github.com/bitwormhole/starter/demo/index/help"
	"github.com/bitwormhole/starter/tools/configen"
	"github.com/bitwormhole/starter/tools/configen2"
)

func Config(cfg application.ConfigBuilder) {
	configDemoIndex(cfg)
}

func injectRunnerConfigen(t *configen.Runner, context application.RuntimeContext) error {

	return nil
}

func injectRunnerConfigen2(t *configen2.Runner, context application.RuntimeContext) error {

	return nil
}

func injectRunnerDemo1(t *demo4gin.Runner, context application.RuntimeContext) error {

	return nil
}

func injectRunnerDemo2(t *demo4gin.Runner, context application.RuntimeContext) error {

	return nil
}

func injectRunnerHelp(t *help.HelpRunner, context application.RuntimeContext) error {

	return nil
}
