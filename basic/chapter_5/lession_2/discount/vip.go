package discount

import "fmt"

// vip会员
type VipUser struct {
	Order *Order
}

// 实现接口
func (f *VipUser) ComputerOrder() {
	fmt.Printf("订单账号: %d,订单金额: %.2f", f.Order.Id, f.Order.Total*0.8)
}
