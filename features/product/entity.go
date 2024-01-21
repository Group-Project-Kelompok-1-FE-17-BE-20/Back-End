package product

import (
	"time"
)

type Core struct {
	ID          uint      `json:"id" form:"id"`
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
	Insert(input Core) error
}

// interface untuk Service Layer
type ProductServiceInterface interface {
	Create(input Core) error
}
