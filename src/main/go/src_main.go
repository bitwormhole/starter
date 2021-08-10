package main

import (
	"github.com/bitwormhole/starter"
)

func main() {
	// fmt.Println("src/main/go")
	i := starter.InitApp()
	i.Use(starter.Module())
	i.Run()
}
