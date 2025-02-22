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
