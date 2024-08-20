package database

import (
	"UserAdresses/internals/logger"
	"UserAdresses/internals/models"
	"go.uber.org/zap"
)

func Migrate() error {
	err := db.AutoMigrate(&models.User{}, &models.Address{})
	if err != nil {
		logger.Log.Fatal("Failed to migrate database", zap.Error(err))
	}
	return nil
}
