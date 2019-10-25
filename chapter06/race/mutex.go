// Package race 这个示例程序展示如何使用互斥锁来定义一段需要同步访问的代码临界区资源的同步访问
package race

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	// counter 是所有 goroutine 都要增加其值的变量
	counter03 int

	// wg03 用来等待程序结束
	wg03 sync.WaitGroup

	// mutex 用来定义一段代码临界区
	mutex sync.Mutex
)

// RunMutex 互斥锁测试入口
func RunMutex() {
	// 计数加 2，表示要等待两个 goroutine
	wg03.Add(2)

	// 创建两个 goroutine
	go incrCounterSafe(1)
	go incrCounterSafe(2)

	// 等待 goroutine 结束
	wg03.Wait()
	fmt.Println("Final Counter03:", counter03)
}

// incrCounterSafe 使用互斥锁来同步并保证安全访问，增加包里 counter 变量的值
func incrCounterSafe(id int) {
	// 在函数退出时调用 Done 来通知 main 函数工作已经完成
	defer wg03.Done()

	for count := 0; count < 7; count++ {
		// 同一时刻只允许一个 goroutine 进入这个临界区
		// 使用大括号只是为了让临界区看起来更清晰，并不是必需的。
		mutex.Lock()
		{
			// 捕获 counter 的值
			value := counter03

			// 当前 goroutine 从线程退出，并放回到队列
			// 强制将当前 goroutine 退出当前线程后，调度器会再次分配这个 goroutine 继续运行。
			runtime.Gosched()

			// 增加本地 value 变量的值
			value++

			// 将该值保存回 counter
			counter03 = value
		}
		// 释放锁，允许其他正在等待的 goroutine 进入临界区
		mutex.Unlock()
	}
}
