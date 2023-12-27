package main

import (
	"demo"

	"github.com/labstack/echo-contrib/echoprometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"net/http"
	"net/http/pprof"
)

func main() {
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(echoprometheus.NewMiddleware("demo_service"))
	e.GET("/metrics", echoprometheus.NewHandler())
	e.Use(middleware.Logger())

	e.GET("/debug/pprof/*", echo.WrapHandler(http.HandlerFunc(pprof.Index)))

	// Create dependency
	// db, _ := demo.CreateConnection()
	// demo.NewHelloRepo(db)

	demo.MyRepo = &MockHelloRepo{}

	repo := MockHelloRepo{}
	e.GET("/", demo.Hello(&repo))

	env := demo.Env{
		Repo: &repo,
	}
	e.GET("/hello2", env.Hello2)

	e.GET("/panic", demo.TryToFail)
	e.Logger.Fatal(e.Start(":1323"))
}

type MockHelloRepo struct{}

func (r *MockHelloRepo) GetData() (string, error) {
	return "Hello World", nil
}
