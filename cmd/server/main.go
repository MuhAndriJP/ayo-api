package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/MuhAndriJP/ayo-api/internal/config"
	"github.com/MuhAndriJP/ayo-api/internal/database"
	"github.com/MuhAndriJP/ayo-api/internal/handler"
	"github.com/MuhAndriJP/ayo-api/internal/repository"
	"github.com/MuhAndriJP/ayo-api/internal/router"
	"github.com/MuhAndriJP/ayo-api/internal/service"
)

func main() {
	if err := config.Load(); err != nil {
		log.Fatalf("config: %v", err)
	}

	if err := database.Connect(); err != nil {
		log.Fatalf("database: %v", err)
	}

	adminRepo := repository.NewAdminRepository(database.DB)

	authSvc := service.NewAuthService(adminRepo)

	authH := handler.NewAuthHandler(authSvc)

	r := router.Setup(authH)

	addr := fmt.Sprintf(":%s", config.App.Port)
	log.Printf("Server running on %s", addr)
	srv := &http.Server{Addr: addr, Handler: r}
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("server: %v", err)
	}
}
