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

// @Summary 获取购物车内容
// @Description 通过RPC调用获取购物车内容
// @Tags Cart
// @Accept json
// @Produce json
// @Param req body cart.Empty true "获取购物车请求"
// @Success 200 {object} cart.GetCartResp "成功响应"
// @Failure 400 {object} cart.GetCartResp "请求参数错误"
// @Router /cart/get [post]
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
	items := make([]*cart.CartItem, len(data.Items))
	// 检查 items 是否为空
	if data.Size() == 0 {
		log.Printf("items length: %d", data.Size())
		return resp, nil // 或者返回一个适当的错误信息
	} else {
		log.Printf("items length 1: %d", data.Size())
		for key, value := range data.Items {
			log.Printf("key: %d, value: %+v", key, value)
		}
	}
	for i := 0; i < len(data.GetItems()); i++ {

		product, err := rpc.ProductClient.GetProduct(h.Context, &rpcproduct.GetProductReq{
			Id: int32(data.GetItems()[i].GetProductId()),
		})
		log.Printf("product %+v", product)
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
