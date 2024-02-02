package data

import (
	"database/sql"
	"errors"
	"fmt"

	"Laptop/app/database"
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
		"JOIN users ON shopping_carts.user_id = users.id WHERE users.id = ? " +
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

// Update implements user.UserDataInterface.
func (pq *paymentQuery) CallbackMid(input payment.PaymentCore) error {
	dataGorm := CoreToModel(input)
	tx := pq.db.Model(&database.Payment{}).Where("order_id = ?", input.OrderID).Updates(dataGorm)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New("error record not found ")
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
