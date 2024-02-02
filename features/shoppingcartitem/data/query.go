package data

import (
	"Laptop/app/database"
	"Laptop/features/product"
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

func (repo *itemQuery) InsertCart(input shoppingcartitem.CoreCart) error {
	// simpan ke DB
	newProductGorm := CoreCartToGorm(input)

	tx := repo.db.Create(&newProductGorm) // proses query insert
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

// Select implements task.TaskDataInterface.
func (repo *itemQuery) SelectCart(userID uint, status string) (shoppingcartitem.CoreCart, error) {
	var cartData database.ShoppingCart
	tx := repo.db.Where("user_id = ? and status = ?", userID, status).First(&cartData)
	if tx.Error != nil {
		return shoppingcartitem.CoreCart{}, tx.Error
	}

	if tx.RowsAffected == 0 {
		return shoppingcartitem.CoreCart{}, errors.New("cart not found")
	}

	coreCart := CartGormToCartCore(cartData)
	return coreCart, nil
}

func (repo *itemQuery) GetCartID(userID uint) (uint, error) {
	var cartData database.ShoppingCart
	tx := repo.db.Where("user_id = ? and status = 'On Going'", userID).First(&cartData)
	if tx.Error != nil {
		return 0, tx.Error
	}

	cartID := cartData.ID
	return cartID, nil
}

func (repo *itemQuery) GetDataProduct(productID uint) (product.Core, error) {
	var productData database.Product
	tx := repo.db.Where("ID = ?", productID).First(&productData)
	if tx.Error != nil {
		return product.Core{}, tx.Error
	}

	productCore := ModelProductToCore(productData)
	return productCore, nil
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

func (repo *itemQuery) GetCartItems(cart_id uint) ([]shoppingcartitem.Core, error) {
	var manyCartItems []database.ShoppingCartItem
	tx := repo.db.Where("shopping_cart_id = ?", cart_id).Find(&manyCartItems)
	if tx.Error != nil {
		return nil, tx.Error
	}

	cartItemsCore := ModelGormToCore(manyCartItems)

	return cartItemsCore, nil
}
