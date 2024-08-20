package main

import (
	"encoding/json"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB

func ConnectDB() (*gorm.DB, error) {
	dsn := "host=localhost user=user password=password dbname=sika port=5432"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func AdrressToDB(data *Address, userID string) error {
	err := DB.Create(data).Error
	if err != nil {
		return err
	}
	return nil
}

func UserToDB(data *User) error {
	addresses := data.Addresses
	err := DB.Create(&data).Error
	if err != nil {
		return err
	}
	for _, add := range addresses {
		err := AdrressToDB(&add, data.ID)
		if err != nil {
			return err
		}
	}
	return nil
}

func ReadUserFromJSON(data string) (*User, error) {
	var userData User
	err := json.Unmarshal([]byte(data), &userData)
	if err != nil {
		return nil, err
	}
	return &userData, nil
}

func ReadUserFile() (string, error) {
	f, err := os.ReadFile("user.json")
	if err != nil {
		return "", err
	}
	return string(f), nil
}

func main() {
	var err error
	DB, err = ConnectDB()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	err = DB.AutoMigrate(&User{})
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	err = DB.AutoMigrate(&Address{})
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	var user string
	user, err = ReadUserFile()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	var uo *User
	uo, err = ReadUserFromJSON(user)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	err = UserToDB(uo)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
