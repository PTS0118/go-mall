package model

import (
	"context"
	"github.com/PTS0118/go-mall/app/order/biz/dal/mysql"
)

type OrderItem struct {
	Base
	UserId     int     `json:"userId" column:"user_id"`
	OrderId    int     `json:"orderId" column:"order_id"`
	TotalPrice float64 `json:"totalPrice" column:"total_price"`
	Status     int     `json:"status" column:"status"`
	Telephone  float32 `json:"telephone" column:"telephone"`
	AddressId  int     `json:"addressId" column:"address_id"`
}

func (p Order) TableName() string {
	return "order_item"
}

// 创建订单
func CreateOrderItem(ctx context.Context, p *Order) (id int32, err error) {
	result := mysql.DB.Create(&p)
	return p.Id, result.Error
}

// 删除订单
func DeleteOrderItem(ctx context.Context, id int32) (err error) {
	result := mysql.DB.Where("id = ?", id).Delete(&Order{Base: Base{Id: id}})
	return result.Error
}

// 更新订单
func UpdateOrderItem(ctx context.Context, p *Order) (err error) {
	result := mysql.DB.Save(p)
	return result.Error
}

// 查找订单列表
func ListOrderItems(ctx context.Context, userId int) (products []*Order, err error) {
	result := mysql.DB.Where(&Order{Base: Base{IsDel: 0}, UserId: userId}).Find(&Order{})
	return products, result.Error
}
