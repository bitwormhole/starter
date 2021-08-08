package collection

import "io"

// Resource 结构包含某个资源的信息
type Resource struct {
	Name         string
	AbsolutePath string
	RelativePath string
	BasePath     string
	IsDir        bool
}

// Res 接口代表一个资源实体
type Res interface {
	// 加载文本
	ReadText() (string, error)

	// 加载二进制数据
	ReadBinary() ([]byte, error)

	// 读数据流
	Reader() (io.ReadCloser, error)

	// 取长度
	Length() int64

	IsFile() bool
	IsDir() bool
	Exists() bool
}

// Resources 接口提供一组获取资源的方法
type Resources interface {

	// 加载文本数据
	GetText(path string) (string, error)

	// 加载二进制数据
	GetBinary(path string) ([]byte, error)

	// 读数据流
	GetReader(path string) (io.ReadCloser, error)

	// 列出所有资源的路径, 相当于{{ List("/",true) }}
	All() []*Resource

	// 列出所有资源的路径
	List(path string, recursive bool) []*Resource

	Get(path string) (Res, error)

	Clear()

	Import(src map[string]Res, override bool)

	Export(dst map[string]Res) map[string]Res
}
