package main

import (
	"github.com/bitwormhole/starter"
	"github.com/bitwormhole/starter/vlog"
)

func main() {
	vlog.Info("src/debug/go")
	starter.InitApp().Use(starter.Module()).Use(innerModule()).Run()
}
