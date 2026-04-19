package database

import (
	"backend/model"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("Database.db"), &gorm.Config{})

	if err != nil {
		return db, err
	}

	return db, nil
}

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&model.Product{},
		&model.Cart{},
		&model.CartItem{},
		&model.Payment{},
	)
}
