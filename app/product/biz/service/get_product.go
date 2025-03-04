package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/PTS0118/go-mall/app/product/biz/dal/redis"
	"github.com/PTS0118/go-mall/app/product/biz/model"
	product "github.com/PTS0118/go-mall/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/kitex/pkg/klog"
	"strings"
	"time"
)

type GetProductService struct {
	ctx context.Context
} // NewGetProductService new GetProductService
func NewGetProductService(ctx context.Context) *GetProductService {
	return &GetProductService{ctx: ctx}
}

// Run create note info
func (s *GetProductService) Run(req *product.GetProductReq) (resp *product.GetProductResp, err error) {
	//查询缓存
	redisKey := fmt.Sprintf("product:%d", req.GetId())
	// 1. 先查Redis缓存
	rData, err := redis.RedisClient.Get(s.ctx, redisKey).Result()
	if err == nil && rData != "" {
		var rProduct model.Product
		if err = json.Unmarshal([]byte(rData), &rProduct); err == nil {
			rCategorylist := strings.Split(rProduct.Categories, ",")
			return &product.GetProductResp{
				Code:    0,
				Message: "获取商品成功",
				Product: &product.Product{
					Id:          rProduct.Id,
					Name:        rProduct.Name,
					Description: rProduct.Description,
					Picture:     rProduct.Picture,
					Price:       rProduct.Price,
					Categories:  rCategorylist,
				},
			}, nil
		}
	}

	//2.查询缓存失败
	data, err := model.GetProductById(s.ctx, req.GetId())
	if err != nil {
		resp = &product.GetProductResp{
			Code:    -1,
			Message: "获取商品失败",
			Product: &product.Product{},
		}
		klog.Error("商品查询失败：%v", err)
	} else if data == nil {
		resp = &product.GetProductResp{
			Code:    -1,
			Message: "商品不存在",
			Product: &product.Product{},
		}
		klog.Info("商品查询失败：商品不存在")
	} else {
		categoryList := strings.Split(data.Categories, ",")
		resp = &product.GetProductResp{
			Code:    0,
			Message: "获取商品成功",
			Product: &product.Product{
				Id:          data.Id,
				Name:        data.Name,
				Description: data.Description,
				Picture:     data.Picture,
				Price:       data.Price,
				Categories:  categoryList,
				Stock:       int32(data.Stock),
			},
		}
	}

	// 3. 写入Redis并设置逻辑过期时间（防缓存击穿）
	go func() {
		jsonData, _ := json.Marshal(data)
		// 设置缓存30分钟过期，但实际逻辑过期时间为24小时
		if err = redis.RedisClient.SetEx(s.ctx, redisKey, string(jsonData), 1800*time.Second).Err(); err != nil {
			klog.Info("Redis 设置商品缓存失败，商品ID:{} ,失败error: {}", req.GetId(), err)
		}
	}()

	return resp, err
}
