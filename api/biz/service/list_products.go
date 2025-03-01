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

// @Summary 获取商品列表
// @Description 通过RPC调用获取商品列表
// @Tags Product
// @Accept json
// @Produce json
// @Param req body product.ListProductsReq true "获取商品列表请求"
// @Success 200 {object} product.ListProductsResp "成功响应"
// @Failure 400 {object} product.ListProductsResp "请求参数错误"
// @Router /product/list [post]
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
