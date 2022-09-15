package utils

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type ResponseGeneric struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func HandleSuccess(c echo.Context, data interface{}) error {
	res := ResponseGeneric{
		Status:  "Success",
		Message: "Posts Loaded",
		Data:    data,
	}
	return c.JSON(http.StatusOK, res)
}

func HandleError(c echo.Context, status int, message string) error {
	res := ResponseGeneric{
		Status:  "Failed",
		Message: message,
	}
	return c.JSON(status, res)
}
