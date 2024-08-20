package main

import (
	"UserAdresses/internals/database"
	"UserAdresses/internals/handlers"
	"UserAdresses/internals/usersimport"
	"fmt"
	"os"
)

func main() {
	// TODO: change it reading from .env
	err := database.Connect("postgres://user:password@localhost:5432/sika")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	err = database.Migrate()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// TODO: read from cli args
	err = usersimport.ImportUsers("users_data.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	err = handlers.SetupAndStart("8080")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
