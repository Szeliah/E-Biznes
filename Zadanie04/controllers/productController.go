package controllers

import (
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
	"myapp/models"
)

var validate = validator.New()

func GetAllProducts(c *echo.Context) error {
	
	db := c.Get("db").(*gorm.DB)
	
	var products []models.Product
	db.Find(&products)

	return c.JSON(http.StatusOK, &products)
}

func GetProductById(c *echo.Context) error {
	
	db := c.Get("db").(*gorm.DB)

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "id must be a number")
	}

	var product models.Product

	db.First(&product, id)

	if product.ID == 0 {
		return c.JSON(http.StatusNotFound, "product not found")
	}

	return c.JSON(http.StatusOK, &product)
}



func CreateProduct(c *echo.Context) error {

	db := c.Get("db").(*gorm.DB)

	var product models.Product

	if err := c.Bind(&product); err != nil {
		return c.JSON(http.StatusBadRequest, "bad request")
	}

	if vErr := validate.Struct(&product); vErr != nil {
		return c.JSON(http.StatusBadRequest, "error")
	}

	db.Create(&product)

	return c.JSON(http.StatusCreated, &product)
}

func UpdateProduct(c *echo.Context) error {

	db := c.Get("db").(*gorm.DB)

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "id must be a number")
	}

	var oldProduct models.Product
	var updatedProduct models.Product

	db.First(&oldProduct, id)

	if oldProduct.ID == 0 {
		return c.JSON(http.StatusNotFound, "product not found")
	}

	if err := c.Bind(&updatedProduct); err != nil {
		return c.JSON(http.StatusBadRequest, "bad request")
	}

	if vErr := validate.Struct(&updatedProduct); vErr != nil {
		return c.JSON(http.StatusBadRequest, "error")
	} 

	db.Save(&updatedProduct)

	return c.JSON(http.StatusOK, &updatedProduct)
}


func DeleteProduct(c *echo.Context) error {

	db := c.Get("db").(*gorm.DB)

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "id must be a number")
	}

	var product models.Product

	db.First(&product, id)
	if product.ID == 0 {
		return c.JSON(http.StatusNotFound, "product not found")
	}

	db.Delete(&models.Product{}, id)

	return c.JSON(http.StatusOK, "product deleted")
}

