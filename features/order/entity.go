package order

import (
	"Laptop/features/payment"
	"Laptop/features/shoppingcartitem"
	"database/sql"
	"time"
)

type Core struct {
	ID             uint       `json:"id" form:"id"`
	ShoppingCartID uint       `gorm:"not null" json:"cartId" form:"cartId"`
	Item           []CoreItem `gorm:"foreignKey:OrderID"`
	Status         string     `gorm:"not null" json:"status" form:"status"`
	CreatedAt      time.Time  `json:"created_at" form:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at" form:"updated_at"`
}

type CoreItem struct {
	ProductID   uint    `gorm:"not null" json:"productId" form:"productId"`
	Jumlah      uint    `gorm:"not null" json:"jumlah" form:"jumlah"`
	TotalAmount float64 `gorm:"not null" json:"totalAmount" form:"totalAmount"`
}

type DetailOrder struct {
	OrderID     uint    `gorm:"not null" json:"orderId" form:"orderId"`
	ProductID   uint    `gorm:"not null" json:"productId" form:"productId"`
	Brand       string  `gorm:"type:string" json:"brand" form:"brand"`
	RAM         string  `gorm:"type:string" json:"ram" form:"ram"`
	Storage     string  `gorm:"type:string" json:"storage" form:"storage"`
	Jumlah      uint    `gorm:"not null" json:"jumlah" form:"jumlah"`
	TotalAmount float64 `gorm:"not null" json:"totalAmount" form:"totalAmount"`
}

type CoreHistory struct {
	PaymentID         string    `gorm:"not null" json:"id" form:"id"`
	TotalAmount       float64   `gorm:"not null" json:"totalAmount" form:"totalAmount"`
	TglOrder          time.Time `gorm:"not null" json:"date_order" form:"date_order"`
	TransactionStatus string    `gorm:"not null" json:"transaction_status" form:"transaction_status"`
}

// interface untuk Data Layer
type OrderDataInterface interface {
	GetCartID(userID uint) (uint, error)
	GetAllCartItem(cartID uint) ([]shoppingcartitem.Core, error)
	Insert(input Core) error
	GetOrderID(cart_id uint) (uint, error)
	CreateOrderItem(order_id uint, input []CoreItem) error
	CreateOrderItemSRaw(db *sql.DB, order_id uint, input []CoreItem) error
	DetailOrder(db *sql.DB, userID uint) ([]DetailOrder, uint, error)
	DateOrder(db *sql.DB, orderID uint) (time.Time, error)
	//CreateHistory(CoreHistory) error
	Cancel(db *sql.DB, orderID uint) error
	GetAllPayments(db *sql.DB, user_id uint) ([]payment.PaymentCore, error)
}

// interface untuk Service Layer
type OrderServiceInterface interface {
	GetCartID(userID uint) (uint, error)
	GetAllCartItem(cartID uint) ([]shoppingcartitem.Core, error)
	Create(input Core) error
	GetOrderID(cart_id uint) (uint, error)
	CreateOrderItem(order_id uint, input []CoreItem) error
	CreateOrderItemSRaw(db *sql.DB, order_id uint, input []CoreItem) error
	DetailOrder(db *sql.DB, userID uint) ([]DetailOrder, uint, error)
	DateOrder(db *sql.DB, orderID uint) (time.Time, error)
	//CreateHistory(CoreHistory) error
	Cancel(db *sql.DB, orderID uint) error
	GetAllPayments(db *sql.DB, user_id uint) ([]payment.PaymentCore, error)
}
