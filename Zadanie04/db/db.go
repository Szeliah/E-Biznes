package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"myapp/models"
)


func Connect() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("DATABASE.db"))

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.Product{})
	db.AutoMigrate(&models.Cart{})

	return db
}
