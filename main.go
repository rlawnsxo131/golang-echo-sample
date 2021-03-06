package main

import (
	"golang-echo-sample/api"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Debug = true
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	api.ApplyRoutes(e)
	e.Logger.Fatal(e.Start(":3001"))
}
