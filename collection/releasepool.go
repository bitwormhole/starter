package collection

import (
	"github.com/bitwormhole/starter/lang"
)

type ReleasePool interface {
	Release() error
	Push(target lang.Disposable)
}

func Release(pool ReleasePool) error {
	if pool == nil {
		return nil
	}
	return pool.Release()
}
