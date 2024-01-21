package data

import (
	"Laptop/features/store"

	"gorm.io/gorm"
)

type Store struct {
	gorm.Model
	UserID     uint   `gorm:"column:user_id"`
	NamaToko   string `gorm:"column:nama_toko;"`
	AlamatToko string `gorm:"column:alamat_toko;"`
	ImageToko  string `gorm:"column:image_toko;"`
}

// Mapping CoreTask to Task Model
func MapCoreStoreToStore(core store.CoreStore) Store {
	return Store{
		NamaToko:   core.NamaToko,
		UserID:     core.UserID,
		AlamatToko: core.AlamatToko,
		ImageToko:  core.ImageToko,
	}
}

func MapStoreToCoreStore(model Store) store.CoreStore {
	return store.CoreStore{
		ID:         model.ID,
		UserID:     model.UserID,
		NamaToko:   model.NamaToko,
		AlamatToko: model.AlamatToko,
		ImageToko:  model.ImageToko,
		CreatedAt:  model.CreatedAt,
		UpdatedAt:  model.UpdatedAt,
		DeletedAt:  model.DeletedAt.Time,
	}
}

// mapping Task Model to CoreTask
func ListMapStoreToCoreStore(models []Store) []store.CoreStore {
	var core []store.CoreStore
	for _, model := range models {
		core = append(core, MapStoreToCoreStore(model))
	}
	return core
}
