package handlers

import (
	"net/http"

	"github.com/ferjoaguilar/rest-echo-repository/server"
	"github.com/labstack/echo/v4"
)

type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func HealhCkeckHandler(s server.Sever) echo.HandlerFunc {
	return func(c echo.Context) error {
		response := Response{
			Message: "Hello to Echo api",
			Data:    map[string]any{"name": "fernando"},
		}

		return c.JSON(http.StatusOK, response)
	}
}
