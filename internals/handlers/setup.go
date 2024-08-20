package handlers

import "github.com/labstack/echo/v4"

var e *echo.Echo

func Setup() *echo.Echo {
	e = echo.New()
	registerRoutes()
	return e
}

func Start(port string) error {
	return e.Start(":" + port)
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
}
