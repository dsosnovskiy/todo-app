package database

import (
	"fmt"
	"log"
	"todo-app/config"
	"todo-app/internal/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect(cfg *config.Config) (*gorm.DB, error) {
	connectionString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.User,
		cfg.Database.DBName,
		cfg.Database.Password,
		cfg.Database.SSLMode)

	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	err = db.AutoMigrate(&model.Task{})
	if err != nil {
		log.Fatalf("Failed to apply migrations: %v", err)
		return nil, err
	}

	return db, nil
}
