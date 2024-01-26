package order

import (
	"Laptop/features/shoppingcartitem"
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
	Productid   uint    `gorm:"not null" json:"prod_id" form:"prod_id"`
	Jumlah      uint    `gorm:"not null" json:"jumlah" form:"jumlah"`
	TotalAmount float64 `gorm:"not null" json:"totalAmount" form:"totalAmount"`
}

// interface untuk Data Layer
type OrderDataInterface interface {
	GetCartID(userID uint) (uint, error)
	GetAllCartItem(cartID uint) ([]shoppingcartitem.Core, error)
	Insert(input Core) error
}

// interface untuk Service Layer
type OrderServiceInterface interface {
	GetCartID(userID uint) (uint, error)
	GetAllCartItem(cartID uint) ([]shoppingcartitem.Core, error)
	Create(input Core) error
}
