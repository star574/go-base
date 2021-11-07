package main

import (
	"fmt"
	"os"
	"time"
)

// 计时器与定时器
func main() {
	// 计时器 time.Second*5 执行周期
	ticker := time.NewTicker(time.Second * 5)

	// 定时器 只执行一次 延时任务 超时
	timer := time.NewTimer(time.Second * 3)

	// 模拟服务器
	//  底层会根据周期向ticker.C 通道发送时间戳 所以可以捕获到一个时间段 充当定时器

	timout := make(chan bool)
	go func() {
		// After 定时器类似 多久之后执行 time.Sleep()差不多 8秒后发送到通道
		// 只会执行一次
		<-time.After(time.Second * 8)
		timout <- true
	}()

	for {
		select {
		case <-ticker.C:
			fmt.Println("上报服务器数据.....")
			fmt.Printf("time.Now(): %v\n", time.Now())
		case <-timer.C:
			fmt.Printf("3秒定时器执行 %v\n", time.Now())

		case <-timout:
			fmt.Println("8秒超时!")
			close(timout)
			os.Exit(0)
		default:
			fmt.Println("服务器运行中")
			time.Sleep(2 * time.Second)
		}
	}
}
