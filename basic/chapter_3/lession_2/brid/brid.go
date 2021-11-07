package brid

import "fmt"

type (
	Brid1 struct {
		Name string
	}
	Brid2 struct {
		Name string
	}
)

// brid 方法调用者 接收者
// brid1 值传递  不会改变原始数据
// *brid1 指针传递 对原有数据操作会改变原始数据
func (brid *Brid1) Fly() {
	brid.Name = "鹰"
	fmt.Println(brid.Name, "会飞")
}

func (brid Brid2) Run() {
	fmt.Println(brid.Name, "会跑")
}
