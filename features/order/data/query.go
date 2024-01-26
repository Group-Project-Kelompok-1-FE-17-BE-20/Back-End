package data

import (
	"Laptop/app/database"
	"Laptop/features/order"
	"Laptop/features/shoppingcartitem"

	"gorm.io/gorm"
)

type orderQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) order.OrderDataInterface {
	return &orderQuery{
		db: db,
	}
}

func (repo *orderQuery) GetCartID(userID uint) (uint, error) {
	var cartData database.ShoppingCart
	tx := repo.db.Where("user_id = ?", userID).First(&cartData)
	if tx.Error != nil {
		return 0, tx.Error
	}

	cartID := cartData.ID
	return cartID, nil
}

func (repo *orderQuery) GetAllCartItem(cartID uint) ([]shoppingcartitem.Core, error) {
	var allItemsData []database.ShoppingCartItem
	tx := repo.db.Where("shopping_cart_id = ?", cartID).Find(&allItemsData)
	if tx.Error != nil {
		return nil, tx.Error
	}

	itemsData := GetItemsData(allItemsData)
	itemsDataCore := GetItemsDataCore(itemsData)

	return itemsDataCore, nil
}

// Insert implements order.OrderDataInterface.
func (repo *orderQuery) Insert(input order.Core) error {
	newOrderGorm := CoreToModel(input)

	tx := repo.db.Create(&newOrderGorm) // proses query insert
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}
