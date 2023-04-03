package main

import (
	"backend-go/app/demo"
	"fmt"
)

func main() {

	var a, b uint64
	a = 4
	b = 10

	c := a - b
	fmt.Println(c)
	demo.DemoString()

	// demo.Test()

	// config.LoadConf()

	// backend.Run(":8080")
}
