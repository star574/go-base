package test

import "fmt"

// User init()  包初始化
type User struct {
	Id int
}

// UserInfo 向外暴露 UserInfo (一个初始化好的User)
var UserInfo *User

func init() {
	fmt.Println("包初始化 UserInfo init")
	UserInfo = &User{
		Id: 1,
	}
}
