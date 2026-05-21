package database

import (
	"fmt"

	"github.com/MuhAndriJP/ayo-api/internal/config"
	"github.com/MuhAndriJP/ayo-api/internal/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Connect() error {
	var err error
	DB, err = gorm.Open(mysql.Open(config.DSN()), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error),
	})
	if err != nil {
		return fmt.Errorf("database connect: %w", err)
	}

	if err = migrate(); err != nil {
		return fmt.Errorf("migrate: %w", err)
	}

	return nil
}

func migrate() error {
	return DB.AutoMigrate(
		&model.Admin{},
		&model.Team{},
		&model.Player{},
		&model.Match{},
		&model.Goal{},
	)
}
