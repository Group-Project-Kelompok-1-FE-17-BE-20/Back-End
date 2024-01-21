package service

import (
	"Laptop/features/product"
)

type productService struct {
	productData product.ProductDataInterface
}

// dependency injection
func New(repo product.ProductDataInterface) product.ProductServiceInterface {
	return &productService{
		productData: repo,
	}
}

func (service *productService) Create(input product.Core) error {
	// logic validation
	err := service.productData.Insert(input)
	return err
}
