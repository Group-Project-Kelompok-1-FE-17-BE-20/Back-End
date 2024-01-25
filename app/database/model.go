package database

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username     string `gorm:"not null" json:"username" form:"username"`
	NamaLengkap  string `gorm:"not null" json:"nama_lengkap" form:"nama_lengkap"`
	Email        string `gorm:"not null;unique" json:"email" form:"email"`
	Password     string `gorm:"not null" json:"password" form:"password"`
	NomorHP      string `json:"NomorHP" form:"NomorHP"`
	Alamat       string `json:"alamat" form:"alamat"`
	JenisKelamin string `json:"jenis_kelamin" form:"jenis_kelamin"`
	ImageProfil  string `json:"image_profil" form:"image_profil"`
	Store        Store
	ShoppingCart ShoppingCart
}

type Store struct {
	gorm.Model
	UserID     uint   `gorm:"not null" json:"user_id" form:"user_id"`
	NamaToko   string `gorm:"not null" json:"nama_toko" form:"nama_toko"`
	AlamatToko string `gorm:"not null" json:"alamat_toko" form:"alamat_toko"`
	ImageToko  string `json:"image_toko" form:"image_toko"`
}

type Product struct {
	gorm.Model
	StoreID          uint    `gorm:"not null" json:"store_id" form:"store_id"`
	Storage          string  `gorm:"type:string" json:"storage" form:"storage"`
	RAM              string  `gorm:"type:string" json:"ram" form:"ram"`
	Price            float64 `gorm:"type:decimal(10,2)" json:"price" form:"price"`
	Description      string  `gorm:"type:string" json:"description" form:"description"`
	Tipe             string  `gorm:"type:string" json:"model" form:"model"`
	Gambar           string  `json:"image" form:"image"`
	Brand            string  `gorm:"type:string" json:"brand" form:"brand"`
	Processor        string  `gorm:"type:string" json:"processor" form:"processor"`
	Categories       string  `gorm:"type:string" json:"categories" form:"categories"`
	Stock            int     `gorm:"type:integer" json:"stock" form:"stock"`
	Store            Store
	ShoppingCartItem ShoppingCartItem
}

type ShoppingCart struct {
	gorm.Model
	UserID uint   `gorm:"column:user_id"`
	Status string `gorm:"type:string" json:"status" form:"status"`
}

type ShoppingCartItem struct {
	gorm.Model
	ShoppingCartID uint    `gorm:"not null" json:"cartId" form:"cartId"`
	ProductID      uint    `gorm:"not null" json:"productId" form:"productId"`
	Quantity       uint    `gorm:"not null" json:"quantity" form:"quantity"`
	UnitPrice      float64 `gorm:"not null" json:"unitPrice" form:"unitPrice"`
	TotalPrice     float64 `gorm:"not null" json:"totalPrice" form:"totalPrice"`
	ShoppingCart   ShoppingCart
}

type Admin struct {
	gorm.Model
	UserID   uint64 `gorm:"user_id"`
	Email    string `gorm:"column:email;not null;unique"`
	Password string `gorm:"column:password;not null"`
	// Users    User   `gorm:"foreignKey:AdminID"`
	// Stores       []data.Store `
}
