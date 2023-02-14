package main

import (
	"demo/app"
	"demo/service"
	"fmt"
	"strings"
)

func main() {

	app.Run() // APP

	service.Run() // SERVICE

	node, err := service.NewWorker(1) // snowflake
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(node.GetId())

	str1 := "IN THE WHOLE WORLD, I AM THE ONLY ONE"

	fmt.Println(strings.ToLower(str1))
	fmt.Println(strings.ToTitle(str1))

	fmt.Println("ALL MISSION SUCCESS !!!")
}
