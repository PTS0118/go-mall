package service

import (
	"context"

	order "github.com/PTS0118/go-mall/api/hertz_gen/api/order"
	"github.com/cloudwego/hertz/pkg/app"
)

type ListOrdersService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewListOrdersService(Context context.Context, RequestContext *app.RequestContext) *ListOrdersService {
	return &ListOrdersService{RequestContext: RequestContext, Context: Context}
}

func (h *ListOrdersService) Run(req *order.Empty) (resp *order.Empty, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
