package lang

import "log"

////////////////////////////////////////////////////////////////////////////////
//  ErrorHandler interface

// ErrorHandler 表示一个错误处理器
type ErrorHandler interface {
	HandleError(err error) error
}

// ErrorHandlerFunc 是ErrorHandler的函数形式
type ErrorHandlerFunc func(err error) error

////////////////////////////////////////////////////////////////////////////////
// defaultErrorHandler struct

type defaultErrorHandler struct {
}

func (inst *defaultErrorHandler) HandleError(err error) error {
	log.Output(0, err.Error())
	return err
}

////////////////////////////////////////////////////////////////////////////////
// adapterErrorHandler struct

type adapterErrorHandler struct {
	fn ErrorHandlerFunc
}

func (inst *adapterErrorHandler) HandleError(err error) error {
	fn := inst.fn
	if fn != nil {
		return fn(err)
	}
	return err
}

// NewErrorHandlerForFunc 创建一个新的 ErrorHandler 作为fn的代理
func NewErrorHandlerForFunc(fn ErrorHandlerFunc) ErrorHandler {
	return &adapterErrorHandler{fn: fn}
}

////////////////////////////////////////////////////////////////////////////////
// func

// DefaultErrorHandler 创建一个默认的错误处理器
func DefaultErrorHandler() ErrorHandler {
	return &defaultErrorHandler{}
}
