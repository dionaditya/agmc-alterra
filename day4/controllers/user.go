package controllers

import (
	"agmc/day2/entity"
	"agmc/day2/models"
	"agmc/day2/utils"
	"fmt"
	"net/http"

	"agmc/day2/middleware"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type UserController struct {
	userModel      models.UserModel
	privateKey     []byte
	pubKey         []byte
	authMiddleware middleware.AuthMiddleware
}

func NewUserController(privateKey []byte, pubKey []byte) UserController {
	return UserController{
		privateKey: privateKey,
		pubKey:     pubKey,
	}
}

func (userController *UserController) GetAllUser(c echo.Context) error {
	users, err := userController.userModel.GetAllUser()
	if err != nil {
		return utils.HandleError(c, http.StatusInternalServerError, err.Error())
	}

	return utils.HandleSuccess(c, users)
}

func (userController *UserController) GetUserByID(c echo.Context) error {
	id := c.Param("id")
	users, err := userController.userModel.GetUserByID(id)
	if err != nil {
		if fmt.Sprint(err) == "No content found" {
			return utils.HandleError(c, http.StatusNotFound, err.Error())
		}
		return utils.HandleError(c, http.StatusInternalServerError, err.Error())
	}

	return utils.HandleSuccess(c, users)
}

func (userController *UserController) AddUser(c echo.Context) error {
	var user = entity.User{}
	err := c.Bind(&user)

	fmt.Println(err)
	if err != nil {
		fmt.Printf("[UserController.AddUser] error bind data %v \n", err)
		return utils.HandleError(c, http.StatusInternalServerError, "Oppss server someting wrong")
	}
	if user.Email == "" || user.Name == "" {
		return utils.HandleError(c, http.StatusBadRequest, "field are required")
	}
	userData, err := userController.userModel.AddUser(&user)

	fmt.Println(err)

	if err != nil {
		return utils.HandleError(c, http.StatusInternalServerError, err.Error())
	}
	return utils.HandleSuccess(c, userData)
}

func (userController *UserController) UpdateUser(c echo.Context) error {
	var user = entity.User{}
	id := c.Param("id")

	err := c.Bind(&user)
	if err != nil {
		fmt.Printf("[UserController.UpdateUser] error bind data %v \n", err)
		return utils.HandleError(c, http.StatusInternalServerError, "Oppss server someting wrong")
	}
	if user.Email == "" || user.Name == "" || user.Password == "" {
		return utils.HandleError(c, http.StatusBadRequest, "field are required")
	}
	mUser, err := userController.userModel.UpdateUser(id, &user)
	if err != nil {
		return utils.HandleError(c, http.StatusInternalServerError, err.Error())
	}
	return utils.HandleSuccess(c, mUser)
}

func (userController *UserController) DeleteUser(c echo.Context) error {
	id := c.Param("id")

	err := userController.userModel.DeleteUser(id)

	if err != nil {
		return utils.HandleError(c, http.StatusInternalServerError, err.Error())
	}
	return utils.HandleSuccess(c, struct{}{})
}

func (userController *UserController) SignIn(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")

	var user *entity.User

	err := userController.userModel.GetUserByEmail(email).First(&user).Error

	if err != nil {
		return utils.HandleError(c, http.StatusInternalServerError, err.Error())
	}

	if err != nil {
		return echo.ErrUnauthorized
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if email == user.Email && err == nil {
		token, err := userController.authMiddleware.GenerateToken(userController.privateKey)

		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, middleware.Token{
			Token: token,
		})
	}

	return echo.ErrUnauthorized
}
