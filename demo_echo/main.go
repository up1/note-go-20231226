package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", hello)
	e.Logger.Fatal(e.Start(":1323"))
}

type Message struct {
	Message string `json:"message_123"`
}

func hello(c echo.Context) error {
	m := Message{"Hello, World!"}
	return c.JSON(http.StatusOK, m)
}
