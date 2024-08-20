package controllers

import (
	"UserAdresses/internals/database"
	"UserAdresses/internals/logger"
	"UserAdresses/internals/models"
	"fmt"
	"go.uber.org/zap"
)

func AddUser(data *models.User) error {
	for _, add := range data.Addresses {
		add.UserID = data.ID
	}
	if _, err := GetUser(data.ID); err == nil {
		return fmt.Errorf("User with id %s already exists", data.ID)
	}
	tx := database.GetDB().Begin()
	err := database.GetDB().Create(data).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	logger.Log.Debug("User added successfully.", zap.Any("user", data))
	return nil
}

func GetUser(id string) (*models.User, error) {
	user := models.User{}
	err := database.GetDB().Preload("Addresses").First(&user, "id = ?", id).Error
	if err != nil {
		logger.Log.Error("Failed to get user", zap.Error(err))
		return nil, err
	}
	logger.Log.Debug("User fetched successfully.", zap.Any("user", user))
	return &user, nil
}

func UpdateUser(data *models.User) error {
	tx := database.GetDB().Begin()
	err := database.GetDB().Save(data).Error
	if err != nil {
		logger.Log.Error("Failed to update user", zap.Error(err))
		tx.Rollback()
		return err
	}
	tx.Commit()
	logger.Log.Debug("User updated successfully.", zap.Any("user", data))
	return nil
}

func DeleteUser(id string) error {
	tx := database.GetDB().Begin()
	err := database.GetDB().Delete(&models.User{}, "id = ?", id).Error
	if err != nil {
		logger.Log.Error("Failed to delete user", zap.Error(err))
		tx.Rollback()
		return err
	}
	tx.Commit()
	logger.Log.Debug("User deleted successfully.", zap.String("id", id))
	return nil
}
