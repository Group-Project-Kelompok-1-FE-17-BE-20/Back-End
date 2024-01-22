package data

import (
	"Laptop/app/database"
	"Laptop/features/store"
)

// Mapping CoreTask to Task Model
func MapCoreStoreToStore(core store.CoreStore) database.Store {
	return database.Store{
		NamaToko:   core.NamaToko,
		UserID:     core.UserID,
		AlamatToko: core.AlamatToko,
		ImageToko:  core.ImageToko,
	}
}

func MapStoreToCoreStore(model database.Store) store.CoreStore {
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
func ListMapStoreToCoreStore(models []database.Store) []store.CoreStore {
	var core []store.CoreStore
	for _, model := range models {
		core = append(core, MapStoreToCoreStore(model))
	}
	return core
}
