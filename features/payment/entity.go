package payment

import (
	"time"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type PaymentCore struct {
	gorm.Model
	ID          string
	OrderID     string `validate:"required"`
	Amount      string `validate:"required"`
	BankAccount string `validate:"required"`
	VANumber    string
	Status      string
	UserID      uint
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type PaymentHandler interface {
	Payment() echo.HandlerFunc
	Notification() echo.HandlerFunc
}

type PaymentService interface {
	Payment(request PaymentCore) (PaymentCore, error)
	UpdatePayment(request PaymentCore) error
}

type PaymentData interface {
	Payment(request PaymentCore) (PaymentCore, error)
	UpdatePayment(request PaymentCore) error
}