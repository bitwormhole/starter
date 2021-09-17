package main

import (
	"github.com/bitwormhole/starter"
	"github.com/bitwormhole/starter/vlog"
)

func main() {

	// args := []string{"-a", "-r", "-g", "-s"}

	vlog.Info("src/debug/go")
	i := starter.InitApp()
	// i.SetArguments(args)
	i.Use(starter.Module()).Use(innerModule()).Run()
}
