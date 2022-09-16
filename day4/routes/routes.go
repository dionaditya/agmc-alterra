package routes

import (
	"agmc/day2/controllers"
	"agmc/day2/middleware"

	"github.com/labstack/echo/v4"
)

type Routing struct {
	user           controllers.UserController
	book           controllers.BookController
	authMiddleware middleware.AuthMiddleware
}

func NewRouting() Routing {
	return Routing{
		user: controllers.UserController{},
		book: controllers.BookController{},
	}
}

func (routing Routing) GetRouting(priv []byte, pub []byte) *echo.Echo {
	e := echo.New()

	e.Use(middleware.MiddlewareLogging)
	e.HTTPErrorHandler = middleware.ErrorHandler

	userController := controllers.NewUserController(priv, pub)

	v1 := e.Group("v1")

	v1.Use(routing.authMiddleware.JWTMiddlewareHandler(pub))

	v1.POST("/signin", userController.SignIn)

	v1.GET("/users", userController.GetAllUser)

	v1.POST("/users", userController.AddUser)

	v1.GET("/users/:id", userController.GetUserByID)

	v1.PUT("/users/:id", userController.UpdateUser)

	v1.DELETE("/users/:id", userController.DeleteUser)

	v1.GET("/books", routing.book.GetAllBooks)

	v1.POST("/books", routing.book.AddBook)

	v1.GET("/books/:id", routing.book.GetBookByID)

	v1.PUT("/books/:id", routing.book.UpdateBook)

	v1.DELETE("/books/:id", routing.book.DeleteBook)

	return e
}
