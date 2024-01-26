package handler

import (
	"Laptop/features/store"
	"time"
)

type StoreResponse struct {
	UserID     uint
	NamaToko   string    `json:"nama_toko" `
	AlamatToko string    `gorm:"type:string"json:"alamat_toko" `
	ImageToko  string    `gorm:"type:string"json:"image_toko"`
	CreatedAt  time.Time `json:"created_at"`
}

// Mapping CorePrject to TaskResponsee
func MapCoreStoreToStoreRes(core store.CoreStore) StoreResponse {
	return StoreResponse{
		UserID:     core.UserID,
		NamaToko:   core.NamaToko,
		AlamatToko: core.AlamatToko,
		ImageToko:  core.ImageToko,
		CreatedAt:  core.CreatedAt,
	}
}
