package handler

import (
	"Laptop/features/shoppingcartitem"
)

type ItemRequest struct {
	ShoppingCartID uint    `gorm:"not null" json:"cartId" form:"cartId"`
	ProductID      uint    `gorm:"not null" json:"productId" form:"productId"`
	Tipe           string  `gorm:"type:string" json:"model" form:"model"`
	Price          float64 `gorm:"type:decimal(10,2)" json:"price" form:"price"`
	Processor      string  `gorm:"type:string" json:"processor" form:"processor"`
	RAM            string  `gorm:"type:string" json:"ram" form:"ram"`
	Storage        string  `gorm:"type:string" json:"storage" form:"storage"`
	Quantity       uint    `gorm:"not null" json:"quantity" form:"quantity"`
	TotalPrice     float64 `gorm:"not null" json:"totalPrice" form:"totalPrice"`
	Gambar         string  `gorm:"type:string" json:"gambar" form:"gambar" binding:"uri"`
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
		Tipe:           input.Tipe,
		Price:          input.Price,
		Processor:      input.Processor,
		RAM:            input.RAM,
		Storage:        input.Storage,
		Quantity:       input.Quantity,
		TotalPrice:     input.TotalPrice,
		Gambar:         input.Gambar,
	}
}
