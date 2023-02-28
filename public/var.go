package public

import (
	"fmt"
	"reflect"
)

func VarInit() {

	// 初始化整数变量，值为10。
	var v int = 10
	fmt.Println(v)
	// 输出: 10

	// 变量声明: 一个slice变量
	var vSlice []int = []int{1, 2, 3, 4}
	fmt.Println(vSlice, "type: ", reflect.TypeOf(vSlice).Kind())
	// 输出: [1 2 3 4] type: slice

	// 短变量声明: 一个map变量，指向的值为[]
	var vMap map[string]int = map[string]int{
		"a": 1,
		"b": 2,
	}
	fmt.Println(vMap)
	// 输出: map[a:1 b:2]

	// 短变量声明: 一个整数变量。
	sdvInt := 10
	fmt.Println(sdvInt, "type: ", reflect.TypeOf(sdvInt).Kind())
	// 输出: 10 type:  int

	// 短变量声明: 一个slice变量
	sdvSlice := []int{1, 2, 3, 4}
	fmt.Println(sdvSlice, "type: ", reflect.TypeOf(sdvSlice).Kind())
	// 输出: [1 2 3 4] type: slice

	// 短变量声明: 一个map变量，指向的值为[]
	sdvMap := map[string]int{
		"a": 1,
		"b": 2,
	}
	fmt.Println(sdvMap)
	// 输出: map[a:1 b:2]

	// 初始化一个整数指针变量，指向的值为0
	var newInt *int = new(int)
	fmt.Println(*newInt)

	// 初始化一个slice指针变量
	var newSlice = new([10]int)[0:5]
	fmt.Println(newSlice, "type: ", reflect.TypeOf(newSlice).Kind())
	// 输出: [0 0 0 0 0] type: slice

	// 初始化一个map指针变量，指向的值为[]
	var newMap *map[string]int = new(map[string]int)
	fmt.Println(*newMap)
	// 输出: map[]

	// 初始化一个chan指针变量，指向的值为nil
	var newChan *chan int = new(chan int)
	fmt.Println(*newChan)
	// 输出: nil

	// make只能用于创建slice, map, channel
	// 切片类型(slice)
	makeSlice := make([]int, 5, 10)
	fmt.Println(makeSlice)
	// 输出: [0 0 0 0 0]

	// Map 类型
	var makeMap map[string]int = make(map[string]int)
	fmt.Println(makeMap)
	// 输出: map[]

	// Channel 类型
	var makeChan chan int32 = make(chan int32, 100)
	fmt.Println(makeChan)
	// 输出: 0xc000112000

}
