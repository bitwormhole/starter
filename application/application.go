package application

type ExitCodeGenerator interface {
	GetExitCode() int
}

// Run 函数启动一个应用实例，返回应用上下文
func Run(config Configuration, args []string) (RuntimeContext, error) {
	return config.GetLoader().Load(config, args)
}

// Exit 函数用于退出应用
func Exit(context RuntimeContext) (int, error) {

	exitcodegen := tryGetExitCodeGenerator(context)
	// errHandler := context.GetErrorHandler()

	err := context.GetReleasePool().Release()
	if err != nil {
		return 0, err
	}

	if exitcodegen == nil {
		return 0, nil
	}

	code := exitcodegen.GetExitCode()
	return code, nil
}

func tryGetExitCodeGenerator(context RuntimeContext) ExitCodeGenerator {
	obj, err := context.GetComponents().GetComponentByClass("ExitCodeGenerator")
	if err != nil {
		return nil
	}
	gen, ok := obj.(ExitCodeGenerator)
	if ok {
		return gen
	}
	return nil
}
