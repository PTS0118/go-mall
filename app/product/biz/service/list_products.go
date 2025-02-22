package service

import (
	"context"
	"fmt"
	"github.com/PTS0118/go-mall/app/product/biz/model"
	product "github.com/PTS0118/go-mall/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/kitex/pkg/klog"
	"strings"
)

type ListProductsService struct {
	ctx context.Context
} // NewListProductsService new ListProductsService
func NewListProductsService(ctx context.Context) *ListProductsService {
	return &ListProductsService{ctx: ctx}
}

// Run create note info
func (s *ListProductsService) Run(req *product.ListProductsReq) (resp *product.ListProductsResp, err error) {
	data, err := model.ListProducts(s.ctx, req.Page, req.PageSize)
	if err != nil {
		resp = &product.ListProductsResp{
			Code:     -1,
			Message:  "获取商品失败",
			Products: make([]*product.Product, 0),
		}
		klog.Error("商品查询失败：%v", err)
	} else if data == nil {
		resp = &product.ListProductsResp{
			Code:     -1,
			Message:  "商品不存在",
			Products: make([]*product.Product, 0),
		}
		klog.Info("商品查询失败：商品不存在")
	} else {
		list := make([]*product.Product, req.PageSize)
		for key, value := range data {
			fmt.Println("key: %v, value: %+v", key, value)
			categoryList := strings.Split(value.Categories, ",")
			list[key] = &product.Product{
				Id:          value.Id,
				Name:        value.Name,
				Description: value.Description,
				Picture:     value.Picture,
				Price:       value.Price,
				Categories:  categoryList,
			}
		}
		resp = &product.ListProductsResp{
			Code:     0,
			Message:  "获取商品成功",
			Products: list,
		}
	}
	return resp, err
}
