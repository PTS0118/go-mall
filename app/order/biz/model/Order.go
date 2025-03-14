package model

import (
	"context"
	"github.com/PTS0118/go-mall/app/order/biz/dal/mysql"
)

type Order struct {
	Base
	UserId     int     `json:"userId" column:"user_id"`
	OrderId    string  `json:"orderId" column:"order_id"`
	TotalPrice float64 `json:"totalPrice" column:"total_price"`
	Status     int     `json:"status" column:"status"`
	Telephone  string  `json:"telephone" column:"telephone"`
	AddressId  int     `json:"addressId" column:"address_id"`
}

func (p Order) TableName() string {
	return "order"
}

// 创建订单
func CreateOrder(ctx context.Context, p *Order) (id int32, err error) {
	result := mysql.DB.Create(&p)
	return p.Id, result.Error
}

// 删除订单
func DeleteOrder(ctx context.Context, id int32) (err error) {
	result := mysql.DB.Where("id = ?", id).Delete(&Order{Base: Base{Id: id}})
	return result.Error
}

// 更新订单
func UpdateOrder(ctx context.Context, p *Order) (err error) {
	result := mysql.DB.Save(p)
	return result.Error
}

// 查找订单
func GetOrder(ctx context.Context, orderId string) (order *Order, err error) {
	result := mysql.DB.Where(&Order{Base: Base{IsDel: 0}, OrderId: orderId}).First(&order)
	return order, result.Error
}

// 查找订单列表
func ListOrders(ctx context.Context, userId int) (orders []*Order, err error) {
	result := mysql.DB.Where(&Order{Base: Base{IsDel: 0}, UserId: userId}).Find(&orders)
	return orders, result.Error
}

// 设置订单状态
func MarkOrderStatus(status int, orderId string) (err error) {
	result := mysql.DB.Model(&Order{}).Where("order_id = ?", orderId).Update("status", status)
	return result.Error
}
