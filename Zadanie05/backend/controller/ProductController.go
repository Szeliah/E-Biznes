package controller

import (
	"backend/model"
	"net/http"

	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

type ProductResponse struct {
	ID          uint    `json:"id"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
}

type ProductController struct {
	DB *gorm.DB
}

func ToProductResponse(p model.Product) ProductResponse {
	return ProductResponse{
		ID:          p.ID,
		Name:        p.Name,
		Price:       p.Price,
		Description: p.Description,
	}
}

func NewProductController(db *gorm.DB) *ProductController {
	return &ProductController{DB: db}
}

func (ctrl *ProductController) GetProducts(ctx *echo.Context) error {
	var products []model.Product
	ctrl.DB.Find(&products)

	var res []ProductResponse
	for _, p := range products {
		res = append(res, ToProductResponse(p))
	}

	return ctx.JSON(http.StatusOK, &res)
}
