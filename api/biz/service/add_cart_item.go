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

// @Summary 添加商品到购物车
// @Description 通过RPC调用添加商品到购物车
// @Tags Cart
// @Accept json
// @Produce json
// @Param req body cart.AddCartReq true "添加购物车请求"
// @Success 200 {object} cart.AddCartResp "成功响应"
// @Failure 400 {object} cart.AddCartResp "请求参数错误"
// @Router /cart/add [post]
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
