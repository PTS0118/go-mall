package dal

import (
	"github.com/PTS0118/go-mall/app/order/biz/dal/kafka"
	"github.com/PTS0118/go-mall/app/order/biz/dal/mysql"
	"github.com/cloudwego/kitex/pkg/klog"
	"os"
	"os/signal"
	"syscall"
)

func Init() {
	//redis.Init()
	mysql.Init()
	kafka.Init()
	// 通过 select{} 或信号监听保持主线程运行
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	klog.Info("程序终止")

}
