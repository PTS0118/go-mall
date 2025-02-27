package enum

type OrderStatus int

const (
	Unpaid    OrderStatus = 0 //待支付
	Paid      OrderStatus = 1 //已支付
	Cancelled OrderStatus = 2 //已取消
	Completed OrderStatus = 3
)
