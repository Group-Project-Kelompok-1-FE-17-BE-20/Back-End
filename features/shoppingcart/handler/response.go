package handler

import (
	"Laptop/features/shoppingcart"
)

type CartResponse struct {
	UserID uint   `json:"user_id" form:"user_id"`
	Status string `gorm:"type:string" json:"status" form:"status"`
}

// Mapping CorePrject to TaskResponsee
func MapCoreStoreToStoreRes(core shoppingcart.CoreCart) CartResponse {
	return CartResponse{
		UserID: core.UserID,
		Status: core.Status,
	}
}
