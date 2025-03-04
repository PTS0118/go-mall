package service

import (
	"context"
	"github.com/PTS0118/go-mall/app/order/biz/model"
	order "github.com/PTS0118/go-mall/rpc_gen/kitex_gen/order"
)

type UpdateOrderService struct {
	ctx context.Context
} // NewUpdateOrderService new UpdateOrderService
func NewUpdateOrderService(ctx context.Context) *UpdateOrderService {
	return &UpdateOrderService{ctx: ctx}
}

// Run create note info
func (s *UpdateOrderService) Run(req *order.UpdateOrderReq) (resp *order.UpdateOrderResp, err error) {
	//仅支持修改地址信息以及手机号等
	//获取主键ID
	orderData, err := model.GetOrder(s.ctx, req.GetOrderId())
	if err != nil {
		return &order.UpdateOrderResp{
			Code:    -1,
			Message: "修改订单失败，不存在该订单",
		}, err
	}

	err = model.UpdateOrder(s.ctx, &model.Order{
		Base:      model.Base{Id: orderData.Id},
		Telephone: req.GetTelephone(),
		AddressId: int(req.GetAddressId()),
	})
	if err != nil {
		return &order.UpdateOrderResp{
			Code:    -1,
			Message: "修改订单失败",
		}, err
	}
	return &order.UpdateOrderResp{
		Code:    0,
		Message: "修改订单成功",
	}, err
}
