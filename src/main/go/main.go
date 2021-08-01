package main

import (
	"fmt"

	"github.com/bitwormhole/starter"
)

func main() {
	res := starter.GetResources()
	starter.Init().MountResources(res, "/").Use(nil).Run()
	fmt.Println("this is src/main/go")
}
