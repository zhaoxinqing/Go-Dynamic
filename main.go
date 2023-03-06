package main

import (
	"demo/backend"
	"demo/config"
	"demo/demo"
)

func main() {

	demo.Test()

	config.LoadConf()

	backend.Run(":8080")
}
