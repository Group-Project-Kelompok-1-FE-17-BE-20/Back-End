package data

import (
	"database/sql"
	"errors"
	"fmt"

	"Laptop/app/middlewares"
	"Laptop/features/payment"

	"gorm.io/gorm"
)

var log = middlewares.Log()

type paymentQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) payment.PaymentData {
	return &paymentQuery{
		db: db,
	}
}

func (pq *paymentQuery) GetOrderItems(dbRaw *sql.DB, userID uint) (uint, float64) {
	var order_id uint
	var total_amount float64

	query := "SELECT order_items.order_id, SUM(order_items.total_amount) " +
		"FROM order_items " +
		"JOIN orders ON order_items.order_id = orders.id " +
		"JOIN shopping_carts ON orders.shopping_cart_id = shopping_carts.id " +
		"JOIN users ON shopping_carts.user_id = users.id WHERE users.id = ? and orders.status = 'On Going' " +
		"GROUP BY order_items.order_id;"

	rowID := dbRaw.QueryRow(query, userID)

	if err := rowID.Scan(&order_id, &total_amount); err != nil {
		log.Fatal("cannot scan data: ")
	}

	return order_id, total_amount
}

func (pq *paymentQuery) Payment(request payment.PaymentCore) (payment.PaymentCore, error) {
	paymentData := chargeMidtrans(request)

	result := pq.db.Create(&paymentData)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		log.Error("payment record not found")
		return payment.PaymentCore{}, errors.New("payment record not found")
	}

	if result.RowsAffected == 0 {
		log.Warn("no charge payment has been created")
		return payment.PaymentCore{}, errors.New("row affected : 0")
	}

	if result.Error != nil {
		log.Error("error while charging payment")
		return payment.PaymentCore{}, errors.New("internal server error")
	}
	fmt.Printf("log ppayment data : %v\n", paymentData)
	fmt.Printf("log ppayment model: %v\n", paymentModels(paymentData))
	return paymentModels(paymentData), nil
}

func (pq *paymentQuery) UpdateStatus(dbRaw *sql.DB, pay payment.PaymentCore) error {
	// Buat pernyataan SQL UPDATE
	query := "UPDATE payments SET status = ? WHERE transaction_id = ?"

	// Eksekusi pernyataan SQL
	_, err := dbRaw.Exec(query, pay.Status, pay.ID)
	if err != nil {
		return err
	}

	return nil
}

func (pq *paymentQuery) UpdateStock(dbRaw *sql.DB, orderID string) error {
	// menghitung stock dari product yang dibeli
	query1 := "SELECT products.id, products.stock FROM products " +
		"JOIN order_items ON products.id = order_items.product_id " +
		"JOIN orders ON order_items.order_id = orders.id " +
		"WHERE orders.id = ?"

	rows, err1 := dbRaw.Query(query1, orderID)
	if err1 != nil {
		return err1
	}

	var listProductID []uint
	var listStock []int
	for rows.Next() {
		var product_id uint
		var stock int
		if err := rows.Scan(&product_id, &stock); err != nil {
			return err
		}
		listProductID = append(listProductID, product_id)
		listStock = append(listStock, stock)

	}

	fmt.Println("daftar product id: ", listProductID)
	fmt.Println("daftar stock: ", listStock)

	// menghitung banyak produk yang dibeli
	query2 := "SELECT order_items.jumlah FROM order_items " +
		"JOIN orders ON order_items.order_id = orders.id " +
		"WHERE orders.id = ?"

	rows, err2 := dbRaw.Query(query2, orderID)
	if err2 != nil {
		return err2
	}

	var listJumlah []int
	for rows.Next() {
		var jumlah int
		if err := rows.Scan(&jumlah); err != nil {
			return err
		}
		listJumlah = append(listJumlah, jumlah)
	}

	// proses pengurangan stock
	if len(listStock) != len(listJumlah) {
		fmt.Println("Panjang slice tidak sama.")
	}

	var result []int
	for i := 0; i < len(listStock); i++ {
		listStock[i] -= listJumlah[i]
		itemRes := listStock[i]
		result = append(result, itemRes)
	}

	// proses update stock baru ke database
	for i := 0; i < len(result); i++ {
		query := "UPDATE products SET stock = ? WHERE id = ?"
		tx, err := dbRaw.Exec(query, result[i], listProductID[i])
		if err != nil {
			return err
		}

		stockRows, _ := tx.RowsAffected()
		fmt.Println("Stock rows affected: ", stockRows)
	}

	return nil
}

// Update implements user.UserDataInterface.
func (pq *paymentQuery) CallbackMid(dbRaw *sql.DB, input payment.PaymentCore) error {
	// dataGorm := CoreToModel(input)
	// tx := pq.db.Model(&database.Payment{}).Where("order_id = ?", dataGorm.OrderID).Updates(dataGorm.Status)
	// if tx.Error != nil {
	// 	fmt.Println("Error updating payment status:", tx.Error)
	// 	return errors.New("failed to update payment status")
	// }

	fmt.Println("data input buat update: ", input)

	// update status payment
	query := "UPDATE payments SET status = ? WHERE order_id = ?"
	tx, err := dbRaw.Exec(query, input.Status, input.OrderID)
	if err != nil {
		return errors.New("failed to update payment status")
	}

	affected, _ := tx.RowsAffected()
	if affected == 0 {
		return errors.New("error record not found ")
	} else if affected != 0 {
		// update orders
		query1 := "UPDATE orders SET status = 'Selesai' WHERE id = ?"
		_, err := dbRaw.Exec(query1, input.OrderID)
		if err != nil {
			return err
		}

		// update shopping carts
		query2 := "UPDATE shopping_carts " +
			"SET status = 'Selesai' " +
			"WHERE id IN ( " +
			"SELECT shopping_carts.id " +
			"FROM shopping_carts " +
			"JOIN orders ON shopping_carts.id = orders.shopping_cart_id " +
			"WHERE orders.id = ? " +
			");"

		_, err2 := dbRaw.Exec(query2, input.OrderID)
		if err2 != nil {
			return err2
		}

		// update stock
		errStock := pq.UpdateStock(dbRaw, input.OrderID)
		if errStock != nil {
			return errStock
		}
	}

	return nil
}

// UpdatePayment implements payment.PaymentData
func (pq *paymentQuery) UpdatePayment(request payment.PaymentCore) error {
	req := paymentEntities(request)
	log.Sugar().Infof("callback midtrans status: %s, order ID: %s, transaction ID: %s",
		req.Status, req.OrderID, req.ID)
	query := pq.db.Table("payments").
		Where("id = ? AND order_id = ?", request.ID, request.OrderID).
		Updates(map[string]interface{}{
			"status": request.Status,
		})
	if errors.Is(query.Error, gorm.ErrRecordNotFound) {
		log.Error("user profile record not found")
		return errors.New("user profile record not found")
	}

	if query.RowsAffected == 0 {
		log.Warn("no payment record has been updated")
		return errors.New("no payment record has been updated")
	}

	if query.Error != nil {
		log.Error("error while updating payment status")
		return errors.New("internal server error")
	}

	return nil
}
