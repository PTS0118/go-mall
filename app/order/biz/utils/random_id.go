package utils

import (
	"fmt"
	"math/rand"
	"time"
)

// GenerateOrderID 生成固定长度的订单ID
func GenerateOrderID(userID int) string {
	const (
		timestampLength = 13                                                // 时间戳精确到毫秒，长度为13
		randomLength    = 4                                                 // 随机数长度
		userIDLength    = 8                                                 // 用户ID填充后的长度
		totalLength     = timestampLength + randomLength + userIDLength + 2 // 加上分隔符长度
	)

	// 获取当前时间戳
	timestamp := time.Now().UnixNano() / int64(time.Millisecond)
	// 生成随机数
	randNum := rand.Intn(9999-1000) + 1000 // 生成4位随机数
	// 将userID转换为字符串并用前导零填充到8位
	userIDStr := fmt.Sprintf("%08d", userID)
	// 组合订单ID
	orderID := fmt.Sprintf("%d-%04d-%s", timestamp, randNum, userIDStr)
	if len(orderID) > totalLength {
		// 如果由于某些原因导致长度超出预期，可以采取截断或其他措施
		orderID = orderID[:totalLength]
	}
	return orderID
}
