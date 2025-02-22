package service

import (
	"context"
	"github.com/PTS0118/go-mall/api/infra/rpc"
	rpcproduct "github.com/PTS0118/go-mall/rpc_gen/kitex_gen/product"

	product "github.com/PTS0118/go-mall/api/hertz_gen/api/product"
	"github.com/cloudwego/hertz/pkg/app"
)

type UpdateProductService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewUpdateProductService(Context context.Context, RequestContext *app.RequestContext) *UpdateProductService {
	return &UpdateProductService{RequestContext: RequestContext, Context: Context}
}

func (h *UpdateProductService) Run(req *product.UpdateProductReq) (resp *product.UpdateProductResp, err error) {
	//判断参数是否为nil
	if req == nil {
		resp = &product.UpdateProductResp{
			StatusCode: -1,
			StatusMsg:  "参数为空",
		}
		return resp, nil
	}
	productData := &rpcproduct.Product{
		Name:        req.Product.Name,
		Description: req.Product.Description,
		Picture:     req.Product.Picture,
		Price:       req.Product.Price,
		Categories:  req.Product.Categories,
	}
	data, err := rpc.ProductClient.UpdateProduct(h.Context, &rpcproduct.UpdateProductReq{Product: productData})
	if err != nil {
		resp = &product.UpdateProductResp{
			StatusCode: data.GetCode(),
			StatusMsg:  data.GetMessage(),
		}
	} else {
		resp = &product.UpdateProductResp{
			StatusCode: data.GetCode(),
			StatusMsg:  data.GetMessage(),
		}
	}
	return resp, err
}
