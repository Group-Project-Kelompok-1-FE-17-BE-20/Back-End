package handler

import (
	"Laptop/features/shoppingcartitem"
	"time"
)

type ItemResponse struct {
	ID             uint      `json:"id" form:"id"`
	ShoppingCartID uint      `gorm:"not null" json:"cartId" form:"cartId"`
	ProductID      uint      `gorm:"not null" json:"productId" form:"productId"`
	Quantity       uint      `gorm:"not null" json:"quantity" form:"quantity"`
	UnitPrice      float64   `gorm:"not null" json:"unitPrice" form:"unitPrice"`
	TotalPrice     float64   `gorm:"not null" json:"totalPrice" form:"totalPrice"`
	CreatedAt      time.Time `json:"created_at" form:"created_at"`
	UpdatedAt      time.Time `json:"updated_at" form:"updated_at"`
}

type CartResponse struct {
	ID     uint   `json:"id" form:"id"`
	UserID uint   `json:"user_id" form:"user_id"`
	Status string `gorm:"type:string" json:"status" form:"status"`
}

// Mapping CorePrject to TaskResponsee
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
		Quantity:       input.Quantity,
		UnitPrice:      input.UnitPrice,
		TotalPrice:     input.TotalPrice,
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
