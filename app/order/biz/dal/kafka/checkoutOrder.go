package kafka

import (
	"context"
	"github.com/IBM/sarama"
	"github.com/PTS0118/go-mall/app/order/biz/model"
	"github.com/cloudwego/kitex/pkg/klog"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type ConsumerRta struct{}

func ConsumerToRequestRta(consumerGroup sarama.ConsumerGroup) {
	var (
		signals = make(chan os.Signal, 1)
	)
	signal.Notify(signals, syscall.SIGTERM, syscall.SIGINT, os.Interrupt)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// 监听终止信号
	go func() {
		<-signals
		klog.Info("收到终止信号，关闭消费者组")
		cancel()
		if err := consumerGroup.Close(); err != nil {
			klog.Error("关闭消费者组失败", zap.Error(err))
		}
	}()

	consumer := NewConsumerRta()
	klog.Info("[ConsumerToRequestRta] 消费者组启动")
	retryCount := 0
	for {
		select {
		case <-ctx.Done():
			klog.Info("消费者组正常退出")
			return
		default:
			if err := consumerGroup.Consume(ctx, []string{RealTopic}, consumer); err != nil {
				klog.Error("[ConsumerToRequestRta] 消费错误", zap.Error(err))
				// 指数退避重试
				retryCount++
				waitTime := time.Duration(retryCount*2) * time.Second
				time.Sleep(waitTime)
				continue
			}
		}
	}
}

func (c *ConsumerRta) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {
		err := model.MarkOrderStatus(2, string(message.Value))
		if err != nil {
			klog.Error("订单状态更新失败",
				zap.Error(err),
				zap.String("orderID", string(message.Value)),
			)
			// 可选：将错误消息发送到死信队列
			continue
		}
		session.MarkMessage(message, "")
		session.Commit()
	}
	return nil
}

func NewConsumerRta() *ConsumerRta {
	return &ConsumerRta{}
}

func (c *ConsumerRta) Setup(sarama.ConsumerGroupSession) error {
	klog.Info("消费者组初始化完成")
	return nil
}

func (c *ConsumerRta) Cleanup(sarama.ConsumerGroupSession) error {
	klog.Info("消费者组清理完成")
	return nil
}

//package kafka
//
//import (
//	"context"
//	"github.com/IBM/sarama"
//	"github.com/PTS0118/go-mall/app/order/biz/model"
//	"github.com/cloudwego/kitex/pkg/klog"
//	"go.uber.org/zap"
//	"os"
//	"os/signal"
//	"syscall"
//	"time"
//)
//
//type ConsumerRta struct {
//}
//
//func ConsumerToRequestRta(consumerGroup sarama.ConsumerGroup) {
//	var (
//		signals = make(chan os.Signal, 1)
//		//wg      = &sync.WaitGroup{}
//	)
//	signal.Notify(signals, syscall.SIGTERM, syscall.SIGINT, os.Interrupt)
//	//wg.Add(1)
//	ctx, cancel := context.WithCancel(context.Background())
//	defer cancel()
//
//	// 监听终止信号
//	go func() {
//		<-signals
//		klog.Info("收到终止信号，关闭消费者组")
//		cancel()
//		if err := consumerGroup.Close(); err != nil {
//			klog.Error("关闭消费者组失败", zap.Error(err))
//		}
//	}()
//
//	// 启动消费者协程
//	go func() {
//		consumer := NewConsumerRta()
//		klog.Info("[ConsumerToRequestRta] consumer group start")
//		for {
//			select {
//			case <-ctx.Done():
//				klog.Info("消费者组正常退出")
//				return
//			default:
//				if err := consumerGroup.Consume(context.Background(), []string{RealTopic}, consumer); err != nil {
//					klog.Error("[ConsumerToRequestRta] 消费组错误:", zap.Error(err))
//					continue
//				}
//				time.Sleep(2 * time.Second) // 等待一段时间后重试
//			}
//		}
//	}()
//
//	//go func() {
//	//	//defer wg.Done()
//	//
//	//	// 执行消费者组消费
//	//	for {
//	//		if err := consumerGroup.Consume(context.Background(), []string{RealTopic}, consumer); err != nil {
//	//			klog.Error("[ConsumerToRequestRta] 消费组错误:", zap.Error(err))
//	//			break
//	//		}
//	//		time.Sleep(2 * time.Second) // 等待一段时间后重试
//	//
//	//		// 检查是否接收到中断信号，如果是则退出循环
//	//		select {
//	//		case sin := <-signals:
//	//			klog.Info("get signal,{}", zap.Any("signal", sin))
//	//			return
//	//		default:
//	//		}
//	//	}
//	//}()
//	//wg.Wait()
//	klog.Info("[ConsumerToRequestRta] consumer end & exit")
//}
//
//func (c *ConsumerRta) ConsumeClaim(session sarama.ConsumerGroupSession,
//	claim sarama.ConsumerGroupClaim) error {
//	for message := range claim.Messages() {
//		// 消费逻辑
//		err := model.MarkOrderStatus(2, string(message.Value))
//		if err != nil {
//			klog.Info("kafka 定时取消失败，{}", err)
//			return err
//		}
//		session.MarkMessage(message, "")
//		session.Commit()
//	}
//	return nil
//}
//
//func NewConsumerRta() *ConsumerRta {
//	return &ConsumerRta{}
//}
//
//func (c *ConsumerRta) Setup(sarama.ConsumerGroupSession) error {
//	return nil
//}
//
//func (c *ConsumerRta) Cleanup(sarama.ConsumerGroupSession) error {
//	return nil
//}
