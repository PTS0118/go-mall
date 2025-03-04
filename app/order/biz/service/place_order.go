package service

import (
	"context"
	"github.com/IBM/sarama"
	"github.com/PTS0118/go-mall/app/order/biz/dal/kafka"
	"github.com/PTS0118/go-mall/app/order/biz/dal/mysql"
	"github.com/PTS0118/go-mall/app/order/biz/model"
	"github.com/PTS0118/go-mall/app/order/biz/utils"
	order "github.com/PTS0118/go-mall/rpc_gen/kitex_gen/order"
	"github.com/cloudwego/kitex/pkg/klog"
	"time"
)

type PlaceOrderService struct {
	ctx context.Context
} // NewPlaceOrderService new PlaceOrderService
func NewPlaceOrderService(ctx context.Context) *PlaceOrderService {
	return &PlaceOrderService{ctx: ctx}
}

// Run create note info
func (s *PlaceOrderService) Run(req *order.PlaceOrderReq) (resp *order.PlaceOrderResp, err error) {
	//新建订单
	//生成随机订单号
	orderId := utils.GenerateOrderID(int(req.UserId))
	//创建订单项
	orderItems := make([]*model.OrderItem, len(req.GetOrderItems()))
	totalPrice := 0.0
	for key, value := range req.OrderItems {
		orderItems[key] = &model.OrderItem{
			ProductId:  int(value.ProductId),
			OrderId:    orderId,
			TotalPrice: float64(value.TotalPrice),
			UnitPrice:  float64(value.UnitPrice),
			Count:      int(value.Count),
		}
		totalPrice = totalPrice + float64(value.TotalPrice)
	}

	orderData := &model.Order{
		UserId:     int(req.UserId),
		OrderId:    orderId,
		TotalPrice: totalPrice,
		Status:     0, //初始状态
		Telephone:  req.Telephone,
		AddressId:  int(req.AddressId),
	}

	// 开始事务
	tx := mysql.DB.Begin()
	// 注意：最好在这里使用 defer 来确保即使发生 panic 也能正确回滚
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	//创建订单
	_, err = model.CreateOrder(s.ctx, orderData)
	if err != nil {
		resp = &order.PlaceOrderResp{
			Code:    -1,
			Message: "创建订单失败",
			OrderId: "",
		}
		tx.Rollback()
		return resp, err
	}

	//创建订单项
	for _, value := range orderItems {
		_, err = model.CreateOrderItem(s.ctx, value)
		if err != nil {
			resp = &order.PlaceOrderResp{
				Code:    -1,
				Message: "创建订单项失败",
				OrderId: "",
			}
			tx.Rollback()
			return resp, err
		}
	}

	// 提交事务
	if err = tx.Commit().Error; err != nil {
		resp = &order.PlaceOrderResp{
			Code:    -1,
			Message: "提交订单失败",
			OrderId: "",
		}
		tx.Rollback()
		return resp, err
	}
	resp = &order.PlaceOrderResp{
		Code:    0,
		Message: "创建订单成功",
		OrderId: orderId,
	}

	// 发送消息到kafka延迟队列，实现定时订单未支付取消
	msg := &sarama.ProducerMessage{
		Topic:     kafka.DelayTopic,
		Timestamp: time.Now(),
		Key:       sarama.StringEncoder(orderId),
		Value:     sarama.StringEncoder(orderId),
	}
	klog.Info("kafka 准备发送消息, message:", msg)
	_, _, err = kafka.KafkaDelayQueue.SendMessage(msg)
	if err != nil {
		klog.Info("kafka 发送消息到延时队列失败, error:{}", err)
	} else {
		klog.Info("kafka 发送消息成功, message:", msg)
	}

	return resp, err
}
