package main

import "fmt"

// 函数
func main() {
	fmt.Printf("add(10, 5): %v\n", add(10, 5))
	fmt.Printf("add2(10, 5): %v\n", add2(10, 5))
	// 匿名函数
	fmt.Printf("res %v\n", func(a, b int) int {
		return a + b
	}(1, 2))

	fmt.Printf("op(\"-\")(1, 2): %v\n", op("*")(1, 2))

	a, b := swap(1, 2)
	fmt.Printf("a: %v\n", a)
	fmt.Printf("b: %v\n", b)

	// 闭包交换地址
	func() {
		a, b = b, a
	}()

	c, d := func(a, b *int) (*int, *int) {
		return b, a
	}(&a, &b)

	fmt.Printf("&a: %v\n", &a)
	fmt.Printf("a: %v\n", a)
	fmt.Printf("&b: %v\n", &b)
	fmt.Printf("b: %v\n", b)
	fmt.Printf("c: %v\n", c)
	fmt.Printf("*c: %v\n", *c)
	fmt.Printf("d: %v\n", d)
	fmt.Printf("*d: %v\n", *d)

	fmt.Printf("test1(1, 2, 3, 4, 5): %v\n", test1(1, 2, 3, 4, 5))

}
func add(a, b int) int {
	defer func() {
		if recover() != nil {
			fmt.Printf("recover(): %v\n", recover())
		}
		fmt.Printf("%v+%v=%v\n", a, b, a+b)
	}()
	return a + b
}

func add2(a, b int) (r int) {
	defer func() {
		fmt.Printf("%v+%v=%v\n", a, b, a+b)
	}()
	r = a + b
	return r
}

// 一个匿名函数简单计算器
func op(s string) func(a, b int) int {

	switch s {
	case "+":
		return func(a, b int) int { return a + b }
	case "-":
		return func(a, b int) int { return a - b }
	case "*":
		return func(a, b int) int { return a * b }
	case "/":
		return func(a, b int) int {
			if b == 0 {
				fmt.Println("出现错误 0 ")
			} else {

				return a / b
			}
			return 0
		}
	case "%":
		return func(a, b int) int { return a % b }
	}
	return nil
}
func swap(a, b int) (int, int) {
	return b, a
}

// 任意参数长度的传递
func test1(a ...int) int {
	return test(a...)
}
func test(a ...int) int {
	s := 0
	for i := 0; i < len(a); i++ {
		s += a[i]
	}
	return s
}
