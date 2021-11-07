package main

import "fmt"

// select 管理 channel
func main() {
	ch1 := make(chan int)
	go func() {
		sum := 0
		for i := 1; i < 1000; i++ {
			sum += i
			// 发送数据到通道
			ch1 <- sum
		}
		close(ch1)
	}()

	// 1. 死循环读取数据
	// 2. 在另一个循环读取值
	for {
		value, ok := <-ch1
		if ok {
			fmt.Printf("value: %v\n", value)
		} else {
			break
		}
	}
	fmt.Println("结束")

	// 3 for range
	for value := range ch1 {
		fmt.Println("range ch1=", value)
	}

	// 前面几种无法处理多个通道
	// 用select对channel进行管理
	ch2 := make(chan int)
	go func() {
		for i := 0; i < 1000; i++ {
			ch2 <- i + 2
		}
	}()

	// 多个分支 能同时读到数据 那么会随机读取一个
	// 如果其中一个读不到 会读取另一个
	// 如果全部读不到
	// 一般配合死循环 会无限等待 break无法生效 os.Exit(0)可以退出
	for {
		select {
		// 必须对channel进行io操作(读,写)
		case value, ok := <-ch1:
			if ok {
				fmt.Println("ch1 value", value)
			}
		case value, ok := <-ch2:
			if ok {
				fmt.Println("ch2 value", value)
			}
		default:
			fmt.Println("未读到数据")
		}
	}
}
