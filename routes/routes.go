package routes

import (
	"github.com/labstack/echo/v4"
	bookCtrl "learn_orm/controllers/bookController"
	usrCtrl "learn_orm/controllers/userController"
)

func New() *echo.Echo {
	e := echo.New()

	usrRoute := e.Group("/users")
	usrRoute.GET("", usrCtrl.GetUsersController)
	usrRoute.GET("/:id", usrCtrl.GetUserController)
	usrRoute.POST("", usrCtrl.CreateUserController)
	usrRoute.DELETE("/:id", usrCtrl.DeleteUserController)
	usrRoute.PUT("/:id", usrCtrl.UpdateUserController)

	bookRoute := e.Group("/books")
	bookRoute.GET("", bookCtrl.GetBooksController)
	bookRoute.GET("/:id", bookCtrl.GetBookController)
	bookRoute.POST("", bookCtrl.CreateBookController)
	bookRoute.DELETE("/:id", bookCtrl.DeleteBookController)
	bookRoute.PUT("/:id", bookCtrl.UpdateBookController)

	return e
}
