package dal

import (
	"github.com/PTS0118/go-mall/app/order/biz/dal/mysql"
)

func Init() {
	//redis.Init()
	mysql.Init()
}
