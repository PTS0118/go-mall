package service

import (
	"context"
	"log"

	"github.com/PTS0118/go-mall/app/cart/biz/model"
	cart "github.com/PTS0118/go-mall/rpc_gen/kitex_gen/cart"
	"github.com/cloudwego/kitex/pkg/klog"
)

type GetCartService struct {
	ctx context.Context
} // NewGetCartService new GetCartService
func NewGetCartService(ctx context.Context) *GetCartService {
	return &GetCartService{ctx: ctx}
}

// Run create note info
func (s *GetCartService) Run(req *cart.GetCartReq) (resp *cart.GetCartResp, err error) {
	res, err := model.ListProductsByUserId(s.ctx, int32(req.UserId))
	if err != nil {
		items := make([]*cart.CartItem, len(res))
		log.Printf("len: %d", len(res))
		if res != nil {
			for i, v := range res {
				items[i] = &cart.CartItem{
					ProductId: uint32(v.ProductId),
					Count:     v.Count,
				}
			}
		}

		resp = &cart.GetCartResp{
			Code:    0,
			Message: "获取购物车成功",
			Items:   items,
			UserId:  req.UserId,
		}
		klog.Error("获取购物车成功：%v", err)
	} else {
		resp = &cart.GetCartResp{
			Code:    0,
			Message: "获取购物车失败",
			UserId:  req.UserId,
		}
	}

	return resp, err
}
