package handler

import (
	"Laptop/features/product"
)

type ProductRequest struct {
	Storage     string  `gorm:"type:string" json:"storage" form:"storage"`
	RAM         string  `gorm:"type:string" json:"ram" form:"ram"`
	Price       float64 `gorm:"type:decimal(10,2)" json:"price" form:"price"`
	Description string  `gorm:"type:string" json:"description" form:"description"`
	Tipe        string  `gorm:"type:string" json:"model" form:"model"`
	Gambar      string  `gorm:"type:string" json:"gambar" form:"gambar" binding:"uri"`
	Brand       string  `gorm:"type:string" json:"brand" form:"brand"`
	Processor   string  `gorm:"type:string" json:"processor" form:"processor"`
	Categories  string  `gorm:"type:string" json:"categories" form:"categories"`
	Stock       int     `gorm:"type:integer" json:"stock" form:"stock"`
}

func RequestToCore(input ProductRequest) product.Core {
	return product.Core{
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
}
