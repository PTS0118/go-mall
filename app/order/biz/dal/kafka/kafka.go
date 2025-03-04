package kafka

import (
	"context"
	"fmt"
	"github.com/IBM/sarama"
	"github.com/PTS0118/go-mall/app/order/conf"
	"github.com/cloudwego/kitex/pkg/klog"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type KafkaConfig struct {
	BrokerList []string
	Topic      []string
	GroupId    []string
	Cfg        *sarama.Config
	PemPath    string
	KeyPath    string
	CaPemPath  string
}

var (
	Producer           sarama.SyncProducer
	ConsumerGroupReal  sarama.ConsumerGroup
	ConsumerGroupDelay sarama.ConsumerGroup
	KafkaDelayQueue    *KafkaDelayQueueProducer
)

const (
	DelayTime  = time.Minute * 5
	DelayTopic = "delayTopic"
	RealTopic  = "realTopic"
)

// KafkaDelayQueueProducer 延迟队列生产者，包含了生产者和延迟服务
type KafkaDelayQueueProducer struct {
	producer   sarama.SyncProducer // 生产者
	delayTopic string              // 延迟服务主题
}

func Init() {
	klog.Info("[Init] 开始初始化 kafkas")
	// 1.获取kafka地址
	brokerList := conf.GetConf().Kafka.Address
	//2. 注册groupId
	groupId := []string{"real_group", "delay_group"}
	// 3. 创建消费者组配置
	config := sarama.NewConfig()
	config.Version = sarama.V2_5_0_0 // 设置与 Kafka 版本一致
	config.Producer.Return.Successes = true
	config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRange
	config.Consumer.Offsets.Initial = sarama.OffsetOldest // 从最早消息开始消费
	config.Consumer.Offsets.AutoCommit.Enable = true      // 启用自动提交 Offset
	config.Consumer.Offsets.AutoCommit.Interval = 1 * time.Second

	cfg := KafkaConfig{
		BrokerList: brokerList,
		Topic:      nil,
		GroupId:    groupId,
		Cfg:        config,
		PemPath:    "",
		KeyPath:    "",
		CaPemPath:  "",
	}
	err := NewKafkaConfig(cfg)
	if err != nil {
		fmt.Printf("NewKafkaConfig：%v\n", err) // 修改：增加换行符
		return
	}
	// 启动延时队列
	GetKafkaDelayQueue()
	// 启动消费服务 -- 定时取消订单
	go func() {

	}()
	//wg.Wait() // 等待消费者组启动完成（需在 ConsumerToRequestRta 内部标记
	//ConsumerToRequestRta(ConsumerGroupReal)
	klog.Info("[Init] 初始化完成") // 添加日志
}

func NewKafkaConfig(cfg KafkaConfig) (err error) {
	Producer, err = sarama.NewSyncProducer(cfg.BrokerList, cfg.Cfg)
	if err != nil {
		klog.Error("[NewKafkaConfig] 创建同步生产者失败", zap.Error(err)) // 添加日志
		return err
	}
	ConsumerGroupReal, err = sarama.NewConsumerGroup(cfg.BrokerList, cfg.GroupId[0], cfg.Cfg)
	if err != nil {
		klog.Error("[NewKafkaConfig] 创建实时消费者组失败", zap.Error(err)) // 添加日志
		return err
	}
	ConsumerGroupDelay, err = sarama.NewConsumerGroup(cfg.BrokerList, cfg.GroupId[1], cfg.Cfg)
	if err != nil {
		klog.Error("[NewKafkaConfig] 创建延迟消费者组失败", zap.Error(err)) // 添加日志
		return err
	}

	klog.Info("[NewKafkaConfig] 配置初始化成功") // 添加日志
	return nil
}

func GetKafkaDelayQueue() {
	klog.Info("[GetKafkaDelayQueue] 创建延迟队列成功") // 添加日志
	KafkaDelayQueue = NewKafkaDelayQueueProducer(Producer, ConsumerGroupDelay, DelayTime, DelayTopic, RealTopic)
}

// NewKafkaDelayQueueProducer 创建延迟队列生产者
// producer 生产者
// delayServiceConsumerGroup 延迟服务消费者组
// delayTime 延迟时间
// delayTopic 延迟服务主题
// realTopic 真实队列主题
func NewKafkaDelayQueueProducer(producer sarama.SyncProducer, delayServiceConsumerGroup sarama.ConsumerGroup,
	delayTime time.Duration, delayTopic, realTopic string) *KafkaDelayQueueProducer {
	var (
		signals = make(chan os.Signal, 1)
	)
	signal.Notify(signals, syscall.SIGTERM, syscall.SIGINT, os.Interrupt)
	// 启动延迟服务
	consumer := NewDelayServiceConsumer(producer, delayTime, realTopic)
	klog.Info("[NewKafkaDelayQueueProducer] kafka延迟队列消费开始")
	go func() {
		for {
			if err := delayServiceConsumerGroup.Consume(context.Background(),
				[]string{delayTopic}, consumer); err != nil {
				klog.Error("[NewKafkaDelayQueueProducer] delay queue consumer failed,err: ", zap.Error(err))
				break
			}
			time.Sleep(2 * time.Second)
			klog.Info("[NewKafkaDelayQueueProducer] 检测消费函数是否一直执行")
			// 检查是否接收到中断信号，如果是则退出循环
			select {
			case sin := <-signals:
				klog.Info("[NewKafkaDelayQueueProducer] get signal,", zap.Any("signal", sin))
				return
			default:
			}
		}
		klog.Info("[NewKafkaDelayQueueProducer] consumer func exit")
	}()
	klog.Info("[NewKafkaDelayQueueProducer] return KafkaDelayQueueProducer")

	return &KafkaDelayQueueProducer{
		producer:   producer,
		delayTopic: delayTopic,
	}
}

// SendMessage 发送消息
func (q *KafkaDelayQueueProducer) SendMessage(msg *sarama.ProducerMessage) (partition int32, offset int64, err error) {
	msg.Topic = q.delayTopic
	partition, offset, err = q.producer.SendMessage(msg)
	if err != nil {
		klog.Error("[SendMessage] 发送消息失败", zap.Error(err)) // 添加日志
	}
	return partition, offset, err
}

// DelayServiceConsumer 延迟服务消费者
type DelayServiceConsumer struct {
	producer  sarama.SyncProducer
	delay     time.Duration
	realTopic string
}

func NewDelayServiceConsumer(producer sarama.SyncProducer, delay time.Duration,
	realTopic string) *DelayServiceConsumer {
	return &DelayServiceConsumer{
		producer:  producer,
		delay:     delay,
		realTopic: realTopic,
	}
}

func (c *DelayServiceConsumer) ConsumeClaim(session sarama.ConsumerGroupSession,
	claim sarama.ConsumerGroupClaim) error {
	klog.Info("[ConsumeClaim] 开始处理消息")
	for message := range claim.Messages() {
		now := time.Now()
		klog.Info("[ConsumeClaim] 处理消息",
			zap.Any("send real topic res", now.Sub(message.Timestamp) >= c.delay),
			zap.Any("message.Timestamp", message.Timestamp),
			zap.Any("c.delay", c.delay),
			zap.Any("claim.Messages len", len(claim.Messages())),
			zap.Any("sub:", now.Sub(message.Timestamp)),
			zap.Any("meskey:", message.Key),
			zap.Any("message:", string(message.Value)),
		)
		// 如果消息已经超时，把消息发送到真实队列
		if now.Sub(message.Timestamp) >= c.delay {
			klog.Info("[ConsumeClaim] 发送真实队列", zap.Any("msg: ", string(message.Value)))
			_, _, err := c.producer.SendMessage(&sarama.ProducerMessage{
				Topic:     c.realTopic,
				Timestamp: message.Timestamp,
				Key:       sarama.ByteEncoder(message.Key),
				Value:     sarama.ByteEncoder(message.Value),
			})
			if err != nil {
				klog.Error("[ConsumeClaim] 延时队列转发真实队列失败", zap.Error(err)) // 修改：移除多余的大括号
				return nil
			}
			session.MarkMessage(message, "")
			continue
		}
		// 否则休眠一秒
		time.Sleep(time.Second)
	}

	klog.Info("[ConsumeClaim] 结束处理消息")
	return nil
}

func (c *DelayServiceConsumer) Setup(sarama.ConsumerGroupSession) error {
	klog.Info("[Setup] 消费者组设置完成")
	return nil
}

func (c *DelayServiceConsumer) Cleanup(sarama.ConsumerGroupSession) error {
	klog.Info("[Cleanup] 消费者组清理完成")
	return nil
}
