package payment

import (
	"database/sql"
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
	NamaLengkap string `gorm:"not null" json:"nama_lengkap" form:"nama_lengkap"`
	Alamat      string `gorm:"type:string" json:"alamat" form:"alamat"`
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
	GetOrderItems(dbRaw *sql.DB, userID uint) (uint, float64)
	Payment(request PaymentCore) (PaymentCore, error)
	UpdatePayment(request PaymentCore) error
}

type PaymentData interface {
	GetOrderItems(dbRaw *sql.DB, userID uint) (uint, float64)
	Payment(request PaymentCore) (PaymentCore, error)
	UpdatePayment(request PaymentCore) error
}
