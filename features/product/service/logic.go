package service

import (
	"Laptop/features/product"
	"errors"
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

func (service *productService) Update(id int, input product.Core) error {
	//validasi
	if id <= 0 {
		return errors.New("invalid id")
	}
	//validasi inputan
	// ...
	err := service.productData.Update(id, input)
	return err
}
