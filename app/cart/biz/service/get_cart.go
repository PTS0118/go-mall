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
	if req.UserId == 0 {
		resp = &cart.GetCartResp{
			Code:    -1,
			Message: "无效的用户ID",
			UserId:  req.UserId,
		}
		return resp, nil
	}
	res, err := model.ListProductsByUserId(s.ctx, int32(req.UserId))
	if err != nil {
		// 查询失败，返回失败响应
		resp = &cart.GetCartResp{
			Code:    -1,
			Message: "获取购物车失败",
			UserId:  req.UserId,
		}
		klog.Error("获取购物车失败：%v", err)
	} else {
		// 查询成功，处理结果并返回成功响应
		items := make([]*cart.CartItem, len(res))
		log.Printf("len: %d", len(res))
		for i, v := range res {
			items[i] = &cart.CartItem{
				ProductId: uint32(v.ProductId),
				Count:     v.Count,
			}
		}

		resp = &cart.GetCartResp{
			Code:    0,
			Message: "获取购物车成功",
			Items:   items,
			UserId:  req.UserId,
		}
	}

	return resp, err
}
