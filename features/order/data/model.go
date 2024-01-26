package data

import (
	"Laptop/app/database"
	"Laptop/features/order"
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
			Productid:   input.Productid,
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
