package routes

import (
	"github.com/labstack/echo/v5"
	"myapp/controllers"
)

func ProductRoutes(e *echo.Echo) {

	e.GET("/products", controllers.GetAllProducts)
	e.GET("/products/:id", controllers.GetProductById)

	e.POST("/products", controllers.CreateProduct)
	e.PUT("/products/:id", controllers.UpdateProduct)
	
	e.DELETE("/products/:id", controllers.DeleteProduct)

}