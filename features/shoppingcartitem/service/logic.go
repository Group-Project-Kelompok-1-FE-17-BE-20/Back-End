package service

import (
	"Laptop/features/shoppingcartitem"
)

type itemService struct {
	itemData shoppingcartitem.ItemDataInterface
}

// dependency injection
func New(repo shoppingcartitem.ItemDataInterface) shoppingcartitem.ItemServiceInterface {
	return &itemService{
		itemData: repo,
	}
}

func (service *itemService) GetCartID(input uint) (uint, error) {
	// logic validation
	res, err := service.itemData.GetCartID(input)
	return res, err
}

func (service *itemService) GetPrice(input uint) (float64, error) {
	// logic validation
	res, err := service.itemData.GetPrice(input)
	return res, err
}

func (service *itemService) Create(input shoppingcartitem.Core) error {
	// logic validation
	err := service.itemData.Insert(input)
	return err
}

func (service *itemService) Update(productId uint, input shoppingcartitem.Core) error {
	// logic validation
	err := service.itemData.Update(productId, input)
	return err
}

func (service *itemService) GetItemById(productId uint) (shoppingcartitem.Core, error) {
	// logic validation
	res, err := service.itemData.GetItemById(productId)
	return res, err
}

func (service *itemService) Delete(input shoppingcartitem.Core) error {
	// logic validation
	err := service.itemData.Delete(input)
	return err
}
