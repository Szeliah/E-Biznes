package routes

import (
	"github.com/labstack/echo/v5"
	"myapp/controllers"
)

func CategoryRoutes(e *echo.Echo) {
	
	e.GET("/categories", controllers.GetAllCategories)
	e.GET("/categories/:id", controllers.GetCategoryById)
	
	e.POST("/categories", controllers.CreateCategory)
	e.PUT("/categories/:id", controllers.UpdateCategory)
	
	e.DELETE("/categories/:id", controllers.DeleteCategory)

}