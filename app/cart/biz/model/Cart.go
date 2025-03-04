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

// func (p *Cart) BeforeDelete(tx *gorm.DB) (err error) {
// 	if p.Base.IsDel == 0 {
// 		p.IsDel = 1
// 		return tx.Save(p).Error // 更新 is_del 字段而不实际删除记录
// 	}
// 	return nil
// }

// 创建商品
func AddOrUpdateCart(ctx context.Context, p *Cart) (err error) {
	tx := mysql.DB.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	// 查询商品是否存在
	var existingCart Cart
	if err := tx.Where("user_id = ? AND product_id = ?", p.UserId, p.ProductId).First(&existingCart).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			// 商品不存在，插入新记录
			if err := tx.Create(p).Error; err != nil {
				tx.Rollback()
				return err
			}
		} else {
			tx.Rollback()
			return err
		}
	} else {
		// 商品存在，更新记录
		existingCart.Count += p.Count
		if err := tx.Save(&existingCart).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	// 提交事务
	return tx.Commit().Error
}

// 删除商品
func DeleteCart(ctx context.Context, id int32) (err error) {
	result := mysql.DB.Where("user_id = ?", id).Delete(&Cart{})
	return result.Error
}

// 更新商品
func UpdateCart(ctx context.Context, p *Cart) (err error) {
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
