package payment

import (
	"database/sql"
	"time"

	"github.com/labstack/echo/v4"
)

type PaymentCore struct {
	ID          string
	OrderID     string `validate:"required" json:"order_id"`
	Amount      string `validate:"required"`
	BankAccount string `validate:"required"`
	VANumber    string
	NamaLengkap string `gorm:"not null" json:"nama_lengkap" form:"nama_lengkap"`
	Alamat      string `gorm:"type:string" json:"alamat" form:"alamat"`
	Status      string `json:"transaction_status"`
	UserID      uint
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type PaymentHandler interface {
	Payment() echo.HandlerFunc
	Notification() echo.HandlerFunc
}

type PaymentService interface {
	GetOrderItems(dbRaw *sql.DB, userID uint) (uint, float64)
	Payment(request PaymentCore) (PaymentCore, error)
	UpdateStatus(dbRaw *sql.DB, pay PaymentCore) error
	UpdatePayment(request PaymentCore) error
	CallbackMid(dbRaw *sql.DB, input PaymentCore) error
}

type PaymentData interface {
	GetOrderItems(dbRaw *sql.DB, userID uint) (uint, float64)
	Payment(request PaymentCore) (PaymentCore, error)
	UpdateStatus(dbRaw *sql.DB, pay PaymentCore) error
	UpdatePayment(request PaymentCore) error
	CallbackMid(dbRaw *sql.DB, input PaymentCore) error
}
