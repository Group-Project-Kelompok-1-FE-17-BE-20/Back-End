package data

import (
	"Laptop/features/shoppingcartitem"

	"gorm.io/gorm"
)

type itemQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) shoppingcartitem.ItemDataInterface {
	return &itemQuery{
		db: db,
	}
}

func (repo *itemQuery) Insert(input shoppingcartitem.Core) error {
	// simpan ke DB
	newItemGorm := CoreToModel(input)

	tx := repo.db.Create(&newItemGorm) // proses query insert
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}
