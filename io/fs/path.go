package fs

import (
	"io"

	"github.com/bitwormhole/starter/lang"
)

// FileSystem  代表一个抽象的文件系统
type FileSystem interface {
	Roots() []Path
	Resolve(path string) (Path, error)
	GetPath(path string) Path
	GetPathByURI(uri lang.URI) (Path, error)

	Separator() string
	SeparatorChar() rune
	PathSeparator() string
	PathSeparatorChar() rune

	DefaultReadOptions() *Options
	DefaultWriteOptions() *Options
	SetDefaultOptions(r *Options, w *Options)
}

// Path 代表一个路径
type Path interface {

	// base

	URI() lang.URI
	String() string
	Parent() Path
	FileSystem() FileSystem

	Path() string
	Name() string
	Exists() bool
	IsDir() bool
	IsFile() bool
	IsSymlink() bool
	LastModTime() int64 // ms from unix time(1970-01-01 00:00:00)
	GetMeta() FileMeta
	SetMeta(meta FileMeta)

	Delete() error

	// for file

	Size() int64

	GetIO() FileIO

	CreateFile(opt *Options) error
	CreateFileWithSize(size int64, opt *Options) error

	CopyTo(target Path) error
	MoveTo(target Path) error

	// for dir

	Mkdir() error
	Mkdirs() error

	ListNames() []string // 返回短文件名
	ListPaths() []string // 返回完整路径名
	ListItems() []Path

	GetChild(name string) Path
	GetHref(name string) Path
}

// // IoMode 执行文件IO操作的模式
// type IoMode interface {
// 	Flag() int
// 	Perm() os.FileMode
// }

// FileIO 表示对一个具体文件的IO
type FileIO interface {
	Path() Path
	WriteText(text string, opt *Options, mkdirs bool) error
	WriteBinary(data []byte, opt *Options, mkdirs bool) error
	ReadText(opt *Options) (string, error)
	ReadBinary(opt *Options) ([]byte, error)
	OpenReader(opt *Options) (io.ReadCloser, error)
	OpenWriter(opt *Options, mkdirs bool) (io.WriteCloser, error)
}

// FileMeta  表示对一个具体文件的 posix liked mode
type FileMeta interface {
	LastModTime() int64
	Size() int64

	IsFile() bool
	IsDir() bool
	IsSymlink() bool
	Exists() bool
	CanExecute() bool
	CanRead() bool
	CanWrite() bool
}
