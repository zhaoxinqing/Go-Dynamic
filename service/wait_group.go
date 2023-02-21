package service

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup //定义一个同步等待的组

func task(i int) {
	fmt.Println("task...", i)
	//耗时操作任务，网络请求，读取文件
	time.Sleep(time.Second)
	wg.Done() //减去一个计数
}

func WaitGroup() {

	start := time.Now()
	for i := 0; i < 10; i++ {
		wg.Add(1) //添加一个计数
		go task(i)
	}
	wg.Wait() //阻塞直到所有任务完成

	tc := time.Since(start) //计算耗时

	fmt.Printf("time cost = %v\n", tc)
	fmt.Println("over")
}
