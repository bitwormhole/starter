package tests

import (
	"testing"

	"github.com/bitwormhole/starter"
	"github.com/bitwormhole/starter/lang"
)

// Starter 创建一个测试用的启动器
func Starter(t *testing.T) TestingInitializer {
	i := starter.InitApp()
	i.SetErrorHandler(makeTestingErrorHandler(t))
	i.SetExitEnabled(false)
	i.SetPanicEnabled(true)
	wrapper := wrapInitializer(i, t)
	wrapper.LoadPropertisFromGitConfig(false)
	return wrapper
}

func makeTestingErrorHandler(t *testing.T) lang.ErrorHandler {
	return lang.NewErrorHandlerForFunc(func(err error) error {
		t.Error(err)
		return err
	})
}
