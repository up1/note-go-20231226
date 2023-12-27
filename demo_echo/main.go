package main

import (
	"net/http"

	"github.com/labstack/echo-contrib/echoprometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(echoprometheus.NewMiddleware("demo_service"))
	e.GET("/metrics", echoprometheus.NewHandler())

	e.GET("/", hello)
	e.GET("/panic", tryToFail)
	e.Logger.Fatal(e.Start(":1323"))
}

func tryToFail(c echo.Context) error {
	panic("I'm panicking!")
}

type Message struct {
	Message string `json:"message_123"`
}

func hello(c echo.Context) error {
	m := Message{"Hello, World!"}
	return c.JSON(http.StatusOK, m)
}
