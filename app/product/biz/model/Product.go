package model

import (
	"context"
	"github.com/PTS0118/go-mall/app/product/biz/dal/mysql"
	"gorm.io/gorm"
)

type Product struct {
	Base
	Name        string  `json:"name" column:"name"`
	Description string  `json:"description" column:"description"`
	Picture     string  `json:"picture" column:"picture"`
	Price       float32 `json:"price" column:"price"`
	Categories  string  `json:"categories" column:"categories"` // 使用切片定义 categories
}

func (p Product) TableName() string {
	return "product"
}

func (p *Product) BeforeDelete(tx *gorm.DB) (err error) {
	if p.Base.IsDel == 0 {
		p.IsDel = 1
		return tx.Save(p).Error // 更新 is_del 字段而不实际删除记录
	}
	return nil
}

// 创建商品
func CreateProduct(ctx context.Context, p *Product) (id int32, err error) {
	result := mysql.DB.Create(&p)
	return p.Id, result.Error
}

// 删除商品
func DeleteProduct(ctx context.Context, id int32) (err error) {
	result := mysql.DB.Where("id = ?", id).Delete(&Product{})
	return result.Error
}

// 更新商品
func UpdateProduct(ctx context.Context, p *Product) (err error) {
	result := mysql.DB.Save(&Product{})
	return result.Error
}

// 查找商品
func GetProductById(ctx context.Context, id int32) (product *Product, err error) {
	err = mysql.DB.WithContext(ctx).Model(&Product{}).Where(&Product{Base: Base{Id: id, IsDel: 0}}).First(&product).Error
	return product, err
}

// 查找商品列表
func ListProducts(ctx context.Context, page int32, pageSize int64) (products []*Product, err error) {
	offset := (int(page) - 1) * (int(pageSize))
	result := mysql.DB.Where(&Product{Base: Base{IsDel: 0}}).Limit(int(pageSize)).Offset(offset).Find(&products)
	return products, result.Error
}
