package database

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username     string `gorm:"not null" json:"username" form:"username"`
	NamaLengkap  string `gorm:"not null" json:"nama_lengkap" form:"nama_lengkap"`
	Email        string `gorm:"not null;unique" json:"email" form:"email"`
	Password     string `gorm:"not null" json:"password" form:"password"`
	NomorHP      string `gorm:"type:string" json:"nomor_hp" form:"nomor_hp"`
	Alamat       string `gorm:"type:string" json:"alamat" form:"alamat"`
	JenisKelamin string `gorm:"type:string" json:"jenis_kelamin" form:"jenis_kelamin"`
	ImageProfil  string `gorm:"type:string" json:"image_profil" form:"image_profil" binding:"uri"`
	Store        Store
	ShoppingCart ShoppingCart
}

type Store struct {
	gorm.Model
	UserID     uint   `gorm:"type:string" json:"user_id" form:"user_id"`
	NamaToko   string `gorm:"type:string" json:"nama_toko" form:"nama_toko"`
	AlamatToko string `gorm:"type:string" json:"alamat_toko" form:"alamat_toko"`
	ImageToko  string `gorm:"type:string" json:"image_toko" form:"image_toko" binding:"uri"`
}

type Product struct {
	gorm.Model
	StoreID          uint    `gorm:"not null" json:"store_id" form:"store_id"`
	Storage          string  `gorm:"type:string" json:"storage" form:"storage"`
	RAM              string  `gorm:"type:string" json:"ram" form:"ram"`
	Price            float64 `gorm:"type:decimal(10,2)" json:"price" form:"price"`
	Description      string  `gorm:"type:string" json:"description" form:"description"`
	Tipe             string  `gorm:"type:string" json:"model" form:"model"`
	Gambar           string  `gorm:"type:string" json:"gambar" form:"gambar" binding:"uri"`
	Brand            string  `gorm:"type:string" json:"brand" form:"brand"`
	Processor        string  `gorm:"type:string" json:"processor" form:"processor"`
	Categories       string  `gorm:"type:string" json:"categories" form:"categories"`
	Stock            int     `gorm:"type:integer" json:"stock" form:"stock"`
	Store            Store
	ShoppingCartItem ShoppingCartItem
}

type ShoppingCart struct {
	gorm.Model
	UserID       uint   `gorm:"column:user_id"`
	Status       string `gorm:"type:string" json:"status" form:"status"`
	Order        Order
	OrderHistory OrderHistory
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

type Order struct {
	gorm.Model
	ShoppingCartID uint        `gorm:"not null" json:"cartId" form:"cartId"`
	Item           []OrderItem `gorm:"foreignKey:OrderID"`
	Status         string      `gorm:"not null" json:"status" form:"status"`
	OrderHistory   OrderHistory
}

type OrderItem struct {
	gorm.Model
	OrderID     uint    `gorm:"not null" json:"orderId" form:"orderId"`
	Productid   uint    `gorm:"not null" json:"prod_id" form:"prod_id"`
	Jumlah      uint    `gorm:"not null" json:"jumlah" form:"jumlah"`
	TotalAmount float64 `gorm:"not null" json:"totalAmount" form:"totalAmount"`
}

type OrderHistory struct {
	gorm.Model
	OrderID        uint      `gorm:"not null" json:"orderId" form:"orderId"`
	ShoppingCartID uint      `gorm:"not null" json:"cartId" form:"cartId"`
	TglOrder       time.Time `gorm:"not null" json:"date_order" form:"date_order"`
	TotalBayar     float64   `gorm:"not null" json:"total" form:"total"`
	StatusOrder    string    `gorm:"not null" json:"status_order" form:"status_order"`
}

type Admin struct {
	gorm.Model
	UserID   uint64 `gorm:"user_id"`
	Email    string `gorm:"column:email;not null;unique"`
	Password string `gorm:"column:password;not null"`
	// Users    User   `gorm:"foreignKey:AdminID"`
	// Stores       []data.Store `
}
