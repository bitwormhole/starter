package application

import (
	"errors"
)

const ExitCodeGeneratorClassName = "app-exit-code-generator"

type ExitCodeGenerator interface {
	GetExitCode() int
}

const LooperClassName = "app-looper"

type Looper interface {
	Loop() error
}

// Run 函数启动一个应用实例，返回应用上下文
func Run(config Configuration, args []string) (RuntimeContext, error) {
	return config.GetLoader().Load(config, args)
}

// Loop 函数用于执行应用的主循环
func Loop(context RuntimeContext) error {
	looper, err := tryGetLooper(context)
	if looper == nil || err != nil {
		return nil
	}
	return looper.Loop()
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
	selector := "." + ExitCodeGeneratorClassName
	obj, err := context.FindComponent(selector)
	if err != nil {
		return nil
	}
	gen, ok := obj.(ExitCodeGenerator)
	if ok {
		return gen
	}
	return nil
}

func tryGetLooper(context RuntimeContext) (Looper, error) {
	selector := "." + LooperClassName
	com, err := context.FindComponent(selector)
	if err != nil {
		return nil, err
	}
	looper, ok := com.(Looper)
	if !ok {
		return nil, errors.New("object is not a Looper interface.")
	}
	return looper, nil
}
