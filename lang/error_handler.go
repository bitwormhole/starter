package lang

import "log"

////////////////////////////////////////////////////////////////////////////////
//  ErrorHandler interface

type ErrorHandler interface {
	OnError(err error)
}

////////////////////////////////////////////////////////////////////////////////
// defaultErrorHandler struct

type defaultErrorHandler struct {
}

func (inst *defaultErrorHandler) OnError(err error) {
	log.Output(0, err.Error())
}

////////////////////////////////////////////////////////////////////////////////
// func

func DefaultErrorHandler() ErrorHandler {
	return &defaultErrorHandler{}
}
