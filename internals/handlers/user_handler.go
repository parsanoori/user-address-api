package handlers

import (
	"UserAdresses/internals/controllers"
	"UserAdresses/internals/logger"
	"UserAdresses/internals/models"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"net/http"
)

func createUser(e echo.Context) error {
	var user models.User
	err := e.Bind(&user)
	if err != nil {
		logger.Log.Error("Failed to bind user data", zap.Error(err))
		return e.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid request",
		})
	}
	logger.Log.Info("User data binded successfully.", zap.Any("user", user))
	err = controllers.AddUser(&user)
	if err != nil {
		logger.Log.Error("Failed to add user", zap.Error(err))
		return e.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid request",
		})
	}
	logger.Log.Info("User added successfully on API endpoint.", zap.Any("user", user))
	return e.JSON(http.StatusOK, map[string]string{
		"message": "success",
	})
}

func getUser(e echo.Context) error {
	id := e.Param("id")
	user, err := controllers.GetUser(id)
	if err != nil {
		logger.Log.Error("Failed to get user", zap.Error(err))
		return e.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid request",
		})
	}
	logger.Log.Info("User fetched successfully.", zap.Any("user", user))
	return e.JSON(http.StatusOK, user)
}

func updateUser(e echo.Context) error {
	var user models.User
	err := e.Bind(&user)
	logger.Log.Info("User data binded successfully.", zap.Any("user", user))
	if err != nil {
		logger.Log.Error("Failed to bind user data", zap.Error(err))
		return e.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid request, can't bind user data",
		})
	}
	err = controllers.UpdateUser(&user)
	if err != nil {
		logger.Log.Error("Failed to update user", zap.Error(err))
		return e.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid request, can't update user",
		})
	}
	logger.Log.Info("User updated successfully on API endpoint.", zap.Any("user", user))
	return e.JSON(http.StatusOK, map[string]string{
		"message": "success",
	})
}

func deleteUser(e echo.Context) error {
	id := e.Param("id")
	err := controllers.DeleteUser(id)
	if err != nil {
		logger.Log.Error("Failed to delete user", zap.Error(err))
		return e.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid request",
		})
	}
	logger.Log.Info("User deleted successfully on API endpoint.", zap.String("id", id))
	return e.JSON(http.StatusOK, map[string]string{
		"message": "success",
	})
}
