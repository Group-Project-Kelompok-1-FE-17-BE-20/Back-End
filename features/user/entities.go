package user

import (
	"time"

	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type CoreUser struct {
	gorm.Model
	ID           uint
	Username     string `json:"username" form:"username"`
	NamaLengkap  string `json:"nama_lengkap" form:"nama_lengkap"`
	Email        string `json:"email" form:"email" validate:"required" gorm:"unique"`
	Password     string `json:"password" form:"Password" validate:"required"`
	NomorHP      string `json:"NomorHP" form:"NomorHP"`
	Alamat       string `json:"Alamat" form:"Alamat"`
	JenisKelamin string `json:"JenisKelamin" form:"JenisKelamin"`
	ImageProfil  string `json:"image_profil" form:"image_profil"`
	// NamaToko     string `json:"NamaToko" form:"NamaToko"`
	// AlamatToko   string `json:"AlamatToko" form:"AlamatToko"`
	// ImageToko    string `json:"ImageToko" form:"ImageToko"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type UserDataInterface interface {
	Login(email string) (CoreUser, error)
	Insert(inputUser CoreUser) (uint, error)
	SelectAll() ([]CoreUser, error)
	Select(userId uint) (CoreUser, error)
	Update(userId uint, userData CoreUser) error
	Delete(userId uint) error
	Photo(echo.Context) *uploader.UploadResult
}

type UserServiceInterface interface {
	Login(email string, password string) (CoreUser, string, error)
	Create(inputUser CoreUser) (uint, error)
	GetAll() ([]CoreUser, error)
	GetById(userId uint) (CoreUser, error)
	UpdateById(userId uint, userData CoreUser) error
	DeleteById(userId uint) error
	Photo(echo.Context) *uploader.UploadResult
}
