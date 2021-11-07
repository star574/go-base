package main

import (
	"fmt"
)

type Myint int    // 自定义类型 Myint  基于int
type Myint2 = int // 别名

type logger struct {
	Desc string
}

func (log *logger) PrintLog() {
	fmt.Printf("log: %+v\n", log)
}

// 自定义类型扩展

type MyLogger logger

// 类似与iop
func (log *MyLogger) PrintLog1() {
	log.Desc = "啦啦啦"
	fmt.Printf("修改后log: %+v\n", log) // 修改后log: &{Desc:哈哈哈哈哈}
}

// 自定义类型
// 类型别名
// 自定义类型扩展
func main() {
	// 不能使用短语法
	// i: 1 main.Myint
	// i2: 2 int
	var i Myint = 1
	fmt.Printf("i: %v %T\n", i, i)
	var i2 Myint2 = 2
	fmt.Printf("i2: %+v %T\n", i2, i2)

	// 扩展后log
	var log = MyLogger{Desc: "哈哈哈哈哈"}
	log.PrintLog1()
	fmt.Printf("log.Desc: %v\n", log.Desc)

}
