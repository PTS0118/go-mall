package utils

import (
	"github.com/PTS0118/go-mall/api/biz/model"
	"github.com/cloudwego/hertz/pkg/app"
)

func GetUserId(c *app.RequestContext) (userId int) {
	// 获取用户信息
	userInterface, exist := c.Get("uid")

	if !exist {
		// 处理错误：未找到用户信息
		c.JSON(401, map[string]interface{}{"error": "Unauthorized"})
		return 0
	}

	// 类型断言并获取ID
	user, ok := userInterface.(*model.User)
	if !ok {
		// 处理类型断言失败的情况
		c.JSON(500, map[string]interface{}{"error": "Internal Server Error"})
		return 0
	}

	userId = user.Base.Id
	return userId
}
