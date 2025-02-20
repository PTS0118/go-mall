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

func (h *GetProductService) Run(req *product.ProductReq) (resp *product.ProductResp, err error) {
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
