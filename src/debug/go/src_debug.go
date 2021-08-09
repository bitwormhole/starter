package main

import (
	"fmt"

	"github.com/bitwormhole/starter"
)

func main() {
	fmt.Println("src/debug/go")
	starter.InitApp().Use(starter.Module()).Run()
}
