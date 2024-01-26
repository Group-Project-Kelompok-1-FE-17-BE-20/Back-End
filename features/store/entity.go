package store

import (
	"time"

	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/labstack/echo/v4"
)

type CoreStore struct {
	ID         uint   `json:"id" form:"id"`
	UserID     uint   `gorm:"not null" json:"user_id" form:"user_id"`
	NamaToko   string `gorm:"type:string" json:"nama_toko" form:"nama_toko"`
	AlamatToko string `gorm:"type:string" json:"alamat_toko" form:"alamat_toko"`
	ImageToko  string `gorm:"type:string" json:"image_toko" form:"image_toko"`
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
	Photo(echo.Context) *uploader.UploadResult
}

type StoreServiceInterface interface {
	Create(input CoreStore) (uint, error)
	GetAll(userID uint) ([]CoreStore, error)
	GetById(StoreID uint, userID uint) (CoreStore, error)
	UpdateById(StoreID uint, userID uint, storeData CoreStore) error
	DeleteById(StoreID uint, userID uint) error
	Photo(echo.Context) *uploader.UploadResult
}
