package main

import (
	"fmt"
	"reflect"

	"github.com/star574/go-base/basic/chapter_5/lession_2/discount"
)

func main() {

	// 通过不同实现得出不同的结果
	order := &discount.Order{
		Id:    101,
		Total: 100.00,
	}
	free := discount.FreeUser{
		Order: order,
	}
	vip := discount.VipUser{
		Order: order,
	}
	// 传递接口
	PrintOrderDiscount(&free)
	PrintOrderDiscount(&vip)

	// 断言
	var m = make(map[string]string)
	m["name"] = "张三"
	fmt.Printf("m[1]: %v\n", m["name"])
	fmt.Printf("reflect.TypeOf(m): %v\n", reflect.TypeOf(m))

	// 接口组合
	// 类似于向上转型
	// 注意接收者 调用者必须是值接收者 如果是指针接收者 则不能访问方法(加上&既可以使用指针接收者也可以使用值接收者)
	var dc DiscounterAndCoder = &CommonUser{
		Order: *order,
	}
	dc.ComputerOrder()
	dc.PrintCode()

}

// Discounter 接口 实现等级的用户的不同金额
type Discounter interface {
	ComputerOrder()
}
type Coder interface {
	PrintCode()
}

// 接口嵌套组合
type DiscounterAndCoder interface {
	Discounter
	Coder
}

// 实现两个接口
type CommonUser struct {
	Order discount.Order
}

func (c CommonUser) ComputerOrder() {
	fmt.Printf("订单账号: %d,订单金额: %.2f\n", c.Order.Id, c.Order.Total*0.7)
}
func (c *CommonUser) PrintCode() {
	fmt.Printf("订单账号: %d,订单金额: %.2f 打印二维码!\n", c.Order.Id, c.Order.Total*0.7)
}

// 打印订单 这里传递的是接口
func PrintOrderDiscount(d Discounter) {
	// 断言 ok 是否断言成功
	_, ok := d.(*discount.FreeUser)
	if ok {
		fmt.Println("ok")
	}
	d.ComputerOrder()
}
