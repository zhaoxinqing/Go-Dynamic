package main

import (
	"demo/app"
	"demo/service"
	"fmt"
	"strings"
)

func main() {
	// APP
	app.Run()

	// SERVICE
	service.Run()

	// snowflake
	node, err := service.NewWorker(1)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(node.GetId())

	str1 := "In the whole world, I am the only one"

	fmt.Println(strings.ToLower(str1))
	fmt.Println(strings.ToTitle(str1))

	fmt.Println("ALL MISSION SUCCESS !!!")
}
