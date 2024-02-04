package handler

import (
	"Laptop/app/database"
	"Laptop/features/payment"
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

type OrderHistoryResponse struct {
	PaymentID         string    `gorm:"not null" json:"id" form:"id"`
	TotalAmount       string    `gorm:"not null" json:"totalAmount" form:"totalAmount"`
	TglOrder          time.Time `gorm:"not null" json:"date_order" form:"date_order"`
	TransactionStatus string    `gorm:"not null" json:"transaction_status" form:"transaction_status"`
}

func CoreToResponseHistory(input []payment.PaymentCore) []OrderHistoryResponse {
	var historiesResponses []OrderHistoryResponse
	for _, v := range input {
		var responseInput = OrderHistoryResponse{
			PaymentID:         v.ID,
			TotalAmount:       v.Amount,
			TglOrder:          v.UpdatedAt,
			TransactionStatus: v.Status,
		}
		historiesResponses = append(historiesResponses, responseInput)
	}

	return historiesResponses
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
