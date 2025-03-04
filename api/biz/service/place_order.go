package service

import (
	"context"

	order "github.com/PTS0118/go-mall/api/hertz_gen/api/order"
	"github.com/PTS0118/go-mall/api/infra/rpc"
	rpcorder "github.com/PTS0118/go-mall/rpc_gen/kitex_gen/order"
	rpcproduct "github.com/PTS0118/go-mall/rpc_gen/kitex_gen/product"
	rpcuser "github.com/PTS0118/go-mall/rpc_gen/kitex_gen/user"
	"github.com/cloudwego/hertz/pkg/app"
)

type PlaceOrderService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewPlaceOrderService(Context context.Context, RequestContext *app.RequestContext) *PlaceOrderService {
	return &PlaceOrderService{RequestContext: RequestContext, Context: Context}
}

func (h *PlaceOrderService) Run(req *order.PlaceOrderReq) (resp *order.PlaceOrderResp, err error) {
	//校验参数
	if req == nil {
		resp = &order.PlaceOrderResp{
			StatusCode: -1,
			StatusMsg:  "参数为空",
			OrderId:    "",
		}
	}
	//根据user_id查询用户信息
	userData, err := rpc.UserClient.GetUser(h.Context, &rpcuser.GetUserReq{
		UserId:   req.UserId,
		Username: "",
		Email:    "",
	})
	if err != nil {
		resp = &order.PlaceOrderResp{
			StatusCode: -1,
			StatusMsg:  "下单失败（不存在该用户）",
			OrderId:    "",
		}
		return resp, err
	}
	//判断是否还有库存
	for _, value := range req.OrderItems {
		productData, _ := rpc.ProductClient.GetProduct(h.Context, &rpcproduct.GetProductReq{Id: value.ProductId})
		if productData.Product.Stock < value.Count {
			resp = &order.PlaceOrderResp{
				StatusCode: -1,
				StatusMsg:  "下单失败（商品：" + productData.Product.Name + " 库存不足）",
				OrderId:    "",
			}
			return resp, err
		}
	}

	//构建订单项对象
	orderItems := make([]*rpcorder.OrderItem, len(req.OrderItems))
	for key, value := range req.OrderItems {
		orderItems[key] = &rpcorder.OrderItem{
			ProductId:  value.ProductId,
			UnitPrice:  value.UnitPrice,
			TotalPrice: value.TotalPrice,
			Count:      value.Count,
		}
	}

	orderData := &rpcorder.PlaceOrderReq{
		UserId:     uint32(req.UserId),
		AddressId:  req.AddressId,
		Email:      userData.Email,
		Telephone:  userData.Telephone,
		OrderItems: orderItems,
	}
	placeOrder, err := rpc.OrderClient.PlaceOrder(h.Context, orderData)
	resp = &order.PlaceOrderResp{
		StatusCode: placeOrder.Code,
		StatusMsg:  placeOrder.Message,
		OrderId:    placeOrder.OrderId,
	}
	return resp, err
}
