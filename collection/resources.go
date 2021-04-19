package collection

import "io"

// Resources 接口提供一组获取资源的方法
type Resources interface {
	GetText(path string) (string, error)
	GetBinary(path string) ([]byte, error)
	GetReader(path string) (io.ReadCloser, error)
}
