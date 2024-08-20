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

func AddUser(data *User) error {
	for _, add := range data.Addresses {
		add.UserID = data.ID
	}
	tx := DB.Begin()
	err := DB.Create(data).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func ReadFile(filename string, userChan chan<- User, errChan chan<- string) error {
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
		fmt.Println("here")
		var user User
		err = decoder.Decode(&user)
		if err != nil {
			errChan <- fmt.Sprintf("Error decoding JSON: %v", err)
		}
		userChan <- user
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

func ReadUsers(filename string, userChan chan<- User, errChan chan<- string) {
	err := ReadFile(filename, userChan, errChan)
	if err != nil {
		panic(err)
	}
}

func LogErrors(errChan <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for err := range errChan {
		fmt.Println(err)
	}
}

func AddUsers(userChan <-chan User, errChan chan<- string, wg *sync.WaitGroup) {
	for user := range userChan {
		fmt.Println(user.ID)
		err := AddUser(&user)
		if err != nil {
			errChan <- fmt.Sprintf("Error saving user %s: %v", user.ID, err)
		}
	}
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
	errChan := make(chan string, 100)

	defer close(userChan)
	defer close(errChan)

	var wg sync.WaitGroup

	go ReadUsers("users_data.json", userChan, errChan)

	wg.Add(1)
	go AddUsers(userChan, errChan, &wg)

	wg.Add(1)
	go LogErrors(errChan, &wg)

	wg.Wait()
}
