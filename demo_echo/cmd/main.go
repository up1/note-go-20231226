package main

import (
	"demo"

	"github.com/labstack/echo-contrib/echoprometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(echoprometheus.NewMiddleware("demo_service"))
	e.GET("/metrics", echoprometheus.NewHandler())
	e.Use(middleware.Logger())

	e.GET("/", demo.Hello)
	e.GET("/panic", demo.TryToFail)
	e.Logger.Fatal(e.Start(":1323"))
}
