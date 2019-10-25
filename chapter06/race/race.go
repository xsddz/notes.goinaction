// Package race 这个示例程序展示如何在程序里造成竞争状态，实际上不希望出现这种情况
package race

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	// counter01 是所有 goroutine 都要增加其值的变量
	counter01 int

	// wg01 用来等待程序结束
	wg01 sync.WaitGroup
)

// RunRace 竞争测试入口
func RunRace() {
	// 计数加 2，表示要等待两个 goroutine
	wg01.Add(2)

	// 创建两个 goroutine
	go incrCounterUnsafe(1)
	go incrCounterUnsafe(2)

	// 等待 goroutine 结束
	wg01.Wait()
	fmt.Println("Final Counter01:", counter01)
}

// incrCounterUnsafe 增加包里 counter01 变量的值
func incrCounterUnsafe(id int) {
	// 在函数退出时调用 Done 来通知 main 函数工作已经完成
	defer wg01.Done()

	for count := 0; count < 7; count++ {
		// 捕获 counter01 的值
		value := counter01

		/*
			当前 goroutine 从线程退出，并放回到队列

			1. 调用了 runtime 包的 Gosched 函数，用于将 goroutine 从当前线程退出，
			给其他 goroutine 运行的机会。在两次操作中间这样做的目的是强制调度器切换两个
			goroutine，以便让竞争状态的效果变得更明显。
			2. go build -race 用竞争检测器标志来编译程序
		*/
		runtime.Gosched()

		// 增加本地 value 变量的值
		value++

		// 将该值保存回 counter01
		counter01 = value
	}
}
