package data

import (
	"Laptop/app/database"
	"Laptop/features/shoppingcart"
	"errors"

	"gorm.io/gorm"
)

type CartQuery struct {
	db *gorm.DB
}

func New(database *gorm.DB) shoppingcart.CartDataInterface {
	return &CartQuery{
		db: database,
	}
}

func (repo *CartQuery) Insert(input shoppingcart.CoreCart) error {
	// simpan ke DB
	newProductGorm := CoreCartToGorm(input)

	tx := repo.db.Create(&newProductGorm) // proses query insert
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

// Select implements task.TaskDataInterface.
func (r *CartQuery) SelectCart(userID uint) (shoppingcart.CoreCart, error) {
	var cartData database.ShoppingCart
	tx := r.db.Where("user_id = ?", userID).First(&cartData)
	if tx.Error != nil {
		return shoppingcart.CoreCart{}, tx.Error
	}

	if tx.RowsAffected == 0 {
		return shoppingcart.CoreCart{}, errors.New("cart not found")
	}

	coreCart := CartGormToCartCore(cartData)
	return coreCart, nil
}
