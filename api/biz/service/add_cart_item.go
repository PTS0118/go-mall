package service

import (
	"context"

	cart "github.com/PTS0118/go-mall/api/hertz_gen/api/cart"
	"github.com/PTS0118/go-mall/api/infra/rpc"
	rpccart "github.com/PTS0118/go-mall/rpc_gen/kitex_gen/cart"
	"github.com/cloudwego/hertz/pkg/app"
)

type AddCartItemService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewAddCartItemService(Context context.Context, RequestContext *app.RequestContext) *AddCartItemService {
	return &AddCartItemService{RequestContext: RequestContext, Context: Context}
}

func (h *AddCartItemService) Run(req *cart.AddCartReq) (resp *cart.AddCartResp, err error) {
	//判断参数是否为nil
	if req == nil {
		resp = &cart.AddCartResp{
			StatusCode: -1,
			StatusMsg:  "req为空",
		}
		return resp, nil
	}
	data, err := rpc.CartClient.AddItem(h.Context, &rpccart.AddItemReq{
		UserId:    req.UserId,
		ProductId: req.ProductId,
		Count:     req.Count,
	})
	if err != nil {
		resp = &cart.AddCartResp{
			StatusCode: data.GetCode(),
			StatusMsg:  data.GetMessage(),
		}
	} else {
		resp = &cart.AddCartResp{
			StatusCode: data.GetCode(),
			StatusMsg:  data.GetMessage(),
		}
	}
	return resp, err
}
