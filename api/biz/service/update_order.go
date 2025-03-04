package service

import (
	"context"
	order "github.com/PTS0118/go-mall/api/hertz_gen/api/order"
	"github.com/PTS0118/go-mall/api/infra/rpc"
	rpcOrder "github.com/PTS0118/go-mall/rpc_gen/kitex_gen/order"
	"github.com/cloudwego/hertz/pkg/app"
)

type UpdateOrderService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewUpdateOrderService(Context context.Context, RequestContext *app.RequestContext) *UpdateOrderService {
	return &UpdateOrderService{RequestContext: RequestContext, Context: Context}
}

func (h *UpdateOrderService) Run(req *order.UpdateOrderReq) (resp *order.UpdateOrderResp, err error) {
	data, err := rpc.OrderClient.UpdateOrder(h.Context, &rpcOrder.UpdateOrderReq{
		OrderId:   req.OrderId,
		AddressId: req.AddressId,
		Email:     req.Email,
		Telephone: req.Telephone,
		UserId:    req.UserId,
	})
	if err != nil {
		return &order.UpdateOrderResp{
			StatusCode: data.Code,
			StatusMsg:  data.Message,
		}, err
	}
	return &order.UpdateOrderResp{
		StatusCode: data.Code,
		StatusMsg:  data.Message,
	}, nil
}
