package service

import (
	"Laptop/features/shoppingcart"
	// "Laptop/controllers/user"
)

type cartService struct {
	cartRepo shoppingcart.CartDataInterface
	// userRepo  user.UserDataInterface
}

// dependency injection
func New(repo shoppingcart.CartDataInterface) shoppingcart.CartServiceInterface {
	return &cartService{
		cartRepo: repo,
	}
}

func (service *cartService) Create(input shoppingcart.CoreCart) error {
	// logic validation
	err := service.cartRepo.Insert(input)
	return err
}

// GetTaskById implements task.TaskServiceInterface.
func (s *cartService) GetCart(userID uint) (shoppingcart.CoreCart, error) {
	result, err := s.cartRepo.SelectCart(userID)
	if err != nil {
		return shoppingcart.CoreCart{}, err
	}
	return result, nil
}
