package main

import "fmt"

type Priter interface {
	PritOrder() // 方法的签名或者原型
}

// 实现了Priter接口的方法
type Order struct {
	Id    int
	Total float64
}

// 实现接口
func (o *Order) PritOrder() {
	fmt.Printf("o: %+v\n", o)
}

type Qr interface {
	Priter // 方法的签名或者原型
	PritQr()
}
type OrderQr struct {
	Order
}

func (o *Order) PritQr() {
	fmt.Printf("o: %d\n", o.Id)
}
func (o *OrderQr) PritQr() {
	fmt.Printf("打印二维码 o: %d\n", o.Id)
}

// 空接口
type Empty interface{}

// 空接口
func main() {
	// 间接访问Order的方法
	var p Priter = &Order{
		Id:    999,
		Total: 10.00,
	}
	// 必须实现所有方法
	var q Qr = &Order{
		Id:    100,
		Total: 10.00,
	}
	p.PritOrder()
	q.PritOrder()
	q.PritQr()
	q = &OrderQr{
		Order: Order{
			Id:    666,
			Total: 89129,
		},
	}
	// 虽然OrderQr没有明确实现PritOrder  但是可以间接调用PritOrder  匿名结构体 类似于继承
	q.PritOrder() // o: &{Id:666 Total:89129}
	q.PritQr()    // 打印二维码 o: 666

	// 要实现接口必须实现所有方法 空接口中没有任何方法 所以所有的类型都已经默认实现了空接口 类似于继承
	// 空接口依然是接口类型 但是不能相互转换
	var empty Empty = 456
	empty = false
	empty = 100.00
	empty = 10 * 2i
	empty = func() string {
		s := "哈哈哈"
		return s
	}()
	// 使用类型断言赋值 相互转换
	var s string = empty.(string)
	fmt.Printf("empty: %v\n", empty)
	fmt.Printf("empty: %v\n", s)

	// 空接口应用场景
	map1 := map[string]interface{}{
		"id":   1,
		"name": "张三",
	}
	fmt.Printf("map1: %+#v\n", map1) // map1: map[string]interface {}{"id":1, "name":"张三"}

}
