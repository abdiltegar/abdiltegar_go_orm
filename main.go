package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

var (
	DB *gorm.DB
)

func init() {
	InitDB()
	InitialMigration()
}

type Config struct {
	DB_Username string
	DB_Password string
	DB_Port     string
	DB_Host     string
	DB_Name     string
}

func InitDB() {
	config := Config{
		DB_Username: "root",
		DB_Password: "Bismillah12",
		DB_Port:     "3306",
		DB_Host:     "localhost",
		DB_Name:     "crud_go",
	}

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.DB_Username,
		config.DB_Password,
		config.DB_Host,
		config.DB_Port,
		config.DB_Name,
	)

	var err error
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}

type User struct {
	gorm.Model
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

func InitialMigration() {
	DB.AutoMigrate(&User{})
}

// get all users
func GetUsersController(c echo.Context) error {
	var users []User
	if err := DB.Find(&users).Error; err != nil {
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
			"users":   User{},
		})
	}

	var user User
	if err := DB.First(&user, id).Error; err != nil {
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

func CreateUserController(c echo.Context) error {
	user := User{}
	c.Bind(&user)

	if err := DB.Save(&user).Error; err != nil {
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
			"users":   User{},
		})
	}

	var users []User
	if err := DB.Find(&users).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	var user User
	if err := DB.Delete(&user, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get user",
		"users":   users,
	})
}

// update user by id
func UpdateUserController(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 16, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Id must be a number",
			"users":   User{},
		})
	}

	var user User
	if err := DB.First(&user, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	userParam := User{}
	c.Bind(&userParam)
	userParam.ID = uint(id)

	if err := DB.Model(&user).Updates(userParam).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get user",
		"users":   user,
	})
}

func main() {
	e := echo.New()

	e.GET("/users", GetUsersController)
	e.GET("/users/:id", GetUserController)
	e.POST("/users", CreateUserController)
	e.DELETE("/users/:id", DeleteUserController)
	e.PUT("/users/:id", UpdateUserController)

	e.Logger.Fatal(e.Start(":8083"))
}
