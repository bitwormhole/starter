package lang

import "io"

// Disposable 接口用于释放对象持有的资源
type Disposable interface {
	Dispose() error
}

type ReleasePool interface {
	Release() error
	Push(target Disposable)
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

func Release(pool ReleasePool) error {
	if pool == nil {
		return nil
	}
	return pool.Release()
}
