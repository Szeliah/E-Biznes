package controllers

import (
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
	"myapp/models"
)

func GetAllCarts(c *echo.Context) error {
	
	db := c.Get("db").(*gorm.DB)

	var carts []models.Cart
	db.Find(&carts)

	return c.JSON(http.StatusOK, &carts)
}

func GetCartById(c *echo.Context) error {
	
	db := c.Get("db").(*gorm.DB)

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "id must be a number")
	}

	var cart models.Cart

	db.First(&cart, id)

	if cart.ID == 0 {
		return c.JSON(http.StatusNotFound, "cart not found")
	}

	return c.JSON(http.StatusOK, &cart)
}



func CreateCart(c *echo.Context) error {

	db := c.Get("db").(*gorm.DB)
	var validate = validator.New()

	var cart models.Cart

	if err := c.Bind(&cart); err != nil {
		return c.JSON(http.StatusBadRequest, "bad request")
	}

	if vErr := validate.Struct(&cart); vErr != nil {
		return c.JSON(http.StatusBadRequest, "error")
	}

	db.Create(&cart)

	return c.JSON(http.StatusCreated, &cart)
}

func UpdateCart(c *echo.Context) error {

	db := c.Get("db").(*gorm.DB)
	var validate = validator.New()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "id must be a number")
	}

	var oldCart models.Cart
	var updatedCart models.Cart

	db.First(&oldCart, id)

	if oldCart.ID == 0 {
		return c.JSON(http.StatusNotFound, "cart not found")
	}

	if err := c.Bind(&updatedCart); err != nil {
		return c.JSON(http.StatusBadRequest, "bad request")
	}

	if vErr := validate.Struct(&updatedCart); vErr != nil {
		return c.JSON(http.StatusBadRequest, "error")
	} 

	updatedCart.Quantity = oldCart.Quantity
	updatedCart.Description = oldCart.Description
	db.Save(&updatedCart)

	return c.JSON(http.StatusOK, &updatedCart)
}


func DeleteCart(c *echo.Context) error {

	db := c.Get("db").(*gorm.DB)

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "id must be a number")
	}

	var cart models.Cart

	db.First(&cart, id)
	if cart.ID == 0 {
		return c.JSON(http.StatusNotFound, "cart not found")
	}

	db.Delete(&models.Cart{}, id)

	return c.JSON(http.StatusOK, "cart deleted")
}

