package service

import (
	"Laptop/features/order"
	"Laptop/features/shoppingcartitem"
	"database/sql"
	"time"
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

func (service *orderService) DetailOrder(input *sql.DB, userID uint) ([]order.DetailOrder, uint, error) {
	result, id, err := service.orderData.DetailOrder(input, userID)
	return result, id, err
}

func (service *orderService) DateOrder(input *sql.DB, order_id uint) (time.Time, error) {
	result, err := service.orderData.DateOrder(input, order_id)
	return result, err
}

func (service *orderService) CreateHistory(input order.CoreHistory) error {
	err := service.orderData.CreateHistory(input)
	return err
}

func (service *orderService) Cancel(db *sql.DB, order_id uint) error {
	err := service.orderData.Cancel(db, order_id)
	return err
}
