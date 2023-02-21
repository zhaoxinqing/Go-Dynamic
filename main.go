package main

import (
	"demo/app"
	"demo/backend"
	"demo/public"
	"demo/service"
	"fmt"
	"log"
	"strings"
	"time"
)

func main() {

	Test()

	service.WaitGroup()

	backend.Run(":8080")
}

func Test() {

	fmt.Printf("%s - 本地时间.\n", time.Now().Format(public.TIME_FORMAT))

	time.Local = time.UTC // 全局时区设置

	fmt.Printf("%s - 设置全局`UTC`时间.\n", time.Now().Format(public.TIME_FORMAT))

	app.Run() // APP

	service.Run() // SERVICE

	node, err := service.NewWorker(1) // snowflake
	if err != nil {
		log.Println(err)
	}

	log.Println(node.GetId())

	str1 := "IN THE WHOLE WORLD, I AM THE ONLY ONE"

	log.Println(strings.ToLower(str1))
	log.Println(strings.ToTitle(str1))

	log.Println("ALL MISSION SUCCESS !!!")
}
