package config

import (
	"testing"
)

func TestDemo(t *testing.T) {
	code := Demo(nil, "")
	t.Logf("exit with code: %d", code)
}
