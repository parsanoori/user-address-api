package database

import "UserAdresses/internals/models"

func Migrate() error {
	return db.AutoMigrate(&models.User{}, &models.Address{})
}
