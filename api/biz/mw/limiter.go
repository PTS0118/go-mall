package mw

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"sync"
	"time"
)

// TokenBucket 令牌桶结构体
type TokenBucket struct {
	mu         sync.Mutex // 互斥锁保证并发安全
	capacity   int        // 桶容量（最大令牌数）
	tokens     int        // 当前令牌数
	refillRate float64    // 每秒填充速率（令牌/秒）
	lastRefill time.Time  // 上次填充时间
}

func CustomLimiter(capacity int, refillRate float64) app.HandlerFunc {
	// 为每个中间件创建独立的令牌桶实例
	bucket := NewTokenBucket(capacity, refillRate)
	return func(ctx context.Context, c *app.RequestContext) {
		if !bucket.Allow() {
			c.AbortWithStatusJSON(429, map[string]string{"error": "请求次数过多"})
			return
		}
		c.Next(ctx)
	}
}

// NewTokenBucket 令牌桶初始化方法
func NewTokenBucket(capacity int, refillRate float64) *TokenBucket {
	return &TokenBucket{
		capacity:   capacity,
		tokens:     capacity, // 初始时桶满
		refillRate: refillRate,
		lastRefill: time.Now(),
	}
}

// Allow 令牌获取逻辑
func (tb *TokenBucket) Allow() bool {
	tb.mu.Lock()
	defer tb.mu.Unlock()

	// 计算时间差并补充令牌
	now := time.Now()
	elapsed := now.Sub(tb.lastRefill).Seconds()
	tokensToAdd := int(elapsed * tb.refillRate)

	// 更新令牌数量（不超过容量）
	tb.tokens = min(tb.tokens+tokensToAdd, tb.capacity)

	// 更新填充时间（注意：只在补充令牌后更新）
	if tokensToAdd > 0 {
		tb.lastRefill = now
	}

	// 检查可用令牌
	if tb.tokens > 0 {
		tb.tokens--
		return true
	}
	return false
}

// 辅助函数（Go 1.21+ 可使用内置math.Min）
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
