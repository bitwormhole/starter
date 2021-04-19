package fs

import (
	"os"
)

// FileSystem  代表一个抽象的文件系统
type FileSystem interface {
	Roots() []Path
	GetPath(path string) Path
	Separator() string
	SeparatorChar() rune
	PathSeparator() string
	PathSeparatorChar() rune
}

// Path 代表一个路径
type Path interface {

	// base

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

	CreateFile(mode IoMode) error
	CreateFileWithSize(size int64, mode IoMode) error

	CopyTo(target Path) error
	MoveTo(target Path) error

	// for dir

	Mkdir() error
	Mkdirs() error

	GetNameList() []string // 返回短文件名
	GetPathList() []string // 返回完整路径名
	GetItemList() []Path

	GetChild(name string) Path
	GetHref(name string) Path
}

// IoMode 执行文件IO操作的模式
type IoMode interface {
	Flag() int
	Perm() os.FileMode
}

// FileIO 表示对一个具体文件的IO
type FileIO interface {
	Path() Path
	WriteText(text string, mode IoMode) error
	WriteBinary(data []byte, mode IoMode) error
	ReadText() (string, error)
	ReadBinary() ([]byte, error)
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
