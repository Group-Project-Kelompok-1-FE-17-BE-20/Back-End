package store

import "time"

type CoreStore struct {
	ID         uint
	UserID     uint
	NamaToko   string
	AlamatToko string
	ImageToko  string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  time.Time
}
type StoreDataInterface interface {
	Insert(input CoreStore) (uint, error)
	SelectAll(userID uint) ([]CoreStore, error)
	Select(StoreID uint, userID uint) (CoreStore, error)
	Update(StoreID uint, userID uint, storeData CoreStore) error
	Delete(StoreID, userID uint) error
}

type StoreServiceInterface interface {
	Create(input CoreStore) (uint, error)
	GetAll(userID uint) ([]CoreStore, error)
	GetById(StoreID uint, userID uint) (CoreStore, error)
	UpdateById(StoreID uint, userID uint, storeData CoreStore) error
	DeleteById(StoreID uint, userID uint) error
}
