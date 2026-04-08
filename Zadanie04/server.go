package main

import (
	"github.com/labstack/echo/v5/middleware"
	"github.com/labstack/echo/v5"
	
	"myapp/routes"
	"myapp/db"
)

func main() {
	e := echo.New()

	db := db.Connect()

	e.Use(middleware.RequestLogger())
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c *echo.Context) error {
			c.Set("db", db)
			return next(c)
		}
	})
	
	routes.RegisterRoutes(e)

	if err := e.Start(":8080"); err != nil {
		e.Logger.Error("failed to start server", "error", err)
	}
}
