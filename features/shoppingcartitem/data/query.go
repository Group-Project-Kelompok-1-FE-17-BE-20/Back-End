package data

import (
	"Laptop/app/database"
	"Laptop/features/shoppingcartitem"
	"errors"

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

func (repo *itemQuery) GetCartID(userID uint) (uint, error) {
	var cartData database.ShoppingCart
	tx := repo.db.Where("user_id = ?", userID).First(&cartData)
	if tx.Error != nil {
		return 0, tx.Error
	}

	cartID := cartData.ID
	return cartID, nil
}

func (repo *itemQuery) GetPrice(productID uint) (float64, error) {
	var productData database.Product
	tx := repo.db.Where("ID = ?", productID).First(&productData)
	if tx.Error != nil {
		return 0, tx.Error
	}

	productPrice := productData.Price
	return productPrice, nil
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

func (repo *itemQuery) Update(productID uint, input shoppingcartitem.Core) error {
	newUpdateGorm := CoreToModel(input)

	txUpdates := repo.db.Model(&database.ShoppingCartItem{}).Where("product_id = ?", productID).Updates(newUpdateGorm)
	if txUpdates.Error != nil {
		return txUpdates.Error
	}

	return nil
}

func (repo *itemQuery) GetItemById(productID uint) (shoppingcartitem.Core, error) {
	var singleItemGorm database.ShoppingCartItem
	tx := repo.db.Where("product_id = ?", productID).First(&singleItemGorm)
	if tx.Error != nil {
		return shoppingcartitem.Core{}, tx.Error
	}

	singleItemCore := ModelToCore(singleItemGorm)

	return singleItemCore, nil
}

func (repo *itemQuery) Delete(input shoppingcartitem.Core) error {
	itemGorm := CoreToModel(input)

	txDel := repo.db.Where("product_id = ?", itemGorm.ProductID).Delete(&itemGorm)
	if txDel.Error != nil {
		return txDel.Error
	}

	if txDel.RowsAffected == 0 {
		return errors.New("item not found")
	}

	return nil
}
