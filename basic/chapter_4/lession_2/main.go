package main

import (
	"fmt"
    // 引入mysql依赖 _ 代表不用包里的东西 只要导入就会执行init函数
    _ "github.com/go-sql-driver/mysql"

	"github.com/star574/go-base/basic/chapter_4/lession_2/test"

)

// 包的初始化
// init函数在main函数执行之前执行
// 可以又多个init函数 按定义的上下顺序执行
// 只要导入 就算不用包里面的东西也会执行init函数 多个包导入按照导入顺序执行init函数

func main() {
	/* 包初始化 UserInfo init
	main函数执行
	test.UserInfo: {Id:1} */
	fmt.Println("main函数执行")
	fmt.Printf("test.UserInfo: %+v\n", *test.UserInfo)
	/*
		    包初始化 UserInfo init
			test.UserInfo: {Id:1}
	*/

}
