package handler

import (
	"Laptop/features/shoppingcart"
)

type CartResponse struct {
	ID     uint
	UserID uint   `json:"user_id" form:"user_id"`
	Status string `gorm:"type:string" json:"status" form:"status"`
}

// Mapping CorePrject to TaskResponsee
func MapCoreStoreToStoreRes(core shoppingcart.CoreCart) CartResponse {
	return CartResponse{
		ID:     core.ID,
		UserID: core.UserID,
		Status: core.Status,
	}
}
