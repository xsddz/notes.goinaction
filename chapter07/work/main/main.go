// 这个示例程序展示如何使用 work 包创建一个 goroutine 池并完成工作
package main

import (
	"log"
	"sync"
	"time"

	"notes.goinaction/chapter07/work"
)

// names 提供了一组用来显示的名字
var names = []string{
	"张三",
	"李四",
	"王五",
	"赵六",
	"孔七",
	"孙八",
}

// namePrinter 使用特定方式打印名字
type namePrinter struct {
	name string
}

// Task 实现 Worker 接口
func (m *namePrinter) Task() {
	log.Println(m.name)
	time.Sleep(time.Second)
}

// main 程序入口
func main() {
	// 使用两个 goroutine 来创建工作池
	p := work.New(2)

	var wg sync.WaitGroup
	wg.Add(100 * len(names))

	for i := 0; i < 100; i++ {
		// 迭代 names 切片
		for _, name := range names {
			// 创建一个 namePrinter 并提供指定的名字
			np := namePrinter{
				name: name,
			}

			go func() {
				// 将任务提交执行。当 Run 返回时我们就知道任务已经处理完成
				p.Run(&np)
				wg.Done()
			}()
		}
	}

	wg.Wait()

	// 让工作池停止工作，等待所有现有的工作完成
	p.Shutdown()
}
