package tests

import (
	"testing"

	"github.com/bitwormhole/starter"
	"github.com/bitwormhole/starter/lang"
)

// TestingStarter 创建一个测试用的启动器
func TestingStarter(t *testing.T) Initializer {
	i := starter.InitApp()
	i.SetErrorHandler(makeTestingErrorHandler(t))
	return WrapInitializer(i, t)
}

func makeTestingErrorHandler(t *testing.T) lang.ErrorHandler {
	return lang.NewErrorHandlerForFunc(func(err error) error {
		t.Error(err)
		return err
	})
}
