package main

import (
	"backend-go/service"
	"fmt"
)

func main() {
	// RunRandomFile()
	RunRandomInt()
}

// RunRandomFile ... 抽文件
func RunRandomFile() {
	str, err := service.RandomFile("./docs")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(str)
}

// RunRandomFile ... 抽数字
func RunRandomInt() {
	num := service.RandomInt(10000)
	fmt.Println(num)
}
