package shoppingcart

import "time"

type CoreCart struct {
	ID        uint
	UserID    uint
	Status    string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
type CartDataInterface interface {
	Insert(input CoreCart) error
	SelectCart(userID uint) (CoreCart, error)
}

type CartServiceInterface interface {
	Create(input CoreCart) error
	GetCart(userID uint) (CoreCart, error)
}
