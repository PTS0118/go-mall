package service

import (
	"context"
	"github.com/PTS0118/go-mall/app/product/biz/model"
	product "github.com/PTS0118/go-mall/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/kitex/pkg/klog"
	"strings"
)

type GetProductService struct {
	ctx context.Context
} // NewGetProductService new GetProductService
func NewGetProductService(ctx context.Context) *GetProductService {
	return &GetProductService{ctx: ctx}
}

// Run create note info
func (s *GetProductService) Run(req *product.GetProductReq) (resp *product.GetProductResp, err error) {
	data, err := model.GetProductById(s.ctx, req.GetId())
	if err != nil {
		resp = &product.GetProductResp{
			Code:    -1,
			Message: "获取商品失败",
			Product: &product.Product{},
		}
		klog.Error("商品查询失败：%v", err)
	} else if data == nil {
		resp = &product.GetProductResp{
			Code:    -1,
			Message: "商品不存在",
			Product: &product.Product{},
		}
		klog.Info("商品查询失败：商品不存在")
	} else {
		categoryList := strings.Split(data.Categories, ",")
		resp = &product.GetProductResp{
			Code:    0,
			Message: "获取商品成功",
			Product: &product.Product{
				Id:          data.Id,
				Name:        data.Name,
				Description: data.Description,
				Picture:     data.Picture,
				Price:       data.Price,
				Categories:  categoryList,
			},
		}
		klog.Info("商品查询成功")
	}
	return resp, err
}
