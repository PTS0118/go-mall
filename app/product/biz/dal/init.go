package dal

import (
	"github.com/PTS0118/go-mall/app/product/biz/dal/mysql"
	"github.com/PTS0118/go-mall/app/product/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
