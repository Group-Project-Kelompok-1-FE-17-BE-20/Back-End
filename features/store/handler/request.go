package handler

import (
	"Laptop/features/store"
)

type StoreRequest struct {
	UserID     uint `json:"user_id" form:"user_id"`
	NamaToko   string
	AlamatToko string
	ImageToko  string
}

// Mapping dari struct TaskRequest To struct Core Task
func MapStoreReqToCoreStore(req StoreRequest) store.CoreStore {
	return store.CoreStore{
		UserID:     req.UserID,
		NamaToko:   req.NamaToko,
		AlamatToko: req.AlamatToko,
	}
}
