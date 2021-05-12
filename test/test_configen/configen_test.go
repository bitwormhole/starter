package test_configen

import (
	"testing"

	"github.com/bitwormhole/starter/tools/configen2"
)

func TestConfigen2(t *testing.T) {

	// pwd :=  os.Getenv("PWD")

	pwd := "D:\\home\\xukun\\git\\wormhole2020projects\\modules\\starter\\demo\\demo-for-config\\conf"

	runner := &configen2.Runner{}
	runner.RunWithPWD(pwd)
}
