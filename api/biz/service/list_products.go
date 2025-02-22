package service

import (
	"context"
	product "github.com/PTS0118/go-mall/api/hertz_gen/api/product"
	"github.com/PTS0118/go-mall/api/infra/rpc"
	rpcproduct "github.com/PTS0118/go-mall/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/hertz/pkg/app"
)

type ListProductsService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewListProductsService(Context context.Context, RequestContext *app.RequestContext) *ListProductsService {
	return &ListProductsService{RequestContext: RequestContext, Context: Context}
}

func (h *ListProductsService) Run(req *product.ListProductsReq) (resp *product.ListProductsResp, err error) {
	//判断参数是否为nil
	if req == nil {
		resp = &product.ListProductsResp{
			StatusCode: -1,
			StatusMsg:  "商品ID为空",
			Product:    make([]*product.Product, 0),
		}
		return resp, nil
	}
	data, err := rpc.ProductClient.ListProducts(h.Context, &rpcproduct.ListProductsReq{PageSize: req.PageSize, Page: req.Page})
	if err != nil {
		resp = &product.ListProductsResp{
			StatusCode: data.GetCode(),
			StatusMsg:  data.GetMessage(),
			Product:    make([]*product.Product, 0),
		}
	} else {
		list := make([]*product.Product, req.PageSize)
		for key, value := range data.Products {
			list[key] = &product.Product{
				Id:          value.Id,
				Name:        value.Name,
				Description: value.Description,
				Picture:     value.Picture,
				Price:       value.Price,
				Categories:  value.Categories,
			}
		}
		resp = &product.ListProductsResp{
			StatusCode: data.GetCode(),
			StatusMsg:  data.GetMessage(),
			Product:    list,
		}
	}
	return resp, err
}
