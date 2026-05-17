package model

import "gorm.io/gorm"

type CartItem struct {
	gorm.Model
	CartID    uint `json:"cart_id"`
	ProductID uint `json:"product_id"`

	Quantity int     `json:"quantity"`
	Price    float64 `json:"price"`

	// Product Product `gorm:"foreignKey:ProductID"`
}
