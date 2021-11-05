package main

import (
	"fmt"
	"math"
)

// 数据类型
func main() {
	var (
		b bool
		a string
		c int
		d uint
		e float32 = 3.14
		f float64
		g complex64 = 3 + 2i
	)
	fmt.Printf("a: %v\n", a)
	fmt.Printf("b: %v\n", b)
	fmt.Printf("c: %v\n", c)
	fmt.Printf("d: %v\n", d)
	fmt.Printf("e: %v\n", e)
	fmt.Printf("f: %v\n", f)
	fmt.Printf("g: %v\n", g)

	// 类型转换
	f = float64(e)
	fmt.Printf("f: %v\n", f)
	q := math.Sqrt(3*3 + 4*4)
	fmt.Printf("q: %v\n", q)

}
