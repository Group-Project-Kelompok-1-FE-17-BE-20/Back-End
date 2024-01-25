package service

import (
	"Laptop/features/store"

	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/labstack/echo/v4"
	// "Laptop/controllers/user"
)

type StoreService struct {
	storeRepo store.StoreDataInterface
	// userRepo  user.UserDataInterface
}

func New(repo store.StoreDataInterface) store.StoreServiceInterface { // userRepo user.UserDataInterface
	return &StoreService{
		storeRepo: repo,
		// userRepo:  userRepo,
	}
}
func (s *StoreService) Create(input store.CoreStore) (uint, error) {
	storeID, err := s.storeRepo.Insert(input)
	if err != nil {
		return 0, err
	}
	return storeID, nil
}

// GetAll implements project.ProjectServiceInterface.
func (s *StoreService) GetAll(userID uint) ([]store.CoreStore, error) {
	result, err := s.storeRepo.SelectAll(userID)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetTaskById implements task.TaskServiceInterface.
func (s *StoreService) GetById(StoreID uint, userID uint) (store.CoreStore, error) {
	result, err := s.storeRepo.Select(StoreID, userID)
	if err != nil {
		return store.CoreStore{}, err
	}
	return result, nil
}

// UpdateTaskById implements task.TaskServiceInterface.
func (s *StoreService) UpdateById(StoreID uint, userID uint, storeData store.CoreStore) error {
	err := s.storeRepo.Update(StoreID, userID, storeData)
	if err != nil {
		return err
	}
	return nil
}

// DeleteTaskById implements task.TaskServiceInterface.
func (s *StoreService) DeleteById(StoreID uint, userID uint) error {
	err := s.storeRepo.Delete(StoreID, userID)
	if err != nil {
		return err
	}
	return nil
}
func (s *StoreService) Photo(c echo.Context) *uploader.UploadResult {
	res := s.storeRepo.Photo(c)
	return res
}
