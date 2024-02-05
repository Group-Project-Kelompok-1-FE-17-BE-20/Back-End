package data

import (
	"Laptop/app/database"
	"Laptop/features/order"
	"Laptop/features/payment"
	"Laptop/features/shoppingcartitem"
)

func GetItemsData(data []database.ShoppingCartItem) []database.ShoppingCartItem {
	var itemsDataGorm []database.ShoppingCartItem
	for _, input := range data {
		var itemGorm = database.ShoppingCartItem{
			ProductID:  input.ProductID,
			Quantity:   input.Quantity,
			TotalPrice: input.TotalPrice,
		}
		itemsDataGorm = append(itemsDataGorm, itemGorm)
	}

	return itemsDataGorm
}

func GetItemsDataCore(data []database.ShoppingCartItem) []shoppingcartitem.Core {
	var itemsDataCore []shoppingcartitem.Core
	for _, input := range data {
		var itemCore = shoppingcartitem.Core{
			ProductID:  input.ProductID,
			Quantity:   input.Quantity,
			TotalPrice: input.TotalPrice,
		}
		itemsDataCore = append(itemsDataCore, itemCore)
	}

	return itemsDataCore
}

func CoreItemToItemGorm(data []order.CoreItem) []database.OrderItem {
	var results []database.OrderItem
	for _, input := range data {
		var item = database.OrderItem{
			ProductID:   input.ProductID,
			Jumlah:      input.Jumlah,
			TotalAmount: input.TotalAmount,
		}
		results = append(results, item)
	}
	return results
}

func CoreToModel(input order.Core) database.Order {
	itemGorm := CoreItemToItemGorm(input.Item)
	return database.Order{
		ShoppingCartID: input.ShoppingCartID,
		Item:           itemGorm,
		Status:         input.Status,
	}
}

func GormCartItemsToCore(data []database.ShoppingCartItem) []shoppingcartitem.Core {
	var itemsData []shoppingcartitem.Core
	for _, input := range data {
		var itemInput = shoppingcartitem.Core{
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
		}
		itemsData = append(itemsData, itemInput)
	}

	return itemsData
}

func ItemsCoreToModel(orderID uint, data []order.CoreItem) []database.OrderItem {
	var itemsData []database.OrderItem
	for _, input := range data {
		var itemInput = database.OrderItem{
			OrderID:     orderID,
			ProductID:   input.ProductID,
			Jumlah:      input.Jumlah,
			TotalAmount: input.TotalAmount,
		}
		itemsData = append(itemsData, itemInput)
	}

	return itemsData
}

func HistoryToModel(input payment.PaymentCore) database.Payment {
	return database.Payment{
		ID:        input.ID,
		Amount:    input.Amount,
		UpdatedAt: input.UpdatedAt,
		Status:    input.Status,
	}
}

func SliceHistoryToModel(input []payment.PaymentCore) []database.Payment {
	var historiesData []database.Payment
	for _, value := range input {
		var historyInput = database.Payment{
			ID:        value.ID,
			Amount:    value.Amount,
			UpdatedAt: value.UpdatedAt,
			Status:    value.Status,
		}
		historiesData = append(historiesData, historyInput)
	}

	return historiesData
}

// func CoretoModelGorm(data []order.Core) []database.Order {
// 	var orderDataGorm []database.Order
// 	for _, input := range data {
// 		var orderGorm = database.Order{
// 			ShoppingCartID: input.ShoppingCartID,
// 			Productid:      input.Productid,
// 			Jumlah:         input.Jumlah,
// 			TotalAmount:    input.TotalAmount,
// 			Status:         input.Status,
// 		}
// 		orderDataGorm = append(orderDataGorm, orderGorm)
// 	}

// 	return orderDataGorm
// }

// func ModelToCore(input database.Order) order.Core {
// 	return order.Core{
// 		ShoppingCartID: input.ShoppingCartID,
// 		Productid:      input.Productid,
// 		Jumlah:         input.Jumlah,
// 		TotalAmount:    input.TotalAmount,
// 		Status:         input.Status,
// 	}
// }

// func ModelGormToCore(data []database.Order) []order.Core {
// 	var orderData []order.Core
// 	for _, input := range data {
// 		var orderInput = order.Core{
// 			ShoppingCartID: input.ShoppingCartID,
// 			Productid:      input.Productid,
// 			Jumlah:         input.Jumlah,
// 			TotalAmount:    input.TotalAmount,
// 			Status:         input.Status,
// 		}
// 		orderData = append(orderData, orderInput)
// 	}

// 	return orderData
// }
