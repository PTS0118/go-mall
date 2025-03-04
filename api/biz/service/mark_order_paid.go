package service

import (
	"context"
	"github.com/PTS0118/go-mall/api/infra/rpc"
	rpcOrder "github.com/PTS0118/go-mall/rpc_gen/kitex_gen/order"
	"strconv"

	order "github.com/PTS0118/go-mall/api/hertz_gen/api/order"
	"github.com/cloudwego/hertz/pkg/app"
)

type MarkOrderPaidService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewMarkOrderPaidService(Context context.Context, RequestContext *app.RequestContext) *MarkOrderPaidService {
	return &MarkOrderPaidService{RequestContext: RequestContext, Context: Context}
}

func (h *MarkOrderPaidService) Run(req *order.MarkOrderPaidReq) (resp *order.MarkOrderPaidResp, err error) {
	status, _ := strconv.Atoi(req.Status)
	data, err := rpc.OrderClient.MarkOrderPaid(h.Context, &rpcOrder.MarkOrderPaidReq{
		OrderId: req.OrderId,
		UserId:  req.UserId,
		Status:  int32(status),
	})
	if err != nil {
		return &order.MarkOrderPaidResp{
			StatusCode: data.Code,
			StatusMsg:  data.Message,
		}, err
	}
	return &order.MarkOrderPaidResp{
		StatusCode: data.Code,
		StatusMsg:  data.Message,
	}, nil
}
