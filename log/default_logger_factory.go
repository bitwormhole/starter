package log

import "fmt"

type DefaultLoggerFactory struct {
}

func (inst *DefaultLoggerFactory) _Impl() LoggerFactory {
	return inst
}

func (inst *DefaultLoggerFactory) CreateLogger() Logger {

	fmt.Println("")

	return nil
}
