package service

import (
	"context"
	"github.com/PTS0118/go-mall/app/order/biz/model"
	order "github.com/PTS0118/go-mall/rpc_gen/kitex_gen/order"
)

type MarkOrderPaidService struct {
	ctx context.Context
} // NewMarkOrderPaidService new MarkOrderPaidService
func NewMarkOrderPaidService(ctx context.Context) *MarkOrderPaidService {
	return &MarkOrderPaidService{ctx: ctx}
}

// Run create note info
func (s *MarkOrderPaidService) Run(req *order.MarkOrderPaidReq) (resp *order.MarkOrderPaidResp, err error) {
	// 修改订单状态
	err = model.MarkOrderStatus(int(req.Status), req.OrderId)
	if err != nil {
		return &order.MarkOrderPaidResp{
			Code:    -1,
			Message: "修改失败",
		}, err
	}

	return &order.MarkOrderPaidResp{
		Code:    0,
		Message: "修改成功",
	}, err
}
