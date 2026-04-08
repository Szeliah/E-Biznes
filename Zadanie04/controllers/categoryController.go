package controllers

import (
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
	"myapp/models"
)

func GetAllCategories(c *echo.Context) error {
	
	db := c.Get("db").(*gorm.DB)

	var categories []models.Category
	db.Preload("Products").Find(&categories)

	return c.JSON(http.StatusOK, &categories)
}

func GetCategoryById(c *echo.Context) error {
	
	db := c.Get("db").(*gorm.DB)

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "id must be a number")
	}

	var category models.Category

	db.Preload("Products").First(&category, id)

	if category.ID == 0 {
		return c.JSON(http.StatusNotFound, "category not found")
	}

	return c.JSON(http.StatusOK, &category)
}



func CreateCategory(c *echo.Context) error {

	db := c.Get("db").(*gorm.DB)
	var validate = validator.New()

	var category models.Category

	if err := c.Bind(&category); err != nil {
		return c.JSON(http.StatusBadRequest, "bad request")
	}

	if vErr := validate.Struct(&category); vErr != nil {
		return c.JSON(http.StatusBadRequest, "error")
	}

	db.Create(&category)

	return c.JSON(http.StatusCreated, &category)
}

func UpdateCategory(c *echo.Context) error {

	db := c.Get("db").(*gorm.DB)
	var validate = validator.New()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "id must be a number")
	}

	var oldCategory models.Category
	var updatedCategory models.Category

	db.First(&oldCategory, id)

	if oldCategory.ID == 0 {
		return c.JSON(http.StatusNotFound, "category not found")
	}

	if err := c.Bind(&updatedCategory); err != nil {
		return c.JSON(http.StatusBadRequest, "bad request")
	}

	if vErr := validate.Struct(&updatedCategory); vErr != nil {
		return c.JSON(http.StatusBadRequest, "error")
	} 

	updatedCategory.Name = oldCategory.Name
	updatedCategory.Description  = oldCategory.Description 
	db.Save(&updatedCategory)

	return c.JSON(http.StatusOK, &updatedCategory)
}


func DeleteCategory(c *echo.Context) error {

	db := c.Get("db").(*gorm.DB)

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "id must be a number")
	}

	var category models.Category

	db.First(&category, id)
	if category.ID == 0 {
		return c.JSON(http.StatusNotFound, "category not found")
	}

	var productsCnt int64
	db.Model(&models.Product{}).Where("category_id = ?", category.ID).Count(&productsCnt)

	if productsCnt > 0 {
		return c.JSON(http.StatusConflict, "cannot delete category with existing products")
	}

	db.Delete(&models.Category{}, id)

	return c.JSON(http.StatusOK, "category deleted")
}

