package data

import (
	"Laptop/app/database"
	"Laptop/features/order"
	"Laptop/features/shoppingcartitem"
	"database/sql"
	"log"

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

func (repo *orderQuery) DetailOrder(db *sql.DB) ([]order.DetailOrder, error) {
	// *** read / select data all_accounts *** //
	var itemsOrdered []order.DetailOrder

	orderID := 1

	query := "SELECT order_items.order_id, order_items.productid, products.brand, products.ram, products.storage, " +
		"order_items.jumlah, order_items.total_amount " +
		"FROM shopping_carts " +
		"JOIN orders ON shopping_carts.id = orders.shopping_cart_id " +
		"JOIN order_items ON orders.id = order_items.order_id " +
		"JOIN products ON order_items.productid = products.id WHERE order_items.order_id = ?;"

	rows, errSelect := db.Query(query, orderID)
	if errSelect != nil {
		log.Fatal("cannot run select query: ", errSelect)
	}

	for rows.Next() {
		var Row_item order.DetailOrder
		errScan := rows.Scan(&Row_item.OrderID, &Row_item.Productid, &Row_item.Brand, &Row_item.RAM, &Row_item.Storage, &Row_item.Jumlah, &Row_item.TotalAmount)
		if errScan != nil {
			log.Fatal("cannot run scan query: ", errScan.Error())
		}
		itemsOrdered = append(itemsOrdered, Row_item)
	}

	return itemsOrdered, nil
}
