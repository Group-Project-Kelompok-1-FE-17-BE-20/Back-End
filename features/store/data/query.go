package data

import (
	"Laptop/app/database"
	store "Laptop/features/store"
	"errors"

	"gorm.io/gorm"
)

type StoreQuery struct {
	db *gorm.DB
}

func New(database *gorm.DB) store.StoreDataInterface {
	return &StoreQuery{
		db: database,
	}
}

// Insert implements task.TaskDataInterface.
func (r *StoreQuery) Insert(input store.CoreStore) (uint, error) {
	newStore := MapCoreStoreToStore(input)

	//simpan ke db
	tx := r.db.Create(&newStore)
	if tx.Error != nil {
		return 0, tx.Error
	}
	if tx.RowsAffected == 0 {
		return 0, errors.New("project not found")
	}
	return newStore.ID, nil
}
func (r *StoreQuery) SelectAll(userID uint) ([]store.CoreStore, error) {
	var dataStore []database.Store
	tx := r.db.Where("user_id", userID).Find(&dataStore)
	if tx.Error != nil {
		return nil, tx.Error
	}
	if tx.RowsAffected == 0 {
		return nil, errors.New("project not found")
	}
	//mapping Project Model to CoreProject
	coreProjectSlice := ListMapStoreToCoreStore(dataStore)
	return coreProjectSlice, nil
}

// // Select implements task.TaskDataInterface.
func (r *StoreQuery) Select(StoreID uint, userID uint) (store.CoreStore, error) {
	var storeData database.Store
	tx := r.db.Where("id = ? AND user_id = ?", StoreID, userID).First(&storeData)
	if tx.Error != nil {
		return store.CoreStore{}, tx.Error
	}
	// tx = r.db.Preload("Store").First(&storeData, StoreID)
	// if tx.Error != nil {
	// 	return storer.CoreStore{}, tx.Error
	// }
	if tx.RowsAffected == 0 {
		return store.CoreStore{}, errors.New("project not found")
	}
	//Mapping Project to CorePproject
	coreProject := MapStoreToCoreStore(storeData)
	return coreProject, nil
}

// func (r *StoreQuery) Select(userID uint) (store.CoreStore, error) {
// 	var storeData database.Store
// 	tx := r.db.Where("user_id = ?", userID).First(&storeData)
// 	if tx.Error != nil {
// 		return 0, tx.Error
// 	}

// 	storeID := storeData.ID
// 	return storeID., nil
// }

// Update implements task.TaskDataInterface.
func (r *StoreQuery) Update(StoreID uint, userID uint, storeData store.CoreStore) error {
	var Store database.Store
	tx := r.db.Where("id = ? AND user_id = ?", StoreID, userID).First(&Store)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("project not found")
	}

	//Mapping Project to CoreProject
	updatedProject := MapCoreStoreToStore(storeData)

	// Lakukan pembaruan data proyek dalam database
	tx = r.db.Model(&Store).Updates(updatedProject)
	if tx.Error != nil {
		return errors.New(tx.Error.Error() + " failed to update data")
	}
	return nil
}

// Delete implements task.TaskDataInterface.
func (r *StoreQuery) Delete(StoreID uint, userID uint) error {
	var Store database.Store
	tx := r.db.Where("id = ? AND user_id = ?", StoreID, userID).Delete(&Store)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("project not found")
	}
	return nil
}
