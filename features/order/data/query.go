package data

import (
	"Laptop/app/database"
	"Laptop/features/order"
	"Laptop/features/shoppingcartitem"
	"database/sql"
	"fmt"
	"log"
	"time"

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
	tx := repo.db.Where("user_id = ? and status = 'On Going'", userID).First(&cartData)
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

// Insert implements order.OrderDataInterface.
func (repo *orderQuery) GetOrderID(cart_id uint) (uint, error) {
	var orderData database.Order
	tx := repo.db.Where("shopping_cart_id = ? and status = 'On Going'", cart_id).First(&orderData)
	if tx.Error != nil {
		return 0, tx.Error
	}

	orderID := orderData.ID
	return orderID, nil
}

func (repo *orderQuery) CreateOrderItem(orderID uint, input []order.CoreItem) error {
	newOrderItemGorm := ItemsCoreToModel(orderID, input)

	fmt.Println("data item gorm: ", newOrderItemGorm)

	for _, item := range newOrderItemGorm {
		// operasi create untuk setiap elemen item
		tx := repo.db.Create(&item)
		if tx.Error != nil {
			return tx.Error
		}
	}

	// tx := repo.db.Create(&newOrderItemGorm)
	// if tx.Error != nil {
	// 	return tx.Error
	// }

	return nil
}

func (repo *orderQuery) CreateOrderItemSRaw(db *sql.DB, orderID uint, input []order.CoreItem) error {
	// Query SQL dan parameter placeholder
	sqlQuery := "INSERT INTO order_items (order_id, productid, jumlah, total_amount) VALUES (?, ?, ?, ?)"

	// Loop untuk setiap OrderItem
	for _, oi := range input {
		// Eksekusi query SQL
		_, err := db.Exec(sqlQuery, orderID, oi.Productid, oi.Jumlah, oi.TotalAmount)
		if err != nil {
			log.Fatal("cannot run insert query", err)
		}
	}

	return nil
}

func (repo *orderQuery) DetailOrder(db *sql.DB, userID uint) ([]order.DetailOrder, uint, error) {
	// *** read / select data all_accounts *** //
	var itemsOrdered []order.DetailOrder
	var order_id uint

	query := "SELECT order_items.order_id, order_items.productid, products.brand, products.ram, products.storage, " +
		"order_items.jumlah, order_items.total_amount " +
		"FROM shopping_carts " +
		"JOIN orders ON shopping_carts.id = orders.shopping_cart_id " +
		"JOIN order_items ON orders.id = order_items.order_id " +
		"JOIN products ON order_items.productid = products.id WHERE shopping_carts.user_id = ?;"

	rows, errSelect := db.Query(query, userID)
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
		order_id = Row_item.OrderID
	}

	return itemsOrdered, order_id, nil
}

func (repo *orderQuery) DateOrder(db *sql.DB, order_id uint) (time.Time, error) {
	var dateOrder time.Time
	rowID := db.QueryRow("select created_at from orders where id = ?", order_id)
	if err := rowID.Scan(&dateOrder); err != nil {
		//
	}

	return dateOrder, nil
}

// Insert implements order.OrderDataInterface.
func (repo *orderQuery) CreateHistory(input order.CoreHistory) error {
	newHistory := HistoryToModel(input)

	tx := repo.db.Create(&newHistory) // proses query insert
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (repo *orderQuery) Cancel(db *sql.DB, order_id uint) error {
	// result1, errExec := db.Exec("delete from order_items where id = ?", order_id)
	// if errExec != nil {
	// 	log.Fatal("cannot cancel data: ", errExec)
	// }

	// hasilRow1, errRow1 := result1.RowsAffected()
	// if errRow1 != nil {
	// 	log.Fatal("errRow: ", errRow1)
	// } else {
	// 	fmt.Println("berhasil cancel. Row affected ID:", hasilRow1)
	// }

	result2, errExec := db.Exec("delete from orders where id = ?", order_id)
	if errExec != nil {
		log.Fatal("cannot cancel data: ", errExec)
	}

	hasilRow2, errRow2 := result2.RowsAffected()
	if errRow2 != nil {
		log.Fatal("errRow: ", errRow2)
	} else {
		fmt.Println("berhasil cancel. Row affected ID:", hasilRow2)
	}

	return nil
}
