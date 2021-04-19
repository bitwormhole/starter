package config

import (
	"testing"
)

func TestDemo(t *testing.T) {
	code := Demo()
	t.Logf("exit with code: %d", code)
}
