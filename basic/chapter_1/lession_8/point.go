package main

import "fmt"

// 指针
// 指针是用来保存内存地址的特殊变量
// & 取地址
// * 解指针
func main() {
	var p *int
	var q *float64
	var c int
	fmt.Printf("p: %v\n", p)
	fmt.Printf("q: %v\n", q)
	fmt.Printf("c: %v\n", c)

	p = &c
	fmt.Printf("p: %v\n", p)
	fmt.Printf("*p=c: %v\n", *p == c)
	fmt.Printf("p=&c: %v\n", p == &c)

	fmt.Printf("p: %T\n", p)
	fmt.Printf("p: %p\n", p)

	fmt.Printf("nil: %p\n", p) // nil: 0x0

	// 定义并没有分配内存地址
	// 分配内存地址
	q = new(float64)
	fmt.Printf("*q: %v\n", *q)
}
