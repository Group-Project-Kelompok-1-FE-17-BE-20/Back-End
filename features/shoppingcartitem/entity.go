package shoppingcartitem

import (
	"Laptop/features/product"
	"time"
)

type Core struct {
	ID             uint      `json:"id" form:"id"`
	ShoppingCartID uint      `gorm:"not null" json:"cartId" form:"cartId"`
	ProductID      uint      `gorm:"not null" json:"productId" form:"productId"`
	Tipe           string    `gorm:"type:string" json:"model" form:"model"`
	Price          float64   `gorm:"type:decimal(10,2)" json:"price" form:"price"`
	Processor      string    `gorm:"type:string" json:"processor" form:"processor"`
	RAM            string    `gorm:"type:string" json:"ram" form:"ram"`
	Storage        string    `gorm:"type:string" json:"storage" form:"storage"`
	Quantity       uint      `gorm:"not null" json:"quantity" form:"quantity"`
	TotalPrice     float64   `gorm:"not null" json:"totalPrice" form:"totalPrice"`
	Gambar         string    `gorm:"type:string" json:"gambar" form:"gambar" binding:"uri"`
	CreatedAt      time.Time `json:"created_at" form:"created_at"`
	UpdatedAt      time.Time `json:"updated_at" form:"updated_at"`
}

type CoreCart struct {
	ID        uint
	UserID    uint
	Status    string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

// interface untuk Data Layer
type ItemDataInterface interface {
	InsertCart(input CoreCart) error
	SelectCart(userID uint, status string) (CoreCart, error)
	GetCartID(userID uint) (uint, error)
	GetDataProduct(productID uint) (product.Core, error)
	Insert(input Core) error
	Update(productId uint, input Core) error
	GetItemById(productId uint) (Core, error)
	Delete(input Core) error
	GetCartItems(uint) ([]Core, error)
}

// interface untuk Service Layer
type ItemServiceInterface interface {
	CreateCart(input CoreCart) error
	GetCart(userID uint, status string) (CoreCart, error)
	GetCartID(userID uint) (uint, error)
	GetDataProduct(productID uint) (product.Core, error)
	Create(input Core) error
	Update(productId uint, input Core) error
	GetItemById(productId uint) (Core, error)
	Delete(input Core) error
	GetCartItems(uint) ([]Core, error)
}
