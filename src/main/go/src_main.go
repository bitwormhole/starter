package main

import (
	"github.com/bitwormhole/starter"
	"github.com/bitwormhole/starter/vlog"
)

func main() {
	vlog.Debug("src/main/go")
	i := starter.InitApp()
	i.Use(starter.Module())
	i.Run()
}
