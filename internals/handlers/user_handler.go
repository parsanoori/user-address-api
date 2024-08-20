package handlers

import (
	"UserAdresses/internals/controllers"
	"UserAdresses/internals/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

func createUser(e echo.Context) error {
	var user models.User
	err := e.Bind(&user)
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid request",
		})
	}
	err = controllers.AddUser(&user)
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid request",
		})
	}
	return e.JSON(http.StatusOK, map[string]string{
		"message": "success",
	})
}

func getUser(e echo.Context) error {
	id := e.Param("id")
	user, err := controllers.GetUser(id)
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid request",
		})
	}
	return e.JSON(http.StatusOK, user)
}

func updateUser(e echo.Context) error {
	var user models.User
	err := e.Bind(&user)
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid request, can't bind user data",
		})
	}
	err = controllers.UpdateUser(&user)
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid request, can't update user",
		})
	}
	return e.JSON(http.StatusOK, map[string]string{
		"message": "success",
	})
}

func deleteUser(e echo.Context) error {
	id := e.Param("id")
	err := controllers.DeleteUser(id)
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid request",
		})
	}
	return e.JSON(http.StatusOK, map[string]string{
		"message": "success",
	})
}
