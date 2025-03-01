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

// @Summary 清空购物车
// @Description 通过RPC调用清空购物车
// @Tags Cart
// @Accept json
// @Produce json
// @Param req body cart.Empty true "清空购物车请求"
// @Success 200 {object} cart.EmptyCartResp "成功响应"
// @Failure 400 {object} cart.EmptyCartResp "请求参数错误"
// @Router /cart/empty [post]
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
	if err == nil {
		resp = &cart.EmptyCartResp{
			StatusCode: data.GetCode(),
			StatusMsg:  "清空购物车成功",
		}
	} else {
		resp = &cart.EmptyCartResp{
			StatusCode: data.GetCode(),
			StatusMsg:  "清空购物车失败",
		}
	}
	return resp, err
}
