package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
	"net/http"
	"time"
)

// RateLimitMiddleware 每fillInterval 秒 自动添加 cap 个数的令牌 注意参数 要用Time.Second 否则就是2ns了
func RateLimitMiddleware(fillInterval time.Duration, cap int64) func(c *gin.Context) {
	bucket := ratelimit.NewBucket(fillInterval, cap)
	return func(c *gin.Context) {
		// 如果取不到令牌 就直接返回吧
		if bucket.TakeAvailable(1) == 0 {
			c.String(http.StatusOK, "rate limit")
			c.Abort()
			return
		}
		c.Next()
	}
}
