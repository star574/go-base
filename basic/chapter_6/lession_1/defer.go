package main

import (
	"fmt"
	"os"
)

// 错误处理
// defer 延迟调用 函数
// defer 在 return 之前会执行
// defer 在发生错误之后也会执行
// 参数是在defer定义是就已经计算完成了
// 使用操作系统的退出 defer不会执行
// os.Exit(0)
func main() {
	test()
	test1()
	WriterFile("./basic/chapter_6/lession_1/hello.txt", "hello world!\n")
}

func test() {
	defer fmt.Printf("i: %v\n", 10) // 最后输出
	defer fmt.Printf("i: %v\n", 11) // 优先输出 defer 是倒着执行的 类似于栈 先进后出
	for i := 0; i < 10; i++ {
		fmt.Printf("i: %v\n", i)
		if i == 1 {
			return
		}
	}
}
func test1() {
	// 倒序输出
	for i := 0; i < 10; i++ {
		defer fmt.Printf("i: %v\n", i)
	}
}

// 写入文件
func WriterFile(fileName string, content string) {
	// os.O_CREATE 没有就创建  O_APPEND 追加  0666 权限
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		// 抛出错误
		panic(err)
	}
	// 定义defer
	defer file.Close()
	file.WriteString(content)
}
