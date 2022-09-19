package routes

import (
	bookHandler "agmc/day2/internal/book/delivery/http"
	"agmc/day2/internal/factory"
	"agmc/day2/internal/middleware"
	userHandler "agmc/day2/internal/user/delivery/http"

	"github.com/labstack/echo/v4"
)

type Routing struct {
	f              *factory.Factory
	authMiddleware middleware.AuthMiddleware
}

func NewRouting(f *factory.Factory) Routing {
	return Routing{
		f: f,
	}
}

func InitUserRouting(g *echo.Group, u *userHandler.UserHandler) {
	g.POST("/signin", u.SignIn)

	g.GET("/users", u.GetAllUser)

	g.POST("/users", u.AddUser)

	g.GET("/users/:id", u.GetUserByID)

	g.PUT("/users/:id", u.UpdateUser)

	g.DELETE("/users/:id", u.DeleteUser)
}

func InitBookRouting(g *echo.Group, b *bookHandler.BookHandler) {
	g.GET("/books", b.GetAllBooks)

	g.POST("/books", b.AddBook)

	g.GET("/books/:id", b.GetBookByID)

	g.PUT("/books/:id", b.UpdateBook)

	g.DELETE("/books/:id", b.DeleteBook)
}

func (routing Routing) GetRouting(priv []byte, pub []byte) *echo.Echo {
	e := echo.New()

	e.Use(middleware.MiddlewareLogging)

	e.HTTPErrorHandler = middleware.ErrorHandler

	userHandler := userHandler.NewUserHandler(routing.f, priv, pub)

	bookhandler := bookHandler.NewBookHanndler(routing.f)

	v1 := e.Group("v1")

	v1.Use(routing.authMiddleware.JWTMiddlewareHandler(pub))

	InitUserRouting(v1, &userHandler)

	InitBookRouting(v1, &bookhandler)

	return e
}
