package main

import (
	"context"
	"fmt"
	"time"
)

type Mykey string
type MyValue string

// context 并发编程
func main() {
	// 1. runtime.Goexit() 结束协程
	// 2. 利用 channel 结束
	// ch := make(chan bool)

	// go func() {
	// 	go func() {

	// 	}()
	// 	go func() {

	// 	}()
	// 	for {
	// 		fmt.Println("协程")
	// 		time.Sleep(time.Second * 1)
	// 		// runtime.Goexit()
	// 		select {
	// 		case <-ch:
	// 			fmt.Println("结束")
	// 			return
	// 		default:
	// 			fmt.Println("继续")
	// 		}
	// 	}
	// }()
	// <-time.After(time.Second * 3)
	// ch <- true
	// <-time.After(time.Second * 1)

	// 传递 context
	// context.Background() 默认上下文 实现接口就可以自定义上下文
	// 可取消上下文
	// 1.context.WithCancel
	ctx, cancel := context.WithCancel(context.Background())

	go test(ctx)

	<-time.After(time.Second * 5)
	cancel()
	<-time.After(time.Second * 1)

	// 2.定时结束 设置协程死亡时间 5秒后自动结束 不需要 cance()
	ctx, cancel = context.WithDeadline(context.Background(), time.Now().Add(time.Second*5))
	go test1(ctx)
	<-time.After(time.Second * 6)
	cancel()

	// 3.超时结束
	ctx, cancel = context.WithTimeout(context.Background(), 3*time.Second)
	go test2(ctx)
	<-time.After(time.Second * 4)
	cancel()

	// 4 没有取消
	var key Mykey = "key"
	var value MyValue = "value"
	context := context.WithValue(context.Background(), key, value)
	go test3(context, key)
	<-time.After(time.Second * 6)
}

func test(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("WithCancel 主协程取消，子协程同时结束")
			return
		default:
			fmt.Println("WithCancel 子协程运行。。。")
			time.Sleep(time.Second * 1)
		}
	}
}
func test1(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("WithDeadline 主协程取消，子协程同时结束")
			return
		default:
			fmt.Println("WithDeadline 子协程运行。。。")
			time.Sleep(time.Second * 1)
		}
	}
}

func test2(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("WithTimeout 主协程取消，子协程同时结束")
			return
		default:
			fmt.Println("WithTimeout 子协程运行。。。")
			time.Sleep(time.Second * 1)
		}
	}
}
func test3(ctx context.Context, key Mykey) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("WithValue 取消")
		default:
			fmt.Println("WithValue 子协程运行。。。", ctx.Value(key))
			time.Sleep(time.Second * 1)
		}
	}
}
