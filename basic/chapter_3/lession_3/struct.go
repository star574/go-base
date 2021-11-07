package main

import (
	"encoding/json"
	"fmt"
)

// 序列化字段不能为私有的 否在无法序列化私有字段
type User struct {
	// 标签 `json:"id"` id 可以任意定义 最好以原属性一致
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
}
 
// 结构体标签
// 结构体的序列化与反序列化
func main() {
	// 初始化
	user := new(User)
	user.Id = 1
	user.Name = "张三"
	user.Avatar = "https://www.baidu.com"

	// 返回字节数组与错误
	u, err := json.Marshal(user)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	// 字节数组 %s以字符串打印 u:
	fmt.Printf("u: %s\n", u)

	// 反序列化
	jsonStr := `{"id":20,"name":"李四","avatar":"https://www.baidu.com"}`
	user2 := User{}
	// 传递字节数组与指针类型 不返回对象 只返回是否发生错误
	err = json.Unmarshal([]byte(jsonStr), &user2)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	fmt.Printf("%#v\n", user2) // main.User{Id:20, Name:"李四", Avatar:"https://www.baidu.com"}
}
