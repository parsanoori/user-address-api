package cmd

import (
	"UserAdresses/internals/config"
	"UserAdresses/internals/database"
	"UserAdresses/internals/handlers"
	"UserAdresses/internals/logger"
	"UserAdresses/internals/usersimport"
	"go.uber.org/zap"
)

func RunAPI() {
	BaseSetup()
	handlers.SetupAndStart(config.AppConfig.Port)
}

func LoadConfig() {
	config.LoadConfig()
}

func SetupLogger() {
	logger.InitLogger(config.AppConfig.LogLevel)
}

func ConnectDB() {
	if err := database.Connect(config.AppConfig.DatabaseURL); err != nil {
		logger.Log.Fatal("Failed to connect to the database", zap.Error(err))
	}
}

func MigrateDB() {
	err := database.Migrate()
	if err != nil {
		logger.Log.Fatal("Failed to migrate database", zap.Error(err))
	}
}

func importUsers() {
	err := usersimport.ImportUsers("users_data.json")
	if err != nil {
		logger.Log.Fatal("Failed to import users", zap.Error(err))
	}
}

var once bool

func BaseSetup() {
	if !once {
		LoadConfig()
		SetupLogger()
		ConnectDB()
		MigrateDB()
		once = true
	}
}

func ImportUsers() {
	BaseSetup()
	importUsers()
}
