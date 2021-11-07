package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var sy sync.WaitGroup

// 使用互斥锁解决数据竞争(多个线程/协程同时操作同一个对象)的问题
// 互斥锁(保证线程/协程对内存的一致性访问 拿到锁才能修改数据 否则阻塞等待 会有性能影响
var lock sync.Mutex


// 数据竞争与锁
// 使用 go run -race .\race.go 可以看到数据竞争的操作
// 不使用锁 可以使用 atomic 进行原子性操作 性能更好
func main() {
	var x int32

	// 自增加1
	for i := 0; i < 1000; i++ {
		// 协程外add
		sy.Add(1)
		go func() {
			// 协程内 down
			defer sy.Done()
			// lock.Lock()
			// x++
            atomic.AddInt32(&x,1)
			// lock.Unlock()
		}()
	}
	// 自增减1
	for i := 0; i < 1000; i++ {
		sy.Add(1)
		go func() {
			defer sy.Done()
			// 加锁
			// lock.Lock()
			// x--
            atomic.AddInt32(&x,-1)
			// 释放锁
			// lock.Unlock()
		}()
	}
	sy.Wait()
	fmt.Printf("x: %v\n", x) // 0
}
