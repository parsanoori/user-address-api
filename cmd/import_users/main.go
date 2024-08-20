package main

import (
	"UserAdresses/internals/database"
	"UserAdresses/internals/usersimport"
	"fmt"
)

func main() {
	// TODO: change it reading from .env
	err := database.Connect("postgres://user:password@localhost:5432/sika")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = database.Migrate()
	if err != nil {
		fmt.Println(err)
		return
	}
	// TODO: read from cli rgs
	err = usersimport.ImportUsers("users_data.json")
	if err != nil {
		fmt.Println(err)
		return
	}
}
