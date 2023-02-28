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

// 协程计算求和并将结果汇总
func WaitGroup02() {
	// 初始化一个切片
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// 计算每个协程的工作量
	chunkSize := len(nums) / 4
	remainder := len(nums) % 4

	// 使用 WaitGroup 来等待所有协程完成
	var wg sync.WaitGroup
	wg.Add(4)

	// 使用 channel 来收集协程结果
	resultCh := make(chan int)

	// 使用协程计算每个 chunk 的求和
	for i := 0; i < 4; i++ {
		start := i * chunkSize
		end := (i + 1) * chunkSize
		if i == 3 {
			end += remainder
		}

		go func(nums []int) {
			sum := 0
			for _, n := range nums {
				sum += n
			}
			resultCh <- sum
			wg.Done()
		}(nums[start:end])
	}

	// 等待所有协程完成并收集结果
	go func() {
		wg.Wait()
		close(resultCh)
	}()

	// 汇总所有结果
	totalSum := 0
	for sum := range resultCh {
		totalSum += sum
	}

	// 输出结果
	fmt.Printf("Total sum: %d\n", totalSum)
}

// 在这个示例代码中，我们首先初始化一个整数切片。
// 然后，我们计算出每个协程的工作量，即切片的一部分。
// 我们使用 WaitGroup 来等待所有协程完成，并使用 channel 来收集每个协程的结果。
// 接下来，我们使用协程计算每个 chunk 的求和，并将结果发送到 channel 中。
// 最后，我们等待所有协程完成并收集结果，然后将所有结果汇总。最终，我们输出总和。

// 请注意，这里我们使用了 sync.WaitGroup 来等待所有协程完成，并使用了一个额外的协程来关闭结果 channel。
// 这是因为在没有额外协程的情况下，我们无法确定所有协程都已经完成并将结果发送到 channel 中。
// 因此，我们需要在一个额外的协程中等待所有协程完成，并关闭 channel 来通知主协程已经完成了所有工作。
