package route

import (
	"backend/controller"

	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

func PaymentRoute(e *echo.Echo, db *gorm.DB) {
	paymentController := controller.NewPaymentController(db)
	e.POST("/payment", paymentController.CreatePayment)
}
