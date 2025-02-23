package service

import (
	"context"

	"github.com/PTS0118/go-mall/app/cart/biz/model"
	cart "github.com/PTS0118/go-mall/rpc_gen/kitex_gen/cart"
	"github.com/cloudwego/kitex/pkg/klog"
)

type EmptyCartService struct {
	ctx context.Context
} // NewEmptyCartService new EmptyCartService
func NewEmptyCartService(ctx context.Context) *EmptyCartService {
	return &EmptyCartService{ctx: ctx}
}

// Run create note info
func (s *EmptyCartService) Run(req *cart.EmptyCartReq) (resp *cart.EmptyCartResp, err error) {
	res, err := model.ListProductsByUserId(s.ctx, int32(req.UserId))
	for _, v := range res {
		err = model.DeleteProduct(s.ctx, int32(v.ProductId))
	}
	if err != nil {
		resp = &cart.EmptyCartResp{
			Code:    0,
			Message: "清空购物车失败",
		}
		klog.Error("清空购物车失败：%v", err)
	} else {
		resp = &cart.EmptyCartResp{
			Code:    0,
			Message: "清空购物车成功",
		}
	}
}
