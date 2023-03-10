package demo

import (
	"encoding/json"
	"fmt"
	"time"
)

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
