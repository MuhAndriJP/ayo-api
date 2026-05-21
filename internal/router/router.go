package router

import (
	"time"

	"github.com/MuhAndriJP/ayo-api/internal/handler"
	"github.com/MuhAndriJP/ayo-api/internal/middleware"
	"github.com/gin-gonic/gin"
)

func Setup(authH *handler.AuthHandler) *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())

	v1 := r.Group("/api/v1")

	auth := v1.Group("/auth")
	auth.POST("/register", authH.Register)
	auth.POST("/login", middleware.LoginRateLimit(10, 15*time.Minute), authH.Login)

	return r
}
