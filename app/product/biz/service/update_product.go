package service

import (
	"context"
	"github.com/PTS0118/go-mall/app/product/biz/model"
	product "github.com/PTS0118/go-mall/rpc_gen/kitex_gen/product"
	"strings"
	"time"
)

type UpdateProductService struct {
	ctx context.Context
} // NewUpdateProductService new UpdateProductService
func NewUpdateProductService(ctx context.Context) *UpdateProductService {
	return &UpdateProductService{ctx: ctx}
}

// Run create note info
func (s *UpdateProductService) Run(req *product.UpdateProductReq) (resp *product.UpdateProductResp, err error) {
	existingProduct, err := model.GetProductById(s.ctx, req.Product.Id)
	if err != nil {
		resp = &product.UpdateProductResp{
			Code:    -1,
			Message: "商品不存在",
		}
		return resp, err
	}
	// 更新现有产品记录的基础信息（Base）和其他字段
	existingProduct.Base.Id = req.Product.Id
	existingProduct.Name = req.Product.Name
	existingProduct.Description = req.Product.Description
	existingProduct.Picture = req.Product.Picture
	existingProduct.Price = req.Product.Price
	existingProduct.Categories = strings.Join(req.Product.Categories, ",")
	existingProduct.UpdatedAt = time.Now()
	err = model.UpdateProduct(s.ctx, existingProduct)
	if err != nil {
		resp = &product.UpdateProductResp{
			Code:    -1,
			Message: "更新商品失败",
		}

	} else {
		resp = &product.UpdateProductResp{
			Code:    0,
			Message: "更新商品成功",
		}
	}
	return resp, err
}
