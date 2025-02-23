package model

import (
	"context"
	"github.com/PTS0118/go-mall/app/order/biz/dal/mysql"
)

type OrderItem struct {
	Base
	ProductId  int     `json:"productId" column:"product_id"`
	OrderId    string  `json:"orderId" column:"order_id"`
	TotalPrice float64 `json:"totalPrice" column:"total_price"`
	UnitPrice  float64 `json:"unitPrice" column:"unit_price"`
	Count      int     `json:"count" column:"count"`
}

func (p OrderItem) TableName() string {
	return "order_item"
}

// 创建订单项
func CreateOrderItem(ctx context.Context, p *OrderItem) (id int32, err error) {
	result := mysql.DB.Create(&p)
	return p.Id, result.Error
}

// 删除订单项
func DeleteOrderItem(ctx context.Context, id int32) (err error) {
	result := mysql.DB.Where("id = ?", id).Delete(&OrderItem{Base: Base{Id: id}})
	return result.Error
}

// 更新订单
func UpdateOrderItem(ctx context.Context, p *OrderItem) (err error) {
	result := mysql.DB.Save(p)
	return result.Error
}

// 查找订单列表
func ListOrderItems(ctx context.Context, orderId string) (orderItems []*OrderItem, err error) {
	result := mysql.DB.Where(&OrderItem{Base: Base{IsDel: 0}, OrderId: orderId}).Find(&orderItems)
	return orderItems, result.Error
}
