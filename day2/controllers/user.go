package controllers

import (
	"agmc/day2/entity"
	"agmc/day2/models"
	"agmc/day2/utils"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	userModel models.UserModel
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
	if err != nil {
		fmt.Printf("[UserController.AddUser] error bind data %v \n", err)
		return utils.HandleError(c, http.StatusInternalServerError, "Oppss server someting wrong")
	}
	if user.Email == "" || user.Name == "" {
		return utils.HandleError(c, http.StatusBadRequest, "field are required")
	}
	userData, err := userController.userModel.AddUser(&user)
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
