// Package tgoroutine 这个示例程序展示如何创建 goroutine以及调度器的行为
package tgoroutine

import (
	"fmt"
	"runtime"
	"sync"
)

/*
Run goroutine测试入口

1. Go 语言里的并发指的是能让某个函数独立于其他函数运行的能力。当一个函数创建为 goroutine 时，
Go 会将其视为一个独立的工作单元。这个单元会被调度到可用的逻辑处理器上执行。
2. 操作系统会在物理处理器上调度线程来运行，而 Go 语言的运行时会在逻辑处理器上调度 goroutine来运行。
每个逻辑处理器都分别绑定到单个操作系统线程。在 1.5 版本上，Go语言的运行时默认会为每个可用的物理处理器
分配一个逻辑处理器。
3. 如果创建一个 goroutine 并准备运行，这个 goroutine 就会被放到调度器的全局运行队列中。
之后，调度器就将这些队列中的 goroutine 分配给一个逻辑处理器，并放到这个逻辑处理器对应的
本地运行队列中。本地运行队列中的 goroutine 会一直等待直到自己被分配的逻辑处理器执行。
4. 并行的关键是同时做很多事情，而并发是指同时管理很多事情，这些事情可能只做了一半就被暂停去做别的事情了。
5. 只有在有多个逻辑处理器且可以同时让每个 goroutine 运行在一个可用的物理处理器上的时候，goroutine 才会并行运行。
6. 基于调度器的内部算法，一个正运行的 goroutine 在工作结束前，可以被停止并重新调度。
调度器这样做的目的是防止某个 goroutine 长时间占用逻辑处理器。当 goroutine 占用时间过长时，
调度器会停止当前正运行的 goroutine，并给其他可运行的 goroutine 运行的机会。
*/
func Run() {
	// 分配一个逻辑处理器给调度器使用
	runtime.GOMAXPROCS(1)
	// 给每个可用的核心分配一个逻辑处理器
	// runtime.GOMAXPROCS(runtime.NumCPU())

	// wg 用来等待程序完成
	// 计数加 2，表示要等待两个 goroutine
	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("Start Goroutines")

	// 声明一个匿名函数，并创建一个 goroutine
	go func() {
		// 在函数退出时调用 Done 来通知 main 函数工作已经完成
		defer wg.Done()

		// 显示字母表 3 次
		for count := 0; count < 3; count++ {
			for char := 'a'; char < 'a'+26; char++ {
				fmt.Printf("%c ", char)
			}
		}
	}()

	// 声明一个匿名函数，并创建一个 goroutine
	go func() {
		// 在函数退出时调用 Done 来通知 main 函数工作已经完成
		defer wg.Done()

		// 显示 5000 以内的素数值
		for outer := 2; outer < 5000; outer++ {
			isPrime := true
			for inner := 2; inner < outer; inner++ {
				if outer%inner == 0 {
					isPrime = false
					break
				}
			}
			if isPrime {
				fmt.Printf("%d ", outer)
			}
		}
	}()

	// 等待 goroutine 结束
	// WaitGroup 是一个计数信号量，可以用来记录并维护运行的 goroutine 。如果 WaitGroup 的值大于 0 ， Wait 方法就会阻塞。
	fmt.Println("Waiting To Finish")
	wg.Wait()

	fmt.Println("\nTerminating Program")
}
