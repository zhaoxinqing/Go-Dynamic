package main

import (
	"backend-go/public"
	"backend-go/public/logger"
	"backend-go/service"
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func main() {

	GetRememberWordsFromTerminal()

	logger.NewLogger(logger.LogOptions("./log")) // logger

	RunRandomInt() // RunRandomFile()

	fmt.Println("APP RUN 00000")

	fmt.Printf("%s - 本地时间.\n", time.Now().Format(public.TIME_FORMAT))

	time.Local = time.UTC // 全局时区设置

	fmt.Printf("%s - 设置全局`UTC`时间.\n", time.Now().Format(public.TIME_FORMAT))

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

	DemoString()
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

var rememberWordList []string

// GetRememberWordsFromTerminal ... 助记词
func GetRememberWordsFromTerminal() {

	for i := 1; i <= 12; i++ {
		fmt.Printf("请输入第 %d 个助记词：", i)
		input := bufio.NewScanner(os.Stdin)
		input.Scan()
		rememberWordList = append(rememberWordList, input.Text())
	}
	if len(rememberWordList) != 12 {
		fmt.Println("助记词位数不正确")
	}

	fmt.Printf("完整助记词：%v\n", rememberWordList)
}

type DemoInfo struct {
	Name    string  `json:"name"`
	Trading string  `json:"trading"`
	Price   float64 `json:"price"`
	Detail  string  `json:"detail"`
}

func DemoString() {
	// var Str = "this is a test"

	demoInfo := DemoInfo{
		Name:    "wanger",
		Trading: "购买",
		Price:   100.00,
		Detail:  time.Now().Format("2006.01.02 15:04:05"),
	}

	fmt.Printf("%+v\n", demoInfo)

	fmt.Printf("%v", demoInfo)

	jsonbytes, _ := json.Marshal(demoInfo)

	jsonStr := string(jsonbytes)

	fmt.Println(jsonStr)
}
