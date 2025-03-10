// Code generated by hertz generator.

package main

import (
	"context"
	"time"

	"github.com/PTS0118/go-mall/api/biz/dal"
	"github.com/PTS0118/go-mall/api/biz/mw"
	_ "github.com/PTS0118/go-mall/api/docs"
	"github.com/PTS0118/go-mall/api/infra/rpc"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/hertz-contrib/swagger"
	swaggerFiles "github.com/swaggo/files"

	"github.com/PTS0118/go-mall/api/biz/router"
	"github.com/PTS0118/go-mall/api/conf"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/middlewares/server/recovery"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/hertz-contrib/cors"
	"github.com/hertz-contrib/gzip"
	"github.com/hertz-contrib/logger/accesslog"
	hertzlogrus "github.com/hertz-contrib/logger/logrus"
	"github.com/hertz-contrib/pprof"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func main() {
	// init dal
	dal.Init()
	rpc.InitClient()
	mw.JWTInit()
	address := conf.GetConf().Hertz.Address
	h := server.New(server.WithHostPorts(address))
	// 注册中间件（确保 enforcer 已正确创建）
	_, err := mw.InitCasbin()
	if err != nil {
		klog.Error("初始化Casbin失败")
		klog.Error(err)
		return
	}
	registerMiddleware(h)

	// add a ping route to test
	h.GET("/ping", func(c context.Context, ctx *app.RequestContext) {
		ctx.JSON(consts.StatusOK, utils.H{"ping": "pong"})
	})

	router.GeneratedRegister(h)
	h.Spin()
}

func registerMiddleware(h *server.Hertz) {
	// log
	logger := hertzlogrus.NewLogger()
	hlog.SetLogger(logger)
	hlog.SetLevel(conf.LogLevel())
	asyncWriter := &zapcore.BufferedWriteSyncer{
		WS: zapcore.AddSync(&lumberjack.Logger{
			Filename:   conf.GetConf().Hertz.LogFileName,
			MaxSize:    conf.GetConf().Hertz.LogMaxSize,
			MaxBackups: conf.GetConf().Hertz.LogMaxBackups,
			MaxAge:     conf.GetConf().Hertz.LogMaxAge,
		}),
		FlushInterval: time.Minute,
	}
	hlog.SetOutput(asyncWriter)
	h.OnShutdown = append(h.OnShutdown, func(ctx context.Context) {
		asyncWriter.Sync()
	})

	// pprof
	if conf.GetConf().Hertz.EnablePprof {
		pprof.Register(h)
	}

	// gzip
	if conf.GetConf().Hertz.EnableGzip {
		h.Use(gzip.Gzip(gzip.DefaultCompression))
	}

	// access log
	if conf.GetConf().Hertz.EnableAccessLog {
		h.Use(accesslog.New())
	}

	// recovery
	h.Use(recovery.Recovery())

	// cores
	h.Use(cors.Default())

	//swagger
	url := swagger.URL("http://localhost:8080/swagger/doc.json") // The url pointing to API definition
	h.GET("/swagger/*any", swagger.WrapHandler(swaggerFiles.Handler, url))
}
