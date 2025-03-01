package service

import (
	"context"
	"log"
	"strconv"

	cart "github.com/PTS0118/go-mall/api/hertz_gen/api/cart"
	"github.com/PTS0118/go-mall/api/infra/rpc"
	rpccart "github.com/PTS0118/go-mall/rpc_gen/kitex_gen/cart"
	rpcproduct "github.com/PTS0118/go-mall/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/hertz/pkg/app"
)

type GetCartService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetCartService(Context context.Context, RequestContext *app.RequestContext) *GetCartService {
	return &GetCartService{RequestContext: RequestContext, Context: Context}
}

func (h *GetCartService) Run(req *cart.Empty) (resp *cart.GetCartResp, err error) {
	//判断参数是否为nil
	if req == nil {
		resp = &cart.GetCartResp{
			StatusCode: -1,
			StatusMsg:  "req为空",
		}
		return resp, nil
	}
	data, err := rpc.CartClient.GetCart(h.Context, &rpccart.GetCartReq{
		UserId: req.UserId,
	})
	items := make([]*cart.CartItem, data.Size())
	// 检查 items 是否为空
	if data.Size() == 0 {
		log.Printf("items length: %d", data.Size())
		return resp, nil // 或者返回一个适当的错误信息
	} else {
		log.Printf("items length 1: %d", data.Size())
	}
	for i := 0; i < data.Size(); i++ {
		log.Printf("test %d", data.GetItems()[i].GetProductId())
		product, err := rpc.ProductClient.GetProduct(h.Context, &rpcproduct.GetProductReq{
			Id: int32(data.GetItems()[i].GetProductId()),
		})
		if err == nil {
			items[i] = &cart.CartItem{
				ProductId:   data.GetItems()[i].GetProductId(),
				ProductName: product.Product.GetName(),
				Count:       data.GetItems()[i].GetCount(),
				Description: product.Product.GetDescription(),
				TotalPrice:  data.GetItems()[i].GetTotalPrice(),
			}
			items[i].TotalPrice = strconv.FormatFloat(float64(product.Product.GetPrice())*float64(data.GetItems()[i].GetCount()), 'f', 2, 64)
		} else {
			log.Printf("err: %s", err)
			resp = &cart.GetCartResp{
				StatusCode: data.GetCode(),
				StatusMsg:  "拼装product错误",
				Items:      items,
			}
		}

	}
	if err != nil {
		resp = &cart.GetCartResp{
			StatusCode: data.GetCode(),
			StatusMsg:  data.GetMessage(),
			Items:      items,
		}
	} else {
		resp = &cart.GetCartResp{
			StatusCode: data.GetCode(),
			StatusMsg:  data.GetMessage(),
			Items:      items,
		}
	}
	return resp, err
}
