package vlog

import "testing"

func TestDefaultLogger(t *testing.T) {
	Debug("a", "b", 6, "c")
}
