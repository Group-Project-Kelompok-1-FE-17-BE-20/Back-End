package data

import (
	"Laptop/app/database"
	"Laptop/features/shoppingcartitem"
)

func CoreToModel(input shoppingcartitem.Core) database.ShoppingCartItem {
	return database.ShoppingCartItem{
		ShoppingCartID: input.ShoppingCartID,
		ProductID:      input.ProductID,
		Quantity:       input.Quantity,
		UnitPrice:      input.UnitPrice,
		TotalPrice:     input.TotalPrice,
	}
}

func CoretoModelGorm(data []shoppingcartitem.Core) []database.ShoppingCartItem {
	var itemsDataGorm []database.ShoppingCartItem
	for _, input := range data {
		var itemGorm = database.ShoppingCartItem{
			ShoppingCartID: input.ShoppingCartID,
			ProductID:      input.ProductID,
			Quantity:       input.Quantity,
			UnitPrice:      input.UnitPrice,
			TotalPrice:     input.TotalPrice,
		}
		itemsDataGorm = append(itemsDataGorm, itemGorm)
	}

	return itemsDataGorm
}

func ModelToCore(input database.ShoppingCartItem) shoppingcartitem.Core {
	return shoppingcartitem.Core{
		ShoppingCartID: input.ShoppingCartID,
		ProductID:      input.ProductID,
		Quantity:       input.Quantity,
		UnitPrice:      input.UnitPrice,
		TotalPrice:     input.TotalPrice,
	}
}

func ModelGormToCore(data []database.ShoppingCartItem) []shoppingcartitem.Core {
	var itemsData []shoppingcartitem.Core
	for _, input := range data {
		var itemInput = shoppingcartitem.Core{
			ShoppingCartID: input.ShoppingCartID,
			ProductID:      input.ProductID,
			Quantity:       input.Quantity,
			UnitPrice:      input.UnitPrice,
			TotalPrice:     input.TotalPrice,
		}
		itemsData = append(itemsData, itemInput)
	}

	return itemsData
}
