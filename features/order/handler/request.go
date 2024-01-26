package handler

import (
	"Laptop/features/order"
	"Laptop/features/shoppingcartitem"
)

type OrderRequest struct {
	ShoppingCartID uint `gorm:"not null" json:"cartId" form:"cartId"`
	Item           []OrderItemRequest
	Status         string `gorm:"not null" json:"status" form:"status"`
}

type OrderItemRequest struct {
	Productid   uint    `gorm:"not null" json:"prod_id" form:"prod_id"`
	Jumlah      uint    `gorm:"not null" json:"jumlah" form:"jumlah"`
	TotalAmount float64 `gorm:"not null" json:"totalAmount" form:"totalAmount"`
}

func ResGetRequest(data []shoppingcartitem.Core) []OrderItemRequest {
	var results []OrderItemRequest
	for _, input := range data {
		var item = OrderItemRequest{
			Productid:   input.ProductID,
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
			Productid:   input.Productid,
			Jumlah:      input.Jumlah,
			TotalAmount: input.TotalAmount,
		}
		results = append(results, item)
	}
	return results
}

func RequestToCore(input OrderRequest) order.Core {
	inputCoreItem := ItemRequestToCoreItem(input.Item)
	return order.Core{
		ShoppingCartID: input.ShoppingCartID,
		Item:           inputCoreItem,
		Status:         input.Status,
	}
}
