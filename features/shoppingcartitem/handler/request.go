package handler

import (
	"Laptop/features/shoppingcartitem"
)

type ItemRequest struct {
	ShoppingCartID uint    `gorm:"not null" json:"cartId" form:"cartId"`
	ProductID      uint    `gorm:"not null" json:"productId" form:"productId"`
	Quantity       uint    `gorm:"not null" json:"quantity" form:"quantity"`
	UnitPrice      float64 `gorm:"not null" json:"unitPrice" form:"unitPrice"`
	TotalPrice     float64 `gorm:"not null" json:"totalPrice" form:"totalPrice"`
}

type CartRequest struct {
	UserID uint   `json:"user_id" form:"user_id"`
	Status string `gorm:"type:string" json:"status" form:"status"`
}

// Mapping dari struct TaskRequest To struct Core Task
func CartReqToCore(req CartRequest) shoppingcartitem.CoreCart {
	return shoppingcartitem.CoreCart{
		UserID: req.UserID,
		Status: req.Status,
	}
}

func RequestToCore(input ItemRequest) shoppingcartitem.Core {
	return shoppingcartitem.Core{
		ShoppingCartID: input.ShoppingCartID,
		ProductID:      input.ProductID,
		Quantity:       input.Quantity,
		UnitPrice:      input.UnitPrice,
		TotalPrice:     input.TotalPrice,
	}
}
