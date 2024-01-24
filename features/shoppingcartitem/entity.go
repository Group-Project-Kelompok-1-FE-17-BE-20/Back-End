package shoppingcartitem

import (
	"time"
)

type Core struct {
	ID             uint      `json:"id" form:"id"`
	ShoppingCartID uint      `gorm:"not null" json:"cartId" form:"cartId"`
	ProductID      uint      `gorm:"not null" json:"productId" form:"productId"`
	Quantity       uint      `gorm:"not null" json:"quantity" form:"quantity"`
	UnitPrice      float64   `gorm:"not null" json:"unitPrice" form:"unitPrice"`
	TotalPrice     float64   `gorm:"not null" json:"totalPrice" form:"totalPrice"`
	CreatedAt      time.Time `json:"created_at" form:"created_at"`
	UpdatedAt      time.Time `json:"updated_at" form:"updated_at"`
}

// interface untuk Data Layer
type ItemDataInterface interface {
	GetCartID(userID uint) (uint, error)
	GetPrice(productID uint) (float64, error)
	Insert(input Core) error
	Update(productId uint, input Core) error
	GetItemById(productId uint) (Core, error)
	Delete(input Core) error
}

// interface untuk Service Layer
type ItemServiceInterface interface {
	GetCartID(userID uint) (uint, error)
	GetPrice(productID uint) (float64, error)
	Create(input Core) error
	Update(productId uint, input Core) error
	GetItemById(productId uint) (Core, error)
	Delete(input Core) error
}
