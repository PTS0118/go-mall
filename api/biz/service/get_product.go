package service

import (
	"context"

	product "github.com/PTS0118/go-mall/api/hertz_gen/api/product"
	"github.com/PTS0118/go-mall/api/infra/rpc"
	rpcproduct "github.com/PTS0118/go-mall/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/hertz/pkg/app"
)

type GetProductService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetProductService(Context context.Context, RequestContext *app.RequestContext) *GetProductService {
	return &GetProductService{RequestContext: RequestContext, Context: Context}
}

// @Summary 获取商品详情
// @Description 通过RPC调用获取商品详情
// @Tags Product
// @Accept json
// @Produce json
// @Param req body product.ProductReq true "获取商品请求"
// @Success 200 {object} product.ProductResp "成功响应"
// @Failure 400 {object} product.ProductResp "请求参数错误"
// @Router /product/get [post]
func (h *GetProductService) Run(req *product.ProductReq) (resp *product.ProductResp, err error) {
	//判断参数是否为nil
	if req == nil {
		resp = &product.ProductResp{
			StatusCode: -1,
			StatusMsg:  "商品ID为空",
			Product:    &product.Product{},
		}
		return resp, nil
	}
	data, err := rpc.ProductClient.GetProduct(h.Context, &rpcproduct.GetProductReq{Id: req.GetId()})
	if err != nil {
		resp = &product.ProductResp{
			StatusCode: data.GetCode(),
			StatusMsg:  data.GetMessage(),
			Product:    &product.Product{},
		}
	} else {
		resp = &product.ProductResp{
			StatusCode: data.GetCode(),
			StatusMsg:  data.GetMessage(),
			Product: &product.Product{
				Id:          data.Product.Id,
				Name:        data.Product.Name,
				Description: data.Product.Description,
				Picture:     data.Product.Picture,
				Price:       data.Product.Price,
				Categories:  data.Product.Categories,
			},
		}
	}
	return resp, err
}
