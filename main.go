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

var userChan chan User
var errChan chan string

func ReadFile(filename string) error {
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
			errChan <- fmt.Sprintf("Error decoding JSON: %v", err)
		}
		userChan <- user
	}
	fmt.Println("end of reading the users")

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

func ReadUsers(filename string, wg *sync.WaitGroup) {
	userChan = make(chan User)
	defer close(userChan)
	errChan = make(chan string)
	defer close(errChan)
	wg.Done()
	err := ReadFile(filename)
	if err != nil {
		panic(err)
	}
}

func LogErrors(wg *sync.WaitGroup) {
	defer wg.Done()
	for err := range errChan {
		fmt.Println(err)
	}
	fmt.Println("end of logging the errors")
}

func AddUsers(wg *sync.WaitGroup) {
	defer wg.Done()
	for user := range userChan {
		err := AddUser(&user)
		if err != nil {
			errChan <- fmt.Sprintf("Error saving user %s: %v", user.ID, err)
		}
	}
	fmt.Println("end of adding the users")
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

	var wgChannelMaking sync.WaitGroup
	wgChannelMaking.Add(1)
	go ReadUsers("users_data.json", &wgChannelMaking)
	wgChannelMaking.Wait()

	var wg sync.WaitGroup
	wg.Add(1)
	go AddUsers(&wg)

	wg.Add(1)
	go LogErrors(&wg)

	wg.Wait()
}
