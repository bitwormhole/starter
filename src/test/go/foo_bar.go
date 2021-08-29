package main

import "github.com/bitwormhole/starter/vlog"

////////////////////////////////////////////////////////////////////////////////

type Foo struct {
	Items []*Bar
	Value int
}

func (inst *Foo) Begin() error {
	vlog.Debug("foo.begin()")
	return nil
}

func (inst *Foo) End() error {
	vlog.Debug("foo.end()")
	return nil
}

////////////////////////////////////////////////////////////////////////////////

type Bar struct {
	Owner *Foo
	Name  string
}

func (inst *Bar) Start() error {
	vlog.Info("bar.start()")
	return nil
}

func (inst *Bar) Stop() error {
	vlog.Info("bar.stop()")
	return nil
}

////////////////////////////////////////////////////////////////////////////////
