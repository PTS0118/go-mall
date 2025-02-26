package mysql

import (
	"time"

	"github.com/PTS0118/go-mall/app/cart/conf"
	"github.com/cloudwego/kitex/pkg/klog"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	DB  *gorm.DB
	err error
)

type KLogWriter struct{}

func (w KLogWriter) Printf(format string, v ...interface{}) {
	klog.Infof(format, v...)
}

func Init() {
	// 创建一个新的 GORM logger 实例
	newLogger := logger.New(
		KLogWriter{},
		logger.Config{
			SlowThreshold: time.Second, // 慢查询阈值
			LogLevel:      logger.Info, // 日志级别
			Colorful:      false,       // 禁用彩色打印
		},
	)
	DB, err = gorm.Open(mysql.Open(conf.GetConf().MySQL.DSN),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
			Logger:                 newLogger,
		},
	)
	if err != nil {
		panic(err)
	}
}
