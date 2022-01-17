package application

import (
	"os"

	"github.com/bitwormhole/starter/vlog"
)

// Run 函数启动一个应用实例，返回应用上下文
func Run(config Configuration, args []string) (Context, error) {
	return config.GetLoader().Load(config, args)
}

// // Start 函数用于执行应用的 Loop 前处理
// func Start(context Context) error {
// 	return errors.New("no impl")
// }

// Loop 函数用于执行应用的主循环
func Loop(context Context) error {
	defer func() {
		e := recover()
		if e != nil {
			vlog.Error(e)
		}
	}()
	return runMainLooper(context)
}

// // Stop 函数用于执行应用的 Loop 后处理
// func Stop(context Context) error {
// 	return errors.New("no impl")
// }

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

// RunAndLoop 函数依次调用 Run(), Loop(), Exit()
func RunAndLoop(config Configuration) (int, error) {

	args := os.Args
	context, err := Run(config, args)
	if err != nil {
		return -1, err
	}

	err = Loop(context)
	if err != nil {
		return -1, err
	}

	code, err := Exit(context)
	if err != nil {
		return -1, err
	}

	return code, nil
}
