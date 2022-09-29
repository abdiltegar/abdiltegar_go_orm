package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"learn_orm/constants"
	bookCtrl "learn_orm/controllers/bookController"
	usrCtrl "learn_orm/controllers/userController"
	mdlwr "learn_orm/middlewares"
)

func New() *echo.Echo {
	e := echo.New()

	e.POST("/login", usrCtrl.LoginController)

	jwtGroup := e.Group("")
	jwtGroup.Use(middleware.JWT([]byte(constants.SECRET_JWT)))

	// Route For Users
	usrRoute := e.Group("/users")
	usrRoute.GET("", usrCtrl.GetUsersController)
	usrRoute.GET("/:id", usrCtrl.GetUserController)
	usrRoute.POST("", usrCtrl.CreateUserController)
	usrRoute.DELETE("/:id", usrCtrl.DeleteUserController)
	usrRoute.PUT("/:id", usrCtrl.UpdateUserController)

	jwtGroup.GET("/users", usrCtrl.GetUsersController)
	jwtGroup.GET("/users/:id", usrCtrl.GetUserController)
	jwtGroup.DELETE("/users/:id", usrCtrl.DeleteUserController)
	jwtGroup.PUT("/users/:id", usrCtrl.UpdateUserController)

	// Route For Books
	bookRoute := e.Group("/books")
	bookRoute.GET("", bookCtrl.GetBooksController)
	bookRoute.GET("/:id", bookCtrl.GetBookController)
	bookRoute.POST("", bookCtrl.CreateBookController)
	bookRoute.DELETE("/:id", bookCtrl.DeleteBookController)
	bookRoute.PUT("/:id", bookCtrl.UpdateBookController)

	jwtGroup.POST("/books", bookCtrl.CreateBookController)
	jwtGroup.DELETE("/books/:id", bookCtrl.DeleteBookController)
	jwtGroup.PUT("/books/:id", bookCtrl.UpdateBookController)

	mdlwr.LogMiddleware(e)

	return e
}
