package discount

import "fmt"

// 免费会员
type FreeUser struct {
	Order *Order
}

// 实现接口
func (f *FreeUser) ComputerOrder() {
	fmt.Printf("订单账号: %d,订单金额: %.2f\n", f.Order.Id, f.Order.Total)
}
