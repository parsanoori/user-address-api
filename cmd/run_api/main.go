package main

import (
	"UserAdresses/internals/database"
	"UserAdresses/internals/handlers"
	"fmt"
	"os"
)

func main() {
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
	err = handlers.SetupAndStart("8080")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
