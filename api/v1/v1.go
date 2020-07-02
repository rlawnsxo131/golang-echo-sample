package v1

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Hello struct {
	Hello string `json:"hello"`
	World string `json:"world"`
}

func handler(c echo.Context) error {
	hello := &Hello{
		Hello: "hello",
		World: "world",
	}
	return c.JSON(http.StatusOK, hello)
}

func ApplyRoutes(e *echo.Group) {
	v1 := e.Group("/v1")
	v1.GET("/hello", handler)
}
