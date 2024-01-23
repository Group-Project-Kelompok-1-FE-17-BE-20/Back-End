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

func (service *itemService) Create(input shoppingcartitem.Core) error {
	// logic validation
	err := service.itemData.Insert(input)
	return err
}
