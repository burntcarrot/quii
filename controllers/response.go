package controllers

import "github.com/labstack/echo/v4"

type Response struct {
	Status       int         `json:"status"`
	Data         interface{} `json:"data"`
	ErrorMessage string      `json:"error"`
	Success      bool        `json:"success"`
}

func Success(c echo.Context, data interface{}) error {
	res := Response{
		Status:  200,
		Data:    data,
		Success: true,
	}

	return c.JSON(res.Status, res)
}

func Error(c echo.Context, status int, err error) error {
	res := Response{
		Status:       status,
		Data:         nil,
		Success:      false,
		ErrorMessage: err.Error(),
	}

	return c.JSON(res.Status, res)
}
