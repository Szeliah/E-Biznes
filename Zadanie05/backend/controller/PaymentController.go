package controller

import (
	"backend/model"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

type PaymentController struct {
	DB *gorm.DB
}

func NewPaymentController(db *gorm.DB) *PaymentController {
	return &PaymentController{DB: db}
}

func (ctrl *PaymentController) CreatePayment(ctx *echo.Context) error {

	db := ctrl.DB
	var validate = validator.New()

	var payment model.Payment

	if err := ctx.Bind(&payment); err != nil {
		return ctx.JSON(http.StatusBadRequest, "bad request")
	}

	if vErr := validate.Struct(&payment); vErr != nil {
		return ctx.JSON(http.StatusBadRequest, "error")
	}

	db.Create(&payment)
	return ctx.JSON(http.StatusCreated, &payment)
}
