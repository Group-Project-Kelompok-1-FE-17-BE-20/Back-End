package data

import (
	"Laptop/app/database"
	"Laptop/features/product"
	"Laptop/features/shoppingcartitem"
)

func CoreCartToGorm(core shoppingcartitem.CoreCart) database.ShoppingCart {
	return database.ShoppingCart{
		UserID: core.UserID,
		Status: core.Status,
	}
}

func CartGormToCartCore(model database.ShoppingCart) shoppingcartitem.CoreCart {
	return shoppingcartitem.CoreCart{
		ID:     model.ID,
		UserID: model.UserID,
		Status: model.Status,
	}
}

func CoreToModel(input shoppingcartitem.Core) database.ShoppingCartItem {
	return database.ShoppingCartItem{
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
}

func CoretoModelGorm(data []shoppingcartitem.Core) []database.ShoppingCartItem {
	var itemsDataGorm []database.ShoppingCartItem
	for _, input := range data {
		var itemGorm = database.ShoppingCartItem{
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
		itemsDataGorm = append(itemsDataGorm, itemGorm)
	}

	return itemsDataGorm
}

func ModelProductToCore(input database.Product) product.Core {
	return product.Core{
		ID:          input.ID,
		StoreID:     input.StoreID,
		Storage:     input.Storage,
		RAM:         input.RAM,
		Price:       input.Price,
		Description: input.Description,
		Tipe:        input.Tipe,
		Gambar:      input.Gambar,
		Brand:       input.Brand,
		Processor:   input.Processor,
		Categories:  input.Categories,
		Stock:       input.Stock,
	}
}

func ModelToCore(input database.ShoppingCartItem) shoppingcartitem.Core {
	return shoppingcartitem.Core{
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
}

func ModelGormToCore(data []database.ShoppingCartItem) []shoppingcartitem.Core {
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
