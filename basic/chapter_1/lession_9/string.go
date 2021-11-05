package main

import (
	"fmt"
	"unicode/utf8"
)

// len函数计算长度（返回的是字节数而不是字符个数）
// 中文占3个字符(非拉丁字符)
// byte 字符
// rune unicode
// 基础包 utf8
func main() {
	s := "中文 hello world"
	println("--------------------------------")
	fmt.Printf("check(s): %v\n", check(s))
	println("--------------------------------")
	var c byte = 'A'
	fmt.Printf("c: %v\n", c) //65
	fmt.Printf("c: %c\n", c) //A

	// type rune = int32 （int32的缩写）
	// 使用四个字节保存
	var ch rune = '中'
	fmt.Printf("ch: %v\n", ch) // 20013
	fmt.Printf("ch: %c\n", ch) // 中

	// 计算字符个数
	fmt.Printf("utf8.RuneCountInString(s): %v\n", utf8.RuneCountInString(s)) // utf8.RuneCountInString(s): 14

}

func check(s string) bool {
	// for range
	// var v rune
	for _, v := range s {
		fmt.Printf("v: %v-%c\n", v, v)
		fmt.Printf("v: %T\n", v) // int32
		fmt.Printf("v: %U\n", v) // v: U+0064
	}

	return utf8.RuneCountInString(s) > 20
}
