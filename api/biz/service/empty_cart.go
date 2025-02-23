package service

import (
	"context"

	cart "github.com/PTS0118/go-mall/api/hertz_gen/api/cart"
	"github.com/PTS0118/go-mall/api/infra/rpc"
	rpccart "github.com/PTS0118/go-mall/rpc_gen/kitex_gen/cart"
	"github.com/cloudwego/hertz/pkg/app"
)

type EmptyCartService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewEmptyCartService(Context context.Context, RequestContext *app.RequestContext) *EmptyCartService {
	return &EmptyCartService{RequestContext: RequestContext, Context: Context}
}

func (h *EmptyCartService) Run(req *cart.Empty) (resp *cart.EmptyCartResp, err error) {
	//判断参数是否为nil
	if req == nil {
		resp = &cart.EmptyCartResp{
			StatusCode: -1,
			StatusMsg:  "req为空",
		}
		return resp, nil
	}
	data, err := rpc.CartClient.EmptyCart(h.Context, &rpccart.EmptyCartReq{
		UserId: req.UserId,
	})
	if err != nil {
		resp = &cart.EmptyCartResp{
			StatusCode: data.GetCode(),
			StatusMsg:  data.GetMessage(),
		}
	} else {
		resp = &cart.EmptyCartResp{
			StatusCode: data.GetCode(),
			StatusMsg:  data.GetMessage(),
		}
	}
	return resp, err
}
