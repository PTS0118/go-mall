package dal

import (
	"github.com/PTS0118/go-mall/api/biz/dal/mysql"
	"github.com/PTS0118/go-mall/api/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
