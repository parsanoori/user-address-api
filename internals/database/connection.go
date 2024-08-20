package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var db *gorm.DB

func Connect(databaseURL string) error {
	var err error
	db, err = gorm.Open(postgres.Open(databaseURL), &gorm.Config{})

	if err != nil {
		return err
	}
	log.Println("Database connected successfully.")
	return nil
}

func GetDB() *gorm.DB {
	return db
}
