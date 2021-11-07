package main

import (
	"fmt"

	"github.com/star574/go-base/basic/chapter_3/lession_2/brid"
)

// 结构体的方法
// 结构体 中的属性 首字母大写表示其他包可以访问 结构体 方法 同理

func main() {
	b1 := brid.Brid1{
		Name: "喜鹊",
	}
	b1.Fly()
	fmt.Printf("b1.name: %v\n", b1.Name)
	// 匿名调用
	var b2 = brid.Brid1{Name: "斑鸠"}
	b2.Fly()

	b3 := brid.Brid2{
		Name: "狗",
	}
	b3.Run()
}
