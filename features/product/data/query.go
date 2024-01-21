package data

import (
	"Laptop/features/product"

	"gorm.io/gorm"
)

type productQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) product.ProductDataInterface {
	return &productQuery{
		db: db,
	}
}

func (repo *productQuery) Insert(input product.Core) error {
	// simpan ke DB
	newProductGorm := CoreToModel(input)

	tx := repo.db.Create(&newProductGorm) // proses query insert
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (repo *productQuery) Update(idParam int, newUpdate product.Core) error {
	newUpdateGorm := CoreToModel(newUpdate)

	txUpdates := repo.db.Model(&Product{}).Where("id = ?", idParam).Updates(newUpdateGorm)
	if txUpdates.Error != nil {
		return txUpdates.Error
	}

	return nil
}
