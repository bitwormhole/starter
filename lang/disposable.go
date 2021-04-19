package lang

import "io"

// Disposable 接口用于释放对象持有的资源
type Disposable interface {
	Dispose() error
}

func Dispose(t Disposable) error {
	if t == nil {
		return nil
	}
	return t.Dispose()
}

func Close(t io.Closer) error {
	if t == nil {
		return nil
	}
	return t.Close()
}
