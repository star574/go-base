package main

import "fmt"

// 运算符
func main() {
	// 匿名函数
	a, b := 1, 2
	fmt.Println(func() int {
		return a * b
	}())
    
}
