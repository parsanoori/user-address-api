package usersimport

import (
	"UserAdresses/internals/controllers"
	"UserAdresses/internals/models"
	"encoding/json"
	"fmt"
	"os"
	"sync"
)

var userChan chan models.User
var errChan chan string

func readFile(filename string) error {
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
		var user models.User
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

func readUsers(filename string, wg *sync.WaitGroup) {
	userChan = make(chan models.User)
	defer close(userChan)
	errChan = make(chan string)
	defer close(errChan)
	wg.Done()
	err := readFile(filename)
	if err != nil {
		panic(err)
	}
}

func logErrors(wg *sync.WaitGroup) {
	defer wg.Done()
	for err := range errChan {
		fmt.Println(err)
	}
}

func addUsers(wg *sync.WaitGroup) {
	defer wg.Done()
	for user := range userChan {
		err := controllers.AddUser(&user)
		if err != nil {
			errChan <- fmt.Sprintf("Error saving user %s: %v", user.ID, err)
		}
	}
}

func ImportUsers(filename string) error {
	var channelWG, wg sync.WaitGroup

	channelWG.Add(1)
	go readUsers(filename, &channelWG)
	channelWG.Wait()

	wg.Add(1)
	go logErrors(&wg)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go addUsers(&wg)
	}

	wg.Wait()
	return nil
}
