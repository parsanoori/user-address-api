package controllers

import (
	"UserAdresses/internals/database"
	"UserAdresses/internals/models"
	"fmt"
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
	return nil
}

func GetUser(id string) (*models.User, error) {
	user := models.User{}
	err := database.GetDB().Preload("Addresses").First(&user, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func UpdateUser(data *models.User) error {
	tx := database.GetDB().Begin()
	err := database.GetDB().Save(data).Error
	if err != nil {
		tx.Rollback()
	}
	tx.Commit()
	return nil
}

func DeleteUser(id string) error {
	tx := database.GetDB().Begin()
	err := database.GetDB().Delete(&models.User{}, "id = ?", id).Error
	if err != nil {
		tx.Rollback()
	}
	tx.Commit()
	return nil
}
