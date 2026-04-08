package main

import (
	"github.com/labstack/echo/v5/middleware"
	"github.com/labstack/echo/v5"
	"myapp/routes"

	"net/http"
)

func main() {
	e := echo.New()

	routes.RegisterRoutes(e)

	e.Use(middleware.RequestLogger())

	e.GET("/", func(c *echo.Context) error {
		return c.String(http.StatusOK, "Hello, World !")
	})

	if err := e.Start(":8080"); err != nil {
		e.Logger.Error("failed to start server", "error", err)
	}
}
