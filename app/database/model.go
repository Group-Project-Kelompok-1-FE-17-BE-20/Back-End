package database

import (
	"gorm.io/gorm"
)

// struct product gorm model
type Product struct {
	gorm.Model
	StoreID     uint    `gorm:"not null" json:"store_id" form:"store_id"`
	Storage     string  `gorm:"type:string" json:"storage" form:"storage"`
	RAM         string  `gorm:"type:string" json:"ram" form:"ram"`
	Price       float64 `gorm:"type:decimal(10,2)" json:"price" form:"price"`
	Description string  `gorm:"type:string" json:"description" form:"description"`
	Tipe        string  `gorm:"type:string" json:"model" form:"model"`
	Gambar      string  `gorm:"type:string" json:"gambar" form:"gambar" binding:"uri"`
	Brand       string  `gorm:"type:string" json:"brand" form:"brand"`
	Processor   string  `gorm:"type:string" json:"processor" form:"processor"`
	Categories  string  `gorm:"type:string" json:"categories" form:"categories"`
	Stock       int     `gorm:"type:integer" json:"stock" form:"stock"`
	Store       Store   `gorm:"foreignKey:StoreID" json:"store" form:"store"`
}

type Store struct {
	gorm.Model
	UserID     uint   `gorm:"column:user_id"`
	NamaToko   string `gorm:"column:nama_toko;"`
	AlamatToko string `gorm:"column:alamat_toko;"`
	ImageToko  string `gorm:"column:image_toko;"`
}

type User struct {
	gorm.Model
	Username     string `gorm:"column:username;"`
	NamaLengkap  string `gorm:"column:nama_lengkap;not null"`
	Email        string `gorm:"column:email;not null;unique"`
	Password     string `gorm:"column:password;not null"`
	NomorHP      string `gorm:"column:nomer_hp;"`
	Alamat       string `gorm:"column:alamat;"`
	JenisKelamin string `gorm:"column:jenis_kelamin;"`
	ImageProfil  string `gorm:"column:image_profil;"`
	Store        Store
}
