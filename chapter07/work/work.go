// Package work 包管理一个 goroutine 池来完成工作
package work

import "sync"

// Worker 必须满足接口类型，才能使用工作池
type Worker interface {
	Task()
}

// Pool 提供一个 goroutine 池，这个池可以完成任何已提交的 Worker 任务
type Pool struct {
	work chan Worker
	wg   sync.WaitGroup
}

// New 创建一个新工作池
func New(maxGoroutines int) *Pool {
	p := Pool{
		work: make(chan Worker),
	}

	p.wg.Add(maxGoroutines)
	for i := 0; i < maxGoroutines; i++ {
		go func() {
			// 1. for range 循环会一直阻塞，直到从 work 通道收到一个 Worker 接口值
			// 2. 一旦 work 通道被关闭，for range 循环就会结束，并调用 WaitGroup 的 Done 方法。然后 goroutine 终止
			for w := range p.work {
				w.Task()
			}
			p.wg.Done()
		}()
	}

	return &p
}

/*
Run 提交工作到工作池

由于 work 通道是一个无缓冲的通道，调用者必须等待工作池里的某个 goroutine 接收到这个值才会返回。
这正是我们想要的，这样可以保证调用的 Run 返回时，提交的工作已经开始执行。
*/
func (p *Pool) Run(w Worker) {
	p.work <- w
}

// Shutdown 等待所有 goroutine 停止工作
func (p *Pool) Shutdown() {
	close(p.work)
	p.wg.Wait()
}
