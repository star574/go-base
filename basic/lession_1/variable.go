package main

import "fmt"

// 变量的定义
func main() {
	// 1
	var a int = 1
	// 2
	var l = 1
	var name string = "张三"
	// 3
	var b int
	// 4
	c := 1
	d := "hello world!"
	// 5
	e, f := 1, "hello"
	var (
		g int = 1
		h string
	)
	var (
		i, j string = "i", "k"
	)
	fmt.Printf("a: %v\n", a)
	fmt.Printf("name: %v\n", name)
	fmt.Printf("b: %v\n", b)
	fmt.Printf("c: %v\n", c)
	fmt.Printf("d: %+v\n", d)
	fmt.Printf("e: %+v\n", e)
	fmt.Printf("f: %+v\n", f)
	fmt.Printf("g: %+v\n", g)
	fmt.Printf("h: %+v\n", h)
	fmt.Printf("i: %+v\n", i)
	fmt.Printf("j: %+v\n", j)
	fmt.Printf("l: %+v\n", l)

}
