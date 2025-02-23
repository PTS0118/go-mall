package service

import (
	"context"
	"fmt"
	"github.com/PTS0118/go-mall/app/product/biz/model"
	product "github.com/PTS0118/go-mall/rpc_gen/kitex_gen/product"
)

type DeleteProductService struct {
	ctx context.Context
} // NewDeleteProductService new DeleteProductService
func NewDeleteProductService(ctx context.Context) *DeleteProductService {
	return &DeleteProductService{ctx: ctx}
}

// Run create note info
func (s *DeleteProductService) Run(req *product.DeleteProductReq) (resp *product.DeleteProductResp, err error) {
	fmt.Println("ID：", req.Id)
	err = model.DeleteProduct(s.ctx, req.Id)
	if err != nil {
		resp = &product.DeleteProductResp{
			Code:    -1,
			Message: "删除商品失败",
		}
	} else {
		resp = &product.DeleteProductResp{
			Code:    0,
			Message: "删除商品成功",
		}
	}

	return resp, err
}
