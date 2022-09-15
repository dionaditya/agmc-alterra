package routes

import (
	"agmc/day2/controllers"
	"agmc/day2/middleware"

	"github.com/labstack/echo/v4"
)

type Routing struct {
	user controllers.UserController
	book controllers.BookController
}

func (routing Routing) GetRouting() *echo.Echo {
	e := echo.New()

	e.Use(middleware.MiddlewareLogging)
	e.HTTPErrorHandler = middleware.ErrorHandler

	v1 := e.Group("v1")

	v1.Use(middleware.JWTMiddlewareHandler())

	v1.POST("/signin", routing.user.SignIn)

	v1.GET("/users", routing.user.GetAllUser)

	v1.POST("/users", routing.user.AddUser)

	v1.GET("/users/:id", routing.user.GetUserByID)

	v1.PUT("/users/:id", routing.user.UpdateUser)

	v1.DELETE("/users/:id", routing.user.DeleteUser)

	v1.GET("/books", routing.book.GetAllBooks)

	v1.POST("/books", routing.book.AddBook)

	v1.GET("/books/:id", routing.book.GetBookByID)

	v1.PUT("/books/:id", routing.book.UpdateBook)

	v1.DELETE("/books/:id", routing.book.DeleteBook)

	return e
}
