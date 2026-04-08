package controllers

import (
	"net/http"
	"strconv"
	"sync"

	"github.com/labstack/echo/v5"
	"myapp/models"
)

type ProductDTO struct {
  Name  string `json:"name"`
  Price int `json:"price"`
}

var idx = 4
var lock = sync.Mutex{}

var products = map[int]*models.Product{
    1: {ID: 1, Name: "Młotek", Price: 32},
    2: {ID: 2, Name: "Kask", Price: 60},
    3: {ID: 3, Name: "Poziomica", Price: 45},
}


func GetAllProducts(c *echo.Context) error {
	lock.Lock()
	defer lock.Unlock()
	return c.JSON(http.StatusOK, products)
}

func GetProductById(c *echo.Context) error {
	lock.Lock()
	defer lock.Unlock()
	
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "id must be a number")
	}

	product, ok := products[id]
	if !ok {
		return c.JSON(http.StatusNotFound, "product not found")
	}

	return c.JSON(http.StatusOK, product)
}



func CreateProduct(c *echo.Context) error {
	lock.Lock()
	defer lock.Unlock()

	p := new(ProductDTO)
	if err := c.Bind(p); err != nil {
		return c.JSON(http.StatusBadRequest, "bad request")
	}

	product := &models.Product{
		ID: idx,
		Name: p.Name,
		Price: p.Price,
	}

	products[idx] = product
	idx++
	return c.JSON(http.StatusCreated, product)
}

func UpdateProduct(c *echo.Context) error {
	lock.Lock()
	defer lock.Unlock()

	p := new(ProductDTO)
	if err := c.Bind(p); err != nil {
		return c.JSON(http.StatusBadRequest, "bad request")
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "id must be a number")
	}

	product, ok := products[id]
	if !ok {
		return c.JSON(http.StatusNotFound, "product not found")
	}

	product.Name = p.Name
	product.Price = p.Price

	return c.JSON(http.StatusOK, product)
}



func DeleteProduct(c *echo.Context) error {
	lock.Lock()
	defer lock.Unlock()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "id must be a number")
	}

	_, ok := products[id]
	if !ok {
		return c.JSON(http.StatusNotFound, "product not found")
	}

	delete(products, id)

	return c.JSON(http.StatusOK, "product deleted ")
}

