package main

import (
	"encoding/json"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"sync"
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

func UserToDB(data *User) error {
	for _, add := range data.Addresses {
		add.UserID = data.ID
	}
	err := DB.Create(data).Error
	if err != nil {
		return err
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

func ReadFile(filename string, userChan chan<- User, resultChan chan<- string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// this way we can read the line by JSON object
	decoder := json.NewDecoder(file)

	var t json.Token

	// Read open bracket
	t, err = decoder.Token()
	if err != nil {
		return err
	}

	if t != json.Delim('[') {
		return fmt.Errorf("expected '[' but got %v", t)
	}

	for decoder.More() {
		var user User
		err = decoder.Decode(&user)
		if err != nil {
			resultChan <- fmt.Sprintf("Error decoding JSON: %v", err)
		}
		userChan <- user
		resultChan <- fmt.Sprintf("User %s has been decoded and added to the channel", user.ID)
	}

	// Read closing bracket
	t, err = decoder.Token()
	if err != nil {
		return err
	}

	if t != json.Delim(']') {
		return fmt.Errorf("expected ']' but got %v", t)
	}
	return nil
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
	err = DB.AutoMigrate(&User{}, &Address{})
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	userChan := make(chan User, 100)
	resultChan := make(chan string, 100)

	var wg sync.WaitGroup

	go func() {
		err = ReadFile("users_data.json", userChan, resultChan)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		close(userChan)
		close(resultChan)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for user := range userChan {
			fmt.Println("User", user.ID, "is being processed")
		}
	}()

	wg.Wait()
}
