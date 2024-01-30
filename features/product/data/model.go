package data

import (
	"Laptop/app/database"
	"Laptop/features/product"
)

func CoreToModel(input product.Core) database.Product {
	return database.Product{
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

func CoretoModelGorm(data []product.Core) []database.Product {
	var productsDataGorm []database.Product
	for _, input := range data {
		var productGorm = database.Product{
			Storage:     input.Storage,
			RAM:         input.RAM,
			Price:       input.Price,
			Description: input.Description,
			Tipe:        input.Tipe,
			Gambar:      input.Gambar,
			Brand:       input.Gambar,
			Processor:   input.Processor,
			Categories:  input.Categories,
			Stock:       input.Stock,
		}
		productsDataGorm = append(productsDataGorm, productGorm)
	}

	return productsDataGorm
}

func ModelToCore(input database.Product) product.Core {
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

func ModelGormToCore(data []database.Product) []product.Core {
	var productsData []product.Core
	for _, input := range data {
		var productInput = product.Core{
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
		productsData = append(productsData, productInput)
	}

	return productsData
}
