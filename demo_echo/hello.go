package demo

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Message struct {
	Message string `json:"message_123"`
}

func Hello(repo Repository) func(c echo.Context) error {
	return func(c echo.Context) error {
		result, _ := repo.GetData()
		m := Message{
			Message: result,
		}
		return c.JSON(http.StatusOK, m)
	}
}

func TryToFail(c echo.Context) error {
	panic("I'm panicking!")
}
