package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	JWTSecret  string
	Port       string
	UploadDir  string
}

var App Config

func Load() error {
	_ = godotenv.Load()

	App = Config{
		DBHost:     getEnv("DB_HOST", "127.0.0.1"),
		DBPort:     getEnv("DB_PORT", "3306"),
		DBUser:     getEnv("DB_USER", "root"),
		DBPassword: getEnv("DB_PASSWORD", ""),
		DBName:     getEnv("DB_NAME", "ayo_db"),
		JWTSecret:  getEnv("JWT_SECRET", "change-me-in-production"),
		Port:       getEnv("PORT", "8080"),
		UploadDir:  getEnv("UPLOAD_DIR", "uploads"),
	}

	return nil
}

func DSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		App.DBUser, App.DBPassword, App.DBHost, App.DBPort, App.DBName)
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}

	return fallback
}
