package main

import (
	"github.com/bitwormhole/starter"
	"github.com/bitwormhole/starter/application"

	"github.com/bitwormhole/starter/vlog"
)

func main() {

	vlog.Info("src/debug/go")

	i := starter.InitApp()
	i.Use(starter.Module())
	i.UseMain(innerModule())

	rt, err := i.RunEx()
	if err != nil {
		panic(err)
	}
	err = run(rt.Context())
	if err != nil {
		panic(err)
	}
	rt.Exit()
}

func run(ctx application.Context) error {

	//os.Stdout.WriteString("")
	//os.Stdin.SetReadDeadline(  )
	return nil
}
