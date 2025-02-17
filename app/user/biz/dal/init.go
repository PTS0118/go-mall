package dal

import (
	"github.com/PTS0118/go-mall/app/user/biz/dal/mysql"
)

func Init() {
	//redis.Init()
	mysql.Init()
}
