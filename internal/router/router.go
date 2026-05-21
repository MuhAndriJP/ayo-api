package router

import (
	"path/filepath"
	"time"

	"github.com/MuhAndriJP/ayo-api/internal/config"
	"github.com/MuhAndriJP/ayo-api/internal/handler"
	"github.com/MuhAndriJP/ayo-api/internal/middleware"
	"github.com/gin-gonic/gin"
)

func Setup(authH *handler.AuthHandler, teamH *handler.TeamHandler) *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())

	uploadDir := filepath.Join(config.App.UploadDir, "teams")
	r.Static("/uploads/teams", uploadDir)

	v1 := r.Group("/api/v1")

	auth := v1.Group("/auth")
	auth.POST("/register", authH.Register)
	auth.POST("/login", middleware.LoginRateLimit(10, 15*time.Minute), authH.Login)

	api := v1.Group("")
	api.Use(middleware.JWTAuth())

	teams := api.Group("/teams")
	teams.GET("/", teamH.List)
	teams.GET("/:id", teamH.Get)
	teams.POST("/", teamH.Create)
	teams.PUT("/:id", teamH.Update)
	teams.DELETE("/:id", teamH.Delete)

	return r
}
