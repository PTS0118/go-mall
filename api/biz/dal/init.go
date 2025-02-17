package dal

import (
	"github.com/PTS0118/go-mall/api/biz/dal/mysql"
)

func Init() {
	//redis.Init()
	mysql.Init()
}
