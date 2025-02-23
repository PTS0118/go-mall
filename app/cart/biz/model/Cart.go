package model

import (
	"context"

	"github.com/PTS0118/go-mall/app/cart/biz/dal/mysql"
	"gorm.io/gorm"
)

type Cart struct {
	Base
	UserId    uint32 `json:"description" column:"description"`
	ProductId uint32 `json:"picture" column:"picture"`
	Count     int32  `json:"price" column:"price"`
}

func (p Cart) TableName() string {
	return "cart"
}

func (p *Cart) BeforeDelete(tx *gorm.DB) (err error) {
	if p.Base.IsDel == 0 {
		p.IsDel = 1
		return tx.Save(p).Error // 更新 is_del 字段而不实际删除记录
	}
	return nil
}

// 创建商品
func CreateProduct(ctx context.Context, p *Cart) (id int32, err error) {
	result := mysql.DB.Create(&p)
	return p.Id, result.Error
}

// 删除商品
func DeleteProduct(ctx context.Context, id int32) (err error) {
	result := mysql.DB.Where("id = ?", id).Delete(&Cart{})
	return result.Error
}

// 更新商品
func UpdateProduct(ctx context.Context, p *Cart) (err error) {
	result := mysql.DB.Save(&Cart{})
	return result.Error
}

// 查找商品
func GetCartByProductId(ctx context.Context, id uint32) (cart *Cart, err error) {
	err = mysql.DB.WithContext(ctx).Model(&Cart{}).Where(&Cart{ProductId: id}).First(&cart).Error
	return cart, err
}

// 查找商品列表
func ListProductsByUserId(ctx context.Context, user_id int32) (carts []*Cart, err error) {
	result := mysql.DB.Where(&Cart{Base: Base{IsDel: 0}, UserId: uint32(user_id)}).Find(&carts)
	return carts, result.Error
}
