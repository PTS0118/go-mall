package service

import (
	"context"

	product "github.com/PTS0118/go-mall/api/hertz_gen/api/product"
	"github.com/PTS0118/go-mall/api/infra/rpc"
	rpcproduct "github.com/PTS0118/go-mall/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/hertz/pkg/app"
)

type CreateProductService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewCreateProductService(Context context.Context, RequestContext *app.RequestContext) *CreateProductService {
	return &CreateProductService{RequestContext: RequestContext, Context: Context}
}

// @Summary 创建商品
// @Description 通过RPC调用创建商品
// @Tags Product
// @Accept json
// @Produce json
// @Param req body product.CreateProductReq true "创建商品请求"
// @Success 200 {object} product.CreateProductResp "成功响应"
// @Failure 400 {object} product.CreateProductResp "请求参数错误"
// @Router /product/create [post]
func (h *CreateProductService) Run(req *product.CreateProductReq) (resp *product.CreateProductResp, err error) {
	//判断参数是否为nil
	if req == nil {
		resp = &product.CreateProductResp{
			StatusCode: -1,
			StatusMsg:  "商品ID为空",
			Id:         0,
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
	data, err := rpc.ProductClient.CreateProduct(h.Context, &rpcproduct.CreateProductReq{Product: productData})
	if err != nil {
		resp = &product.CreateProductResp{
			StatusCode: data.GetCode(),
			StatusMsg:  data.GetMessage(),
			Id:         0,
		}
	} else {
		resp = &product.CreateProductResp{
			StatusCode: data.GetCode(),
			StatusMsg:  data.GetMessage(),
			Id:         data.ProductId,
		}
	}
	return resp, err
}
