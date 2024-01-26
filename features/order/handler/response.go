package handler

import (
	"Laptop/app/database"
	"time"
)

type OrderResponse struct {
	ID             uint `json:"id" form:"id"`
	ShoppingCartID uint `gorm:"not null" json:"cartId" form:"cartId"`
	Item           []database.OrderItem
	Status         string    `gorm:"not null" json:"status" form:"status"`
	CreatedAt      time.Time `json:"created_at" form:"created_at"`
	UpdatedAt      time.Time `json:"updated_at" form:"updated_at"`
}

// func CoreToResponse(input order.Core) OrderResponse {
// 	return OrderResponse{
// 		ID:             input.ID,
// 		ShoppingCartID: input.ShoppingCartID,
// 		Item:           input.Item,
// 		Status:         input.Status,
// 		CreatedAt:      input.CreatedAt,
// 		UpdatedAt:      input.UpdatedAt,
// 	}
// }

// func CoreToResponseList(data []order.Core) []OrderResponse {
// 	var results []OrderResponse
// 	for _, v := range data {
// 		results = append(results, CoreToResponse(v))
// 	}
// 	return results
// }
