package service

import (
	"context"
	"github.com/PTS0118/go-mall/app/product/biz/model"
	product "github.com/PTS0118/go-mall/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/kitex/pkg/klog"
	"strings"
)

type CreateProductService struct {
	ctx context.Context
} // NewCreateProductService new CreateProductService
func NewCreateProductService(ctx context.Context) *CreateProductService {
	return &CreateProductService{ctx: ctx}
}

// Run create note info
func (s *CreateProductService) Run(req *product.CreateProductReq) (resp *product.CreateProductResp, err error) {
	param := &model.Product{
		Name:        req.Product.Name,
		Description: req.Product.Description,
		Picture:     req.Product.Picture,
		Price:       req.Product.Price,
		Categories:  strings.Join(req.Product.Categories, ","),
	}
	id, err := model.CreateProduct(s.ctx, param)
	if err != nil {
		resp = &product.CreateProductResp{
			Code:      -1,
			Message:   "创建商品失败",
			ProductId: 0,
		}
		klog.Error("创建商品失败：%v", err)
	} else {
		resp = &product.CreateProductResp{
			Code:      0,
			Message:   "创建商品成功",
			ProductId: id,
		}
	}
	return resp, err
}
