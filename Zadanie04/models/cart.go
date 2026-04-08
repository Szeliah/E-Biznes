package models

import "gorm.io/gorm"

type Cart struct {
    gorm.Model
    Quantity int `json:"quantity"`
    Description  string    `json:"description"`
}