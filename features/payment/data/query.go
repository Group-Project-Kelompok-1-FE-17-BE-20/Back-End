package data

import (
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
