package handlers

import (
	"UserAdresses/internals/logger"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

var e *echo.Echo

func Setup() *echo.Echo {
	e = echo.New()
	registerRoutes()
	return e
}

func Start(port string) error {
	err := e.Start(":" + port)
	if err != nil {
		logger.Log.Fatal("Failed to start the server", zap.Error(err))
	}
	logger.Log.Info("Server started successfully.")
	return err
}

func SetupAndStart(port string) error {
	Setup()
	return Start(port)
}

func registerRoutes() {
	e.POST("/user", createUser)
	e.GET("/user/:id", getUser)
	e.PUT("/user", updateUser)
	e.DELETE("/user/:id", deleteUser)
	logger.Log.Info("Routes registered successfully.")
}
