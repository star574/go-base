// 包名 
package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"sync"
	"time"
)

// 同步等待
var sy sync.WaitGroup

// channel 通道 应用类型 没有初始化无法使用 默认值 nil
// channel 又两种 有缓冲 无缓冲
func main() {
	// 不常用
	// int 通道中的元素类型
	// var ch chan int
	// fmt.Printf("未初始化: %v\n", ch) // 未初始化: <nil>

	// // 初始化
	// ch = make(chan int)
	// fmt.Printf("初始化: %v\n", ch) //初始化: 0xc00001a0c0

	// 常用
	ch := make(chan int)

	go func() {
		//  无缓冲通道 发送的同时必须接收 同步的通道
		//  有时候接收到消息但是来不及打印
		fmt.Println("接收数据:", <-ch)
		sy.Done()
	}()
	// 使用
	// 向通道发送数据
	sy.Add(1)
	ch <- 100
	sy.Wait()
	// 有缓冲的通道
	ch2 := make(chan int, 2)
	// 有缓冲通道即使没有接收者也不会报错 不能超出缓冲区大小 不是一个同步的通道
	// 缓冲区满了不能发送数据 缓冲为空时也不能读取数据
	ch2 <- 1
	ch2 <- 2
	fmt.Println(<-ch2) // 1
	fmt.Println(<-ch2) // 2 先进先出类似与队列

	// 关闭通道
	// 关闭之后就无法发送数据了
	// 可以从已关闭的通道读取数据 此时缓冲区已经为空了 默认值为数据类型默认值 0 nil
	// 不能重复关闭
	close(ch)
	close(ch2)

	// 模拟比特币挖矿
	// 使用时间生成随机数
	// 通道存放挖矿成功的数据
	blockCh := make(chan string)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 50000; i++ {
		// 先 sy.Add(1) 否则会偶尔出现  WaitGroup 报错问题
		sy.Add(1)
		go Miner(i, blockCh)
	}
	var slice = []string{}
	// 闭包
	go func(ch chan string) {
		// 死循环 不友好
		// for {
		// 	slice = append(slice, <-blockCh)
		// }
		// 建议使用 range
		for v := range ch {
			slice = append(slice, v)
		}
	}(blockCh)
	sy.Wait()
	close(blockCh)
	fmt.Printf("slice: %+#v  成功人数: %d\n", slice, len(slice))
}

// 常量魔法值
const MagicNum = "88"

// 模拟挖矿
func Miner(n int, ch chan string) {
	defer sy.Done()
	num := rand.Int()
	// 模拟挖矿成功
	// int转string
	numStr := strconv.Itoa(num)
	// 判断前缀 strings.HasPrefix
	if strings.HasPrefix(numStr, MagicNum) {
		// fmt.Println("矿工编号: ", n, "挖矿成功: ", num)
		ch <- numStr
	}
	// } else {
	// 	fmt.Println("矿工编号: ", n, "计算hash: ", num)
	// }
}
