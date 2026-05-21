package middleware

import (
	"net/http"
	"strings"

	"github.com/MuhAndriJP/ayo-api/internal/util"
	"github.com/gin-gonic/gin"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"success": false, "message": "token tidak ditemukan"})
			return
		}
		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := util.ParseJWT(tokenStr)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"success": false, "message": "token tidak valid atau kadaluarsa"})
			return
		}
		c.Set("admin_id", claims.AdminID)
		c.Next()
	}
}
