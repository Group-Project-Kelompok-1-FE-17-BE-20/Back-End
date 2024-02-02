package service

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"Laptop/app/middlewares"
	"Laptop/features/payment"

	"github.com/go-playground/validator/v10"
)

var log = middlewares.Log()

type paymentService struct {
	query    payment.PaymentData
	validate *validator.Validate
}

func New(ud payment.PaymentData, v *validator.Validate) payment.PaymentService {
	return &paymentService{
		query:    ud,
		validate: v,
	}
}

func (ps *paymentService) GetOrderItems(dbRaw *sql.DB, userID uint) (uint, float64) {
	order_id, amount := ps.query.GetOrderItems(dbRaw, userID)
	return order_id, amount
}

// Payment implements payment.PaymentService
func (ps *paymentService) Payment(request payment.PaymentCore) (payment.PaymentCore, error) {
	err := ps.validate.Struct(request)
	if err != nil {
		switch {
		case strings.Contains(err.Error(), "orderID"):
			log.Warn("order_id cannot be empty")
			return payment.PaymentCore{}, errors.New("order_id cannot be empty")
		case strings.Contains(err.Error(), "BankAccount"):
			log.Warn("bank account cannot be empty")
			return payment.PaymentCore{}, errors.New("bank account cannot be empty")
		case strings.Contains(err.Error(), "Amount"):
			log.Warn("amount cannot be empty")
			return payment.PaymentCore{}, errors.New("amount cannot be empty")
		}
	}
	fmt.Printf("log: %v\n", request)

	if request.BankAccount != "bca" && request.BankAccount != "bri" && request.BankAccount != "bni" {
		log.Error("only bca bni, and bni are avalaible atm")
		return payment.PaymentCore{}, errors.New("unsupported bank account")
	}

	result, err := ps.query.Payment(request)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			log.Error("payment record not found")
			return payment.PaymentCore{}, errors.New("payment record not found")
		} else {
			log.Error("internal server error")
			return payment.PaymentCore{}, errors.New("internal server error")
		}
	}

	log.Sugar().Infof("new user has been created: %s", result.ID)
	return result, nil
}

func (ps *paymentService) UpdateStatus(dbRaw *sql.DB, pay payment.PaymentCore) error {
	errUpd := ps.query.UpdateStatus(dbRaw, pay)
	return errUpd
}

// WebhoocksService implements order.OrderServiceInterface.
func (ps *paymentService) CallbackMid(dbRaw *sql.DB, input payment.PaymentCore) error {
	if input.OrderID == "" {
		return errors.New("cannot find order id")
	}

	err := ps.query.CallbackMid(dbRaw, input)
	if err != nil {
		return err
	}

	return nil
}

// UpdatePayment implements payment.PaymentService
func (ps *paymentService) UpdatePayment(request payment.PaymentCore) error {
	err := ps.query.UpdatePayment(request)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			log.Error("not found, error while retrieving payment")
			return errors.New("not found, error while retrieving payment")
		} else if strings.Contains(err.Error(), "no payment record has been updated") {
			log.Error("no payment record has been updated")
			return errors.New("no payment record has been updated")
		} else {
			log.Error("internal server error")
			return errors.New("internal server error")
		}
	}
	return nil
}
