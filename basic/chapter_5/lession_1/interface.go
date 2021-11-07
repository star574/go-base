package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
)

type User struct {
	Id   int    `json:"id" xml:"id,attr"`
	Name string `json:"name" xml:"name,attr"`
}

type JsonParser string

type XmlParser string

func (p JsonParser) Serialize(v interface{}) {
	b, err := json.Marshal(v)
	if err != nil {
		// 抛出错误
		panic(err)
	} else {
		fmt.Printf("%s\n", b)
	}
}
func (p XmlParser) Serialize(v interface{}) {
	// 序列化成xml
	b, err := xml.Marshal(v)
	if err != nil {
		// 抛出错误
		panic(err)
	} else {
		fmt.Printf("%s\n", b)
	}
}

// 接口优化后
// 上面两个Serialize 实现接口的Serialize
type TokenParser interface {
	Serialize(v interface{})
}

func printAny(v interface{}, p TokenParser) {
	p.Serialize(v)
}

// 接口
// 接口的好处:
// 对业务逻辑更高层次的抽象
// 更加优雅与灵活
func main() {
	user := User{
		Id:   101,
		Name: "王五",
	}

	var p JsonParser
	var x XmlParser
	p.Serialize(user)
	x.Serialize(user)
	fmt.Println("=========================")
	// 使用接口后
	printAny(user, p)
	printAny(user, x)

}
