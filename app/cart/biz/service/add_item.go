package service

import (
	"context"

	"github.com/PTS0118/go-mall/app/cart/biz/model"
	cart "github.com/PTS0118/go-mall/rpc_gen/kitex_gen/cart"
	"github.com/cloudwego/kitex/pkg/klog"
)

type AddItemService struct {
	ctx context.Context
} // NewAddItemService new AddItemService
func NewAddItemService(ctx context.Context) *AddItemService {
	return &AddItemService{ctx: ctx}
}

// Run create note info
func (s *AddItemService) Run(req *cart.AddItemReq) (resp *cart.AddItemResp, err error) {
	param := &model.Cart{
		UserId:    req.UserId,
		ProductId: req.ProductId,
		Count:     req.Count,
	}
	id, err := model.CreateProduct(s.ctx, param)
	if err != nil {
		resp = &cart.AddItemResp{
			Code:    0,
			Message: "购物车添加商品成功",
		}
		klog.Error("购物车添加商品失败：%v", err)
	} else {
		resp = &cart.AddItemResp{
			Code:    0,
			Message: "购物车添加商品成功",
		}
	}
	return resp, err
}
