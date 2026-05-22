package router

import (
	"path/filepath"
	"time"

	"github.com/MuhAndriJP/ayo-api/internal/config"
	"github.com/MuhAndriJP/ayo-api/internal/handler"
	"github.com/MuhAndriJP/ayo-api/internal/middleware"
	"github.com/gin-gonic/gin"
)

func Setup(
	authH *handler.AuthHandler,
	teamH *handler.TeamHandler,
	playerH *handler.PlayerHandler,
	matchH *handler.MatchHandler,
) *gin.Engine {
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

	teamPlayers := teams.Group("/:id/players")
	teamPlayers.GET("/", playerH.ListByTeam)
	teamPlayers.POST("/", playerH.Create)

	players := api.Group("/players")
	players.GET("/:id", playerH.Get)
	players.PUT("/:id", playerH.Update)
	players.DELETE("/:id", playerH.Delete)

	matches := api.Group("/matches")
	matches.GET("/", matchH.List)
	matches.GET("/:id", matchH.Get)
	matches.POST("/", matchH.Create)
	matches.PUT("/:id", matchH.Update)
	matches.DELETE("/:id", matchH.Delete)
	matches.POST("/:id/result", matchH.SaveResult)
	matches.PUT("/:id/result", matchH.SaveResult)
	matches.GET("/:id/report", matchH.GetReport)

	return r
}
