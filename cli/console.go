package cli

import (
	"context"
	"errors"
	"io"
	"os"

	"github.com/bitwormhole/starter/contexts"
	"github.com/bitwormhole/starter/io/fs"
)

////////////////////////////////////////////////////////////////////////////////

// Console 接口表示一个跟上下文绑定的控制台
type Console interface {
	GetWD() string
	SetWD(wd string)

	GetWorkingPath() fs.Path
	SetWorkingPath(wd fs.Path)

	Error() io.Writer
	Output() io.Writer
	Input() io.Reader

	SetError(w io.Writer)
	SetOutput(w io.Writer)
	SetInput(r io.Reader)
}

// ConsoleFactory 是创建 Console 的工厂
type ConsoleFactory interface {
	Create() Console
}

////////////////////////////////////////////////////////////////////////////////

// GetConsole  从上下文取控制台接口
func GetConsole(ctx context.Context) (Console, error) {
	holder, err := getConsoleHolder(ctx)
	if err != nil {
		return nil, err
	}
	console := holder.getConsole(true)
	if console == nil {
		return nil, errors.New("(create)console==nil")
	}
	return console, nil
}

// SetupConsole 从上下文取控制台接口
func SetupConsole(ctx context.Context, factory ConsoleFactory) error {
	holder, err := getConsoleHolder(ctx)
	if err != nil {
		return err
	}
	holder.factory = factory
	return nil
}

////////////////////////////////////////////////////////////////////////////////

type consoleHolder struct {
	console Console
	factory ConsoleFactory
}

func getConsoleHolder(ctx1 context.Context) (*consoleHolder, error) {

	const key = "/bitwormhole/starter/cli/consoleHolder#binding"

	o1 := ctx1.Value(key)
	o2, ok := o1.(*consoleHolder)
	if ok {
		return o2, nil
	}

	setter, err := contexts.GetContextSetter(ctx1)
	if err != nil {
		return nil, err
	}

	holder := &consoleHolder{}
	setter.SetValue(key, holder)
	return holder, nil
}

func (inst *consoleHolder) getConsole(create bool) Console {
	console := inst.console
	if console == nil {
		if create {
			console = inst.createConsole()
		}
	}
	return console
}

func (inst *consoleHolder) createConsole() Console {
	factory := inst.factory
	if factory == nil {
		console := &consoleImpl{}
		console.init()
		return console
	}
	return factory.Create()
}

////////////////////////////////////////////////////////////////////////////////

type consoleImpl struct {
	pwd string

	in  io.Reader
	out io.Writer
	err io.Writer
}

func (inst *consoleImpl) _Impl() Console {
	return inst
}

func (inst *consoleImpl) Input() io.Reader {
	return inst.in
}

func (inst *consoleImpl) Output() io.Writer {
	return inst.out
}

func (inst *consoleImpl) Error() io.Writer {
	return inst.err
}

func (inst *consoleImpl) SetInput(s io.Reader) {
	if s == nil {
		return
	}
	inst.in = s
}

func (inst *consoleImpl) SetOutput(s io.Writer) {
	if s == nil {
		return
	}
	inst.out = s
}

func (inst *consoleImpl) SetError(s io.Writer) {
	if s == nil {
		return
	}
	inst.err = s
}

func (inst *consoleImpl) init() error {
	dir, err := os.Getwd()
	if err != nil {
		return err
	}
	inst.pwd = dir
	return nil
}

func (inst *consoleImpl) GetWD() string {
	if inst.pwd == "" {
		inst.init()
	}
	return inst.pwd
}

func (inst *consoleImpl) SetWD(dir string) {
	if dir == "" {
		return
	}
	inst.pwd = dir
}

func (inst *consoleImpl) GetWorkingPath() fs.Path {
	path := fs.Default().GetPath(inst.pwd)
	p := path
	for timeout := 99; p != nil; p = p.Parent() {
		if timeout > 0 {
			timeout--
		} else {
			break
		}
		if p.IsDir() {
			return p
		}
	}
	return path
}

func (inst *consoleImpl) SetWorkingPath(path fs.Path) {
	if path == nil {
		return
	}
	inst.pwd = path.Path()
}