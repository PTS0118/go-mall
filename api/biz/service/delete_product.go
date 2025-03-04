package service

import (
	"context"

	product "github.com/PTS0118/go-mall/api/hertz_gen/api/product"
	rpc "github.com/PTS0118/go-mall/api/infra/rpc"
	rpcproduct "github.com/PTS0118/go-mall/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/hertz/pkg/app"
)

type DeleteProductService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewDeleteProductService(Context context.Context, RequestContext *app.RequestContext) *DeleteProductService {
	return &DeleteProductService{RequestContext: RequestContext, Context: Context}
}

// @Summary 删除商品
// @Description 通过RPC调用删除商品
// @Tags Product
// @Accept json
// @Produce json
// @Param req body product.DeleteProductReq true "删除商品请求"
// @Success 200 {object} product.DeleteProductResp "成功响应"
// @Failure 400 {object} product.DeleteProductResp "请求参数错误"
// @Router /product/delete [post]
func (h *DeleteProductService) Run(req *product.DeleteProductReq) (resp *product.DeleteProductResp, err error) {
	//判断参数是否为nil
	if req == nil {
		resp = &product.DeleteProductResp{
			StatusCode: -1,
			StatusMsg:  "商品ID为空",
		}
		return resp, nil
	}
	data, err := rpc.ProductClient.DeleteProduct(h.Context, &rpcproduct.DeleteProductReq{Id: req.GetId()})
	resp = &product.DeleteProductResp{
		StatusCode: data.GetCode(),
		StatusMsg:  data.GetMessage(),
	}
	return resp, err
}
