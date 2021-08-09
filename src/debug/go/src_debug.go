package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("src/debug/go")
	// starter.InitApp().Use(starter.Module()).Run()

	arg0 := os.Args[0]
	fmt.Println("args[0] =", arg0)

}
