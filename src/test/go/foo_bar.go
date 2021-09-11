package main

import "github.com/bitwormhole/starter/vlog"

////////////////////////////////////////////////////////////////////////////////

// Foo ...
type Foo struct {
	Items []*Bar
	Value int
}

// Begin ...
func (inst *Foo) Begin() error {
	vlog.Debug("foo.begin()")
	return nil
}

// End ...
func (inst *Foo) End() error {
	vlog.Debug("foo.end()")
	return nil
}

////////////////////////////////////////////////////////////////////////////////

// Bar ...
type Bar struct {
	Owner *Foo
	Name  string
}

// Start ...
func (inst *Bar) Start() error {
	vlog.Info("bar.start()")
	return nil
}

// Stop ...
func (inst *Bar) Stop() error {
	vlog.Info("bar.stop()")
	return nil
}

////////////////////////////////////////////////////////////////////////////////
