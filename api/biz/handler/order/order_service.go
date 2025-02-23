package order

import (
	"context"

	"github.com/PTS0118/go-mall/api/biz/service"
	"github.com/PTS0118/go-mall/api/biz/utils"
	order "github.com/PTS0118/go-mall/api/hertz_gen/api/order"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// ListOrders .
// @router /list [GET]
func ListOrders(ctx context.Context, c *app.RequestContext) {
	var err error
	var req order.Empty
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &order.Empty{}
	resp, err = service.NewListOrdersService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// PlaceOrder .
// @router /place [POST]
func PlaceOrder(ctx context.Context, c *app.RequestContext) {
	var err error
	var req order.PlaceOrderReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &order.PlaceOrderResp{}
	//设置用户ID
	req.UserId = int32(utils.GetUserId(c))
	resp, err = service.NewPlaceOrderService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// MarkOrderPaid .
// @router /markPaid [POST]
func MarkOrderPaid(ctx context.Context, c *app.RequestContext) {
	var err error
	var req order.MarkOrderPaidReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &order.MarkOrderPaidResp{}
	resp, err = service.NewMarkOrderPaidService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// UpdateOrder .
// @router /update [POST]
func UpdateOrder(ctx context.Context, c *app.RequestContext) {
	var err error
	var req order.UpdateOrderReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &order.UpdateOrderResp{}
	resp, err = service.NewUpdateOrderService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}
