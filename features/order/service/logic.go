package service

import (
	"Laptop/features/order"
	"Laptop/features/shoppingcartitem"
)

type orderService struct {
	orderData order.OrderDataInterface
}

// dependency injection
func New(repo order.OrderDataInterface) order.OrderServiceInterface {
	return &orderService{
		orderData: repo,
	}
}

func (service *orderService) GetCartID(input uint) (uint, error) {
	// logic validation
	res, err := service.orderData.GetCartID(input)
	return res, err
}

func (service *orderService) GetAllCartItem(input uint) ([]shoppingcartitem.Core, error) {
	// logic validation
	res, err := service.orderData.GetAllCartItem(input)
	return res, err
}

// Create implements order.OrderServiceInterface.
func (service *orderService) Create(input order.Core) error {
	err := service.orderData.Insert(input)
	return err
}
