package application

// Run 函数启动一个应用实例，返回应用上下文
func Run(config Configuration, args []string) (Context, error) {
	return config.GetLoader().Load(config, args)
}

// Loop 函数用于执行应用的主循环
func Loop(context Context) error {
	looper := tryGetLooper(context)
	if looper == nil {
		return nil
	}
	return looper.Loop()
}

// Exit 函数用于退出应用
func Exit(context Context) (int, error) {

	exitcodegen := tryGetExitCodeGenerator(context)
	// errHandler := context.GetErrorHandler()

	err := context.GetReleasePool().Release()
	if err != nil {
		return -1, err
	}

	if exitcodegen == nil {
		return 0, nil
	}

	code := exitcodegen.ExitCode()
	return code, nil
}
