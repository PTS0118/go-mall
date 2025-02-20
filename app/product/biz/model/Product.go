package model

import (
	"context"
	"github.com/PTS0118/go-mall/app/product/biz/dal/mysql"
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

//创建商品

//删除商品

//更新商品

// 查找商品
func GetProductById(ctx context.Context, id int32) (product *Product, err error) {
	err = mysql.DB.WithContext(ctx).Model(&Product{}).Where(&Product{Base: Base{Id: id}}).First(&product).Error
	return product, err
}

//查找商品列表
