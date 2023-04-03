package app

import (
	"context"
	"fmt"
	"time"
)

// snippet[ˈsnɪpɪt]小片，片段; 不知天高地厚的年轻人

func Gen() <-chan int {
	ch := make(chan int)
	go func() {
		var n int
		for {
			ch <- n
			n++
			time.Sleep(time.Second)
		}
	}()
	return ch
}

func Perform(ctx context.Context) {
	for {
		// calculatePos()
		// sendResult()

		time.Sleep(time.Second * 2)
		fmt.Println(time.Now().Format("2006.01.02 15:04:05"))

		select {
		case <-ctx.Done():
			// 被取消，直接返回
			return
		case <-time.After(time.Second):
			// block 1 秒钟
		}
	}
}

func RunPerform() {
	fmt.Println("start-------------------------")
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	Perform(ctx)

	// ……
	// app 端返回页面，调用cancel 函数
	cancel()

	fmt.Println("end-------------------------")

	select {}

}

// start-------------------------
// 2023.04.03 18:39:17
// 2023.04.03 18:39:20
// 2023.04.03 18:39:23
// 2023.04.03 18:39:26
// 2023.04.03 18:39:30
// 2023.04.03 18:39:33
// 2023.04.03 18:39:36
// 2023.04.03 18:39:39
// 2023.04.03 18:39:42
// 2023.04.03 18:39:45
// 2023.04.03 18:39:48
// 2023.04.03 18:39:51
// 2023.04.03 18:39:54
// 2023.04.03 18:39:57
// 2023.04.03 18:40:00
// 2023.04.03 18:40:03
// 2023.04.03 18:40:06
// 2023.04.03 18:40:09
// 2023.04.03 18:40:12
// 2023.04.03 18:40:15
// end-------------------------
