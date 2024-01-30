package handler

import (
	"Laptop/features/shoppingcartitem"
	"time"
)

type ItemResponse struct {
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

type CartResponse struct {
	ID     uint   `json:"id" form:"id"`
	UserID uint   `json:"user_id" form:"user_id"`
	Status string `gorm:"type:string" json:"status" form:"status"`
}

func MapCoreStoreToStoreRes(core shoppingcartitem.CoreCart) CartResponse {
	return CartResponse{
		ID:     core.ID,
		UserID: core.UserID,
		Status: core.Status,
	}
}

func CoreToResponse(input shoppingcartitem.Core) ItemResponse {
	return ItemResponse{
		ID:             input.ID,
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
		CreatedAt:      input.CreatedAt,
		UpdatedAt:      input.UpdatedAt,
	}
}

func CoreToResponseList(data []shoppingcartitem.Core) []ItemResponse {
	var results []ItemResponse
	for _, v := range data {
		results = append(results, CoreToResponse(v))
	}
	return results
}
