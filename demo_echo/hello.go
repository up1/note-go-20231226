package demo

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Message struct {
	Message string `json:"message_123"`
}

func Hello(c echo.Context) error {
	repo := NewHelloRepo()
	result, _ := repo.GetData()
	m := Message{
		Message: result,
	}
	return c.JSON(http.StatusOK, m)
}

func TryToFail(c echo.Context) error {
	panic("I'm panicking!")
}
