package middleware

import (
	"net/http"
	"sync"
	"time"

	"github.com/MuhAndriJP/ayo-api/internal/util"
	"github.com/gin-gonic/gin"
)

type visitor struct {
	count   int
	resetAt time.Time
}

var (
	visitors = map[string]*visitor{}
	mu       sync.Mutex
)

func LoginRateLimit(maxReq int, window time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()
		mu.Lock()
		v, ok := visitors[ip]
		if !ok || time.Now().After(v.resetAt) {
			visitors[ip] = &visitor{count: 1, resetAt: time.Now().Add(window)}
			mu.Unlock()
			c.Next()
			return
		}
		v.count++
		if v.count > maxReq {
			mu.Unlock()
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"success": false,
				"message": util.ErrTooManyRequests.Error(),
			})
			return
		}
		mu.Unlock()
		c.Next()
	}
}
