package handler

import (
	domains "agmc/day2/internal/domains"
	"agmc/day2/internal/factory"
	"agmc/day2/internal/user/usecase"
	"agmc/day2/pkg/constant"
	"agmc/day2/pkg/utils/response"
	"fmt"
	"net/http"

	"agmc/day2/internal/middleware"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userUseCase domains.UserUseCase
	privateKey  []byte
	pubKey      []byte
}

func NewUserHandler(f *factory.Factory, privateKey []byte, pubKey []byte) UserHandler {
	return UserHandler{
		privateKey:  privateKey,
		userUseCase: usecase.NewUserUseCase(f.UserRepo),
		pubKey:      pubKey,
	}
}

func (userHandler *UserHandler) GetAllUser(c echo.Context) error {
	users, err := userHandler.userUseCase.GetAllUser()
	if err != nil {
		return response.HandleError(c, http.StatusInternalServerError, err.Error())
	}

	return response.HandleSuccess(c, users)
}

func (userHandler *UserHandler) GetUserByID(c echo.Context) error {
	id := c.Param("id")
	users, err := userHandler.userUseCase.GetUserByID(id)
	if err != nil {
		if fmt.Sprint(err) == constant.E_NOT_FOUND {
			return response.HandleError(c, http.StatusNotFound, err.Error())
		}
		return response.HandleError(c, http.StatusInternalServerError, err.Error())
	}

	return response.HandleSuccess(c, users)
}

func (userHandler *UserHandler) AddUser(c echo.Context) error {
	var user = domains.User{}
	err := c.Bind(&user)

	if err != nil {
		fmt.Printf("[UserController.AddUser] error bind data %v \n", err)
		return response.HandleError(c, http.StatusInternalServerError, constant.E_UNPROCESSABLE_ENTITY)
	}

	if user.Email == "" || user.Name == "" {
		return response.HandleError(c, http.StatusBadRequest, constant.E_BAD_REQUEST)
	}

	userData, err := userHandler.userUseCase.AddUser(&user)

	if err != nil {
		return response.HandleError(c, http.StatusInternalServerError, err.Error())
	}
	return response.HandleSuccess(c, userData)
}

func (userHandler *UserHandler) UpdateUser(c echo.Context) error {
	var user = domains.User{}
	id := c.Param("id")

	err := c.Bind(&user)
	if err != nil {
		fmt.Printf("[UserController.UpdateUser] error bind data %v \n", err)
		return response.HandleError(c, http.StatusInternalServerError, constant.E_UNPROCESSABLE_ENTITY)
	}
	if user.Email == "" || user.Name == "" || user.Password == "" {
		return response.HandleError(c, http.StatusBadRequest, constant.E_BAD_REQUEST)
	}
	mUser, err := userHandler.userUseCase.UpdateUser(id, &user)
	if err != nil {
		return response.HandleError(c, http.StatusInternalServerError, err.Error())
	}
	return response.HandleSuccess(c, mUser)
}

func (userHandler *UserHandler) DeleteUser(c echo.Context) error {
	id := c.Param("id")

	err := userHandler.userUseCase.DeleteUser(id)

	if err != nil {
		return response.HandleError(c, http.StatusInternalServerError, err.Error())
	}
	return response.HandleSuccess(c, struct{}{})
}

func (userHandler *UserHandler) SignIn(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")

	token, err := userHandler.userUseCase.SignIn(email, password, userHandler.privateKey)

	if err == nil {
		return c.JSON(http.StatusOK, middleware.Token{
			Token: token,
		})
	}

	return echo.ErrUnauthorized
}
