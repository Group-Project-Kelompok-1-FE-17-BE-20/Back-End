package handler

import (
	"Laptop/features/store"
)

type StoreRequest struct {
	UserID     uint   `json:"user_id" form:"user_id"`
	NamaToko   string `gorm:"type:string" json:"nama_toko" form:"nama_toko"`
	AlamatToko string `gorm:"type:string" json:"alamat_toko" form:"alamat_toko"`
	ImageToko  string `gorm:"type:string" json:"image_toko" form:"image_toko"`
}

// Mapping dari struct TaskRequest To struct Core Task
func MapStoreReqToCoreStore(req StoreRequest) store.CoreStore {
	return store.CoreStore{
		UserID:     req.UserID,
		NamaToko:   req.NamaToko,
		AlamatToko: req.AlamatToko,
		ImageToko:  req.ImageToko,
	}
}
