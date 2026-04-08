package routes

import (
	"github.com/labstack/echo/v5"
	"myapp/controllers"
)

func CartRoutes(e *echo.Echo) {
	
	e.GET("/carts", controllers.GetAllCarts)
	e.GET("/carts/:id", controllers.GetCartById)
	
	e.POST("/carts", controllers.CreateCart)
	e.PUT("/carts/:id", controllers.UpdateCart)
	
	e.DELETE("/carts/:id", controllers.DeleteCart)

}