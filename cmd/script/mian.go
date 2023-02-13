package main

import (
	"demo/app"
	"fmt"
)

func main() {
	// script.JsonConv()
	// script.JsonConv02()
	// script.JsonConv3()
	// script.TTY() // es 数据迁移
	// script.ZhiHu()
	app.SendEmail("", "", nil)
	fmt.Println("ALL MISSION SUCCESS !!!")
}
