package vlog

import "testing"

func TestDefaultLogger(t *testing.T) {

	Fatal("fatal")
	Error("error")
	Warn("warn")
	Info("i", "n", "f", "o")
	Debug("debug")
	Trace("trace")

}
