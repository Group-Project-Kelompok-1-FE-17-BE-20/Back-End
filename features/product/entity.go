package product

import (
	"mime/multipart"
	"time"

	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

type Core struct {
	ID          uint      `json:"id" form:"id"`
	StoreID     uint      `gorm:"not null" json:"store_id" form:"store_id"`
	Storage     string    `gorm:"type:string" json:"storage" form:"storage"`
	RAM         string    `gorm:"type:string" json:"ram" form:"ram"`
	Price       float64   `gorm:"type:decimal(10,2)" json:"price" form:"price"`
	Description string    `gorm:"type:string" json:"description" form:"description"`
	Tipe        string    `gorm:"type:string" json:"model" form:"model"`
	Gambar      string    `gorm:"type:string" json:"gambar" form:"gambar" binding:"uri"`
	Brand       string    `gorm:"type:string" json:"brand" form:"brand"`
	Processor   string    `gorm:"type:string" json:"processor" form:"processor"`
	Categories  string    `gorm:"type:string" json:"categories" form:"categories"`
	Stock       int       `gorm:"type:integer" json:"stock" form:"stock"`
	CreatedAt   time.Time `json:"created_at" form:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" form:"updated_at"`
}

// interface untuk Data Layer
type ProductDataInterface interface {
	GetStoreID(userID uint) (uint, error)
	Photo(*multipart.FileHeader) *uploader.UploadResult
	Insert(input Core) error
	Update(id int, input Core) error
	SelectAll() ([]Core, error)
	Delete(input []Core, id int) error
	GetSingleProduct(productID_int int) (Core, error)
	GetStoreProducts(store_id uint) ([]Core, error)
}

// interface untuk Service Layer
type ProductServiceInterface interface {
	GetStoreID(userID uint) (uint, error)
	Photo(*multipart.FileHeader) *uploader.UploadResult
	Create(input Core) error
	Update(id int, input Core) error
	GetAll() ([]Core, error)
	Delete(input []Core, id int) error
	GetSingle(productID_int int) (Core, error)
	GetStoreProducts(store_id uint) ([]Core, error)
}
