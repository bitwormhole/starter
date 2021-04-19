package application

type ExitCodeGenerator interface {
	GetExitCode() int
}

// Run 函数启动一个应用实例，返回应用上下文
func Run(config Configuration, args []string) (RuntimeContext, error) {
	return config.GetLoader().Load(config, args)
}

// Exit 函数用于退出应用
func Exit(context RuntimeContext) int {

	exitcodegen := tryGetExitCodeGenerator(context)
	errList := context.GetReleasePool().Release()

	if errList != nil {
		errHandler := context.GetErrorHandler()
		for index := range errList {
			err := errList[index]
			errHandler.OnError(err)
		}
	}

	if exitcodegen != nil {
		return exitcodegen.GetExitCode()
	}

	return 0
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
