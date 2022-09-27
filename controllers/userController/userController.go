package userController

import (
	"github.com/labstack/echo/v4"
	"learn_orm/config"
	"learn_orm/models"
	"net/http"
	"strconv"
)

// get all users
func GetUsersController(c echo.Context) error {
	var users []models.User
	if err := config.DB.Find(&users).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all users",
		"users":   users,
	})
}

// get user by id
func GetUserController(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 16, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Id must be a number",
			"users":   models.User{},
		})
	}

	var user models.User
	if err := config.DB.First(&user, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if user.ID == 0 {
		return c.JSON(http.StatusNoContent, map[string]interface{}{
			"message": "Data not found",
			"user":    user,
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get user",
		"users":   user,
	})
}

// create user
func CreateUserController(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)

	if err := config.DB.Save(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new user",
		"user":    user,
	})
}

// delete user by id
func DeleteUserController(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 16, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Id must be a number",
			"users":   models.User{},
		})
	}

	var user models.User
	if err := config.DB.Delete(&user, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	var users []models.User
	if err := config.DB.Find(&users).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete user",
		"users":   users,
	})
}

// update user by id
func UpdateUserController(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 16, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Id must be a number",
			"users":   models.User{},
		})
	}

	var user models.User
	if err := config.DB.First(&user, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	userParam := models.User{}
	c.Bind(&userParam)
	userParam.ID = uint(id)
	user = userParam

	if err := config.DB.Save(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update user",
		"users":   user,
	})
}
