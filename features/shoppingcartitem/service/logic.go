package service

import (
	"Laptop/features/product"
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

func (service *itemService) CreateCart(input shoppingcartitem.CoreCart) error {
	// logic validation
	err := service.itemData.InsertCart(input)
	return err
}

// GetTaskById implements task.TaskServiceInterface.
func (service *itemService) GetCart(userID uint, status string) (shoppingcartitem.CoreCart, error) {
	result, err := service.itemData.SelectCart(userID, status)
	if err != nil {
		return shoppingcartitem.CoreCart{}, err
	}
	return result, nil
}

func (service *itemService) GetCartID(input uint) (uint, error) {
	// logic validation
	res, err := service.itemData.GetCartID(input)
	return res, err
}

func (service *itemService) GetDataProduct(input uint) (product.Core, error) {
	// logic validation
	res, err := service.itemData.GetDataProduct(input)
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

func (service *itemService) GetCartItems(input uint) ([]shoppingcartitem.Core, error) {
	// logic validation
	res, err := service.itemData.GetCartItems(input)
	return res, err
}
