package database

import (
	"UserAdresses/internals/logger"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func Connect(databaseURL string) error {
	logger.Log.Debug("Connecting to the database with URL", zap.String("url", databaseURL))
	var err error
	db, err = gorm.Open(postgres.Open(databaseURL), &gorm.Config{})

	if err != nil {
		logger.Log.Fatal("Failed to connect to the database", zap.Error(err))
	}
	logger.Log.Info("Database connected successfully.")
	return nil
}

func GetDB() *gorm.DB {
	return db
}
