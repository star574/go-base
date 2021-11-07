package main

import "sync"

// 并发编程
// 并发 cpu线程之间来回切换(同时处理的还是一个任务)
// 并行 两个独立并同时运行的任务
// 协程 : 轻量级的线程 go语言运行时调度 用户态 (线程在linux中是内核态)
// 同步 WaitGroup: 保存协程计数

var wg sync.WaitGroup

func main() {
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go test(i)
	}
	// 直到计数为0 也就是所有协程执行完毕 才会放行
	wg.Wait()
}

func test(i int) {
	// 通知协程运行完毕 协程减1
	defer wg.Done()
	println(i)
}
