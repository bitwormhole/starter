package fs

import (
	"errors"
	"io"
	"net/url"
	"os"
	"path/filepath"
	"runtime"

	"github.com/bitwormhole/starter/lang"
	"github.com/bitwormhole/starter/util"
)

type apiPlatform interface {
	Roots() []string
	PathSeparatorChar() rune
	SeparatorChar() rune
}

type innerFSCore struct {
	fs       FileSystem
	platform apiPlatform

	separator         string
	separatorChar     rune
	pathSeparator     string
	pathSeparatorChar rune
}

type innerPath struct {
	// impl Path
	core *innerFSCore
	path string
}

type innerFileSystem struct {
	// impl FileSystem
	core *innerFSCore
}

var innerFileSystemDefaultOptionsR *Options
var innerFileSystemDefaultOptionsW *Options

// impl innerFileSystem

// Default 创建一个默认的 FileSystem 实例
func Default() FileSystem {

	sys := runtime.GOOS
	var platform apiPlatform

	if sys == "windows" {
		platform = &innerWindowsPlatform{}
	} else {
		platform = &innerPosixPlatform{}
	}

	// create
	core := &innerFSCore{}
	fs := &innerFileSystem{}

	// binding
	core.fs = fs
	core.platform = platform
	core.pathSeparatorChar = platform.PathSeparatorChar()
	core.pathSeparator = string(platform.PathSeparatorChar())
	core.separatorChar = platform.SeparatorChar()
	core.separator = string(platform.SeparatorChar())

	fs.core = core

	return fs
}

// impl innerFileSystem

func (inst *innerFileSystem) GetPathByURI(uri lang.URI) (Path, error) {
	scheme := uri.Scheme()
	if scheme == "file" {
		path := uri.Path()
		return inst.GetPath("/" + path), nil
	}
	return nil, errors.New("bad scheme:" + scheme)
}

func (inst *innerFileSystem) GetPath(path string) Path {
	sep := inst.Separator()
	pb := &util.PathBuilder{}
	pb.SetSeparator(sep)
	pb.AppendPath(path)
	path = pb.String()
	if sep == "/" {
		path = sep + path
	}
	//	path, _ = filepath.Abs(path)
	return &innerPath{
		core: inst.core,
		path: path,
	}
}

func (inst *innerFileSystem) Roots() []Path {
	roots := inst.core.platform.Roots()
	list := make([]Path, len(roots))
	for index := range list {
		path := roots[index]
		list[index] = inst.GetPath(path)
	}
	return list
}

func (inst *innerFileSystem) DefaultReadOptions() *Options {
	opt := innerFileSystemDefaultOptionsR
	if opt == nil {
		opt = &Options{}
		opt.Flag = os.O_RDONLY
		opt.Mode = 0
		innerFileSystemDefaultOptionsR = opt
	}
	return opt.Clone()
}

func (inst *innerFileSystem) DefaultWriteOptions() *Options {
	opt := innerFileSystemDefaultOptionsW
	if opt == nil {
		opt = &Options{}
		opt.Flag = os.O_WRONLY
		opt.Mode = os.ModePerm
		innerFileSystemDefaultOptionsW = opt
	}
	return opt.Clone()
}

func (inst *innerFileSystem) SetDefaultOptions(r *Options, w *Options) {
	if r != nil {
		innerFileSystemDefaultOptionsR = r.Normalize()
	}
	if w != nil {
		innerFileSystemDefaultOptionsW = w.Normalize()
	}
}

func (inst *innerFileSystem) Separator() string {
	return inst.core.separator
}

func (inst *innerFileSystem) SeparatorChar() rune {
	return inst.core.separatorChar
}

func (inst *innerFileSystem) PathSeparator() string {
	return inst.core.pathSeparator
}

func (inst *innerFileSystem) PathSeparatorChar() rune {
	return inst.core.pathSeparatorChar
}

// impl innerPath

func (inst *innerPath) Name() string {
	return filepath.Base(inst.path)
}

func (inst *innerPath) Path() string {
	return inst.path
}

func (inst *innerPath) Parent() Path {
	parent := inst.FileSystem().GetPath(inst.path + "/..")
	path1 := inst.path
	path2 := parent.Path()
	if path1 == path2 {
		return nil
	}
	return parent
}

func (inst *innerPath) URI() lang.URI {
	pb := &util.PathBuilder{}
	pb.AppendPath(inst.path)
	path, err := pb.Create("/", "")
	if err != nil {
		path = "/"
	}
	location := &url.URL{}
	location.Scheme = "file"
	location.Path = path
	return lang.CreateURI(location)
}

func (inst *innerPath) String() string {
	return inst.path
}

func (inst *innerPath) Exists() bool {
	return inst.GetMeta().Exists()
}

func (inst *innerPath) IsDir() bool {
	return inst.GetMeta().IsDir()
}

func (inst *innerPath) IsSymlink() bool {
	return inst.GetMeta().IsSymlink()
}

func (inst *innerPath) IsFile() bool {
	return inst.GetMeta().IsFile()
}

func (inst *innerPath) Mkdir() error {
	mode := os.ModeDir
	return os.Mkdir(inst.path, mode)
}

func (inst *innerPath) Mkdirs() error {
	mode := os.ModeDir
	return os.MkdirAll(inst.path, mode)
}

func (inst *innerPath) Delete() error {
	return os.Remove(inst.path)
}

func (inst *innerPath) CopyTo(target Path) error {

	src, err := os.Open(inst.path)
	if err != nil {
		return err
	}
	defer src.Close()

	flag := os.O_CREATE | os.O_WRONLY
	mode := os.ModePerm

	dst, err := os.OpenFile(target.Path(), flag, mode)
	if err != nil {
		return err
	}
	defer dst.Close()

	_, err = io.Copy(dst, src)
	return err
}

func (inst *innerPath) MoveTo(target Path) error {
	old := inst.path
	next := target.Path()
	return os.Rename(old, next)
}

func (inst *innerPath) LastModTime() int64 {
	return inst.GetMeta().LastModTime()
}

func (inst *innerPath) Size() int64 {
	return inst.GetMeta().Size()
}

func (inst *innerPath) FileSystem() FileSystem {
	return inst.core.fs
}

func (inst *innerPath) ListNames() []string {
	file, err := os.Open(inst.path)
	if err != nil {
		return []string{}
	}
	defer file.Close()
	names, err := file.Readdirnames(0)
	if err != nil {
		return []string{}

	}
	return names
}

func (inst *innerPath) ListPaths() []string {
	names := inst.ListNames()
	paths := make([]string, len(names))
	for index := range names {
		name := names[index]
		paths[index], _ = filepath.Abs(inst.path + "/" + name)
	}
	return paths
}

func (inst *innerPath) ListItems() []Path {
	names := inst.ListNames()
	paths := make([]Path, len(names))
	for index := range names {
		name := names[index]
		paths[index] = inst.GetChild(name)
	}
	return paths
}

func (inst *innerPath) GetChild(name string) Path {
	path := inst.path
	return inst.FileSystem().GetPath(path + "/" + name)
}

func (inst *innerPath) GetHref(name string) Path {
	if inst.IsFile() {
		path := inst.path
		return inst.FileSystem().GetPath(path + "/../" + name)
	}
	return inst.GetChild(name)
}

func (inst *innerPath) GetIO() FileIO {
	return &innerFileIO{path: inst}
}

func (inst *innerPath) GetMeta() FileMeta {
	mode := &innerFileMeta{}
	mode.load(inst.path)
	return mode
}

func (inst *innerPath) SetMeta(mode FileMeta) {
	// TODO
}

func (inst *innerPath) CreateFile(opt *Options) error {
	return inst.CreateFileWithSize(0, opt)
}

func (inst *innerPath) CreateFileWithSize(size int64, opt *Options) error {

	if inst.Exists() {
		return errors.New("the file exists")
	}

	opt = opt.Normalize()
	file, err := os.OpenFile(inst.path, opt.Flag, opt.Mode)
	if err != nil {
		return err
	}
	defer file.Close()
	if size <= 0 {
		return nil
	}

	const bufferSize int64 = 1024 * 4
	buffer := make([]byte, bufferSize)
	var count int64 = 0

	for count < size {
		todoSize := size - count
		buf := buffer
		if todoSize > bufferSize {
			todoSize = bufferSize
		}
		if todoSize < bufferSize {
			buf = buffer[0:todoSize]
		}
		cb, err := file.Write(buf)
		if err != nil {
			return err
		}
		count += int64(cb)
	}

	return nil
}
