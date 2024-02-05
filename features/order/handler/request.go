package handler

import (
	"Laptop/features/order"
	"Laptop/features/shoppingcartitem"
	"time"
)

type OrderRequest struct {
	ShoppingCartID uint `gorm:"not null" json:"cartId" form:"cartId"`
	Item           []OrderItemRequest
	Status         string `gorm:"not null" json:"status" form:"status"`
}

type OrderItemRequest struct {
	OrderID     uint    `gorm:"not null" json:"orderId" form:"orderId"`
	ProductID   uint    `gorm:"not null" json:"productId" form:"productId"`
	Jumlah      uint    `gorm:"not null" json:"jumlah" form:"jumlah"`
	TotalAmount float64 `gorm:"not null" json:"totalAmount" form:"totalAmount"`
}

type HistoryRequest struct {
	PaymentID         string    `gorm:"not null" json:"id" form:"id"`
	TotalAmount       float64   `gorm:"not null" json:"totalAmount" form:"totalAmount"`
	TglOrder          time.Time `gorm:"not null" json:"date_order" form:"date_order"`
	TransactionStatus string    `gorm:"not null" json:"transaction_status" form:"transaction_status"`
}

func ResGetRequest(data []shoppingcartitem.Core) []OrderItemRequest {
	var results []OrderItemRequest
	for _, input := range data {
		var item = OrderItemRequest{
			ProductID:   input.ProductID,
			Jumlah:      input.Quantity,
			TotalAmount: input.TotalPrice,
		}
		results = append(results, item)
	}
	return results
}

func ItemRequestToCoreItem(data []OrderItemRequest) []order.CoreItem {
	var results []order.CoreItem
	for _, input := range data {
		var item = order.CoreItem{
			ProductID:   input.ProductID,
			Jumlah:      input.Jumlah,
			TotalAmount: input.TotalAmount,
		}
		results = append(results, item)
	}
	return results
}

func IdAndItemToReq(order_id uint, data []order.CoreItem) []OrderItemRequest {
	var results []OrderItemRequest
	for _, input := range data {
		var item = OrderItemRequest{
			OrderID:     order_id,
			ProductID:   input.ProductID,
			Jumlah:      input.Jumlah,
			TotalAmount: input.TotalAmount,
		}
		results = append(results, item)
	}
	return results
}

// func IdAndItemToCore(data []OrderItemRequest) []order.CoreItem {
// 	var results []order.CoreItem
// 	for _, input := range data {
// 		var item = order.CoreItem{
// 			OrderID:     input.OrderID,
// 			Productid:   input.Productid,
// 			Jumlah:      input.Jumlah,
// 			TotalAmount: input.TotalAmount,
// 		}
// 		results = append(results, item)
// 	}
// 	return results
// }

func RequestToCore(input OrderRequest) order.Core {
	inputCoreItem := ItemRequestToCoreItem(input.Item)
	return order.Core{
		ShoppingCartID: input.ShoppingCartID,
		Item:           inputCoreItem,
		Status:         input.Status,
	}
}

func HistoryToCore(input HistoryRequest) order.CoreHistory {
	return order.CoreHistory{
		PaymentID:         input.PaymentID,
		TotalAmount:       input.TotalAmount,
		TglOrder:          input.TglOrder,
		TransactionStatus: input.TransactionStatus,
	}
}
