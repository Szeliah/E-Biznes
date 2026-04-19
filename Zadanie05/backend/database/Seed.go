package database

import (
	"backend/model"

	"gorm.io/gorm"
)

func SeedProducts(db *gorm.DB) error {
	var count int64
	db.Model(&model.Product{}).Count(&count)

	if count > 0 {
		return nil
	}

	products := []model.Product{
		{
			Name:        "Niezwykły Młotek",
			Description: "Młotek do wbijania gwoździ, ale i nie tylko",
			Price:       44.99,
		},
		{
			Name:        "Super Poziomica",
			Description: "Narzędzie służące do sprawdzania pionów i poziomów ale, także do sprawdzania czy jest równa powierzchnia",
			Price:       99.99,
		},
		{
			Name:        "Niesamowity Łom",
			Description: "Narzędzie służące jako 'dźiwgnia' do wyważania, podważania czy rozbijania czegoś",
			Price:       76.99,
		},
	}

	return db.Create(&products).Error
}
