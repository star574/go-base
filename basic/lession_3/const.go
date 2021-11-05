package main

import "fmt"

// 常量的定义 常量未使用不会报错，编译可以通过（黄色波浪线） 变量（红色波浪线）
// 常量可以不使用强制类型转换 在使用的时候会自动使用合适的类型
func main() {
	const a int = 1
	const (
		b        = 20
		c string = "hello"
	)
	const (
		d = iota
		e
		_
		f
	)
	const (
		g = iota
		h
	)
	fmt.Printf("a: %v\n", a)
	fmt.Printf("b: %v\n", b)
	fmt.Printf("c: %v\n", c)
	fmt.Printf("d: %v\n", d)
	fmt.Printf("e: %v\n", e)
	fmt.Printf("f: %v\n", f)
	fmt.Printf("g: %v\n", g)
	fmt.Printf("h: %v\n", h)
}
