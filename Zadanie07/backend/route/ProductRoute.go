package route

import (
	"backend/controller"

	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

func ProductRoute(e *echo.Echo, db *gorm.DB) {
	productController := controller.NewProductController(db)
	e.GET("/products", productController.GetProducts)
}
