package main

import (
	"backend/database"
	"backend/route"
	"log"
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.RequestLogger())
	e.Use(middleware.CORS("http://127.0.0.1:5173"))

	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}

	if err := database.Migrate(db); err != nil {
		log.Fatal(err)
	}

	if err := database.SeedProducts(db); err != nil {
		log.Fatal(err)
	}

	e.GET("/", func(c *echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	route.ProductRoute(e, db)
	route.PaymentRoute(e, db)

	if err := e.Start(":8080"); err != nil {
		e.Logger.Error("failed to start server", "error", err)
	}
}
