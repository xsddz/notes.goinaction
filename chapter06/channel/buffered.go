/*
Package channel 这个示例程序展示如何使用有缓冲的通道和固定数目的 goroutine 来处理一堆工作

1. 有缓冲的通道（buffered channel）是一种在被接收前能存储一个或者多个值的通道。
这种类型的通道并不强制要求 goroutine 之间必须同时完成发送和接收。通道会阻塞发送和接收动作的条件
也会不同。只有在通道中没有要接收的值时，接收动作才会阻塞。只有在通道没有可用缓冲区容纳被发送的值时，
发送动作才会阻塞。这导致有缓冲的通道和无缓冲的通道之间的一个很大的不同：无缓冲的通道保证进行发送和
接收的 goroutine 会在同一时间进行数据交换；有缓冲的通道没有这种保证。
*/
package channel

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	numberGoroutines = 4  // 要使用的 goroutine 的数量
	taskLoad         = 10 // 要处理的工作的数量
)

// wg 用来等待程序完成
var wg02 sync.WaitGroup

// init 初始化包，Go 语言运行时会在其他代码执行之前优先执行这个函数
func init() {
	// 初始化随机数种子
	rand.Seed(time.Now().Unix())
}

// RunBuffered 有缓存的通道测试入口
func RunBuffered() {
	// 创建一个有缓冲的通道来管理工作
	tasks := make(chan string, taskLoad)

	// 启动 goroutine 来处理工作
	wg02.Add(numberGoroutines)
	for gr := 1; gr <= numberGoroutines; gr++ {
		go worker(tasks, gr)
	}

	// 增加一组要完成的工作
	for post := 1; post <= taskLoad; post++ {
		tasks <- fmt.Sprintf("Tasks : %d", post)
	}

	// 当所有工作都处理完时关闭通道，以便所有 goroutine 退出
	// 当通道关闭后，goroutine 依旧可以从通道接收数据，但是不能再向通道里发送数据。
	close(tasks)

	// 等待所有工作完成
	wg02.Wait()
}

// worker 作为 goroutine 启动来处理从有缓冲的通道传入的工作
func worker(tasks chan string, worker int) {
	// 通知函数已经返回
	defer wg02.Done()

	for {
		// 等待分配工作
		// 能够从已经关闭的通道接收数据这一点非常重要，因为这允许通道关闭后依旧能取出其中缓冲的全部值，
		// 而不会有数据丢失。从一个已经关闭且没有数据的通道里获取数据，总会立刻返回，并返回一个通道类型的零值。
		// 如果在获取通道时还加入了可选的标志，就能得到通道的状态信息。
		task, ok := <-tasks
		if !ok {
			// 这意味着通道已经空了，并且已被关闭
			fmt.Printf("Worker : %d : Shutting Down\n", worker)
			return
		}

		// 显示我们开始工作了
		fmt.Printf("Woker : %d : Started %s\n", worker, task)

		// 随机等一段时间来模拟工作
		sleep := rand.Int63n(100)
		time.Sleep(time.Duration(sleep) * time.Millisecond)

		// 显示我们完成了工作
		fmt.Printf("Worker : %d : Completed %s\n", worker, task)
	}
}
