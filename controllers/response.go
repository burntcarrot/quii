package controllers

import "github.com/labstack/echo/v4"

type Response struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
}

func Success(c echo.Context, data interface{}) error {
	res := Response{
		Status: 200,
		Data:   data,
	}

	return c.JSON(res.Status, res)
}
