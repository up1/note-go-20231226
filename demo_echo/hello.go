package demo

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Message struct {
	Message string `json:"message_123"`
}

func Hello(c echo.Context) error {
	m := Message{"Hello, World!"}
	return c.JSON(http.StatusOK, m)
}
