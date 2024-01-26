package data

import (
	"Laptop/app/database"
	store "Laptop/features/store"
	"context"
	"errors"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/labstack/echo/v4"
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

// // // Select implements task.TaskDataInterface.
// func (r *StoreQuery) Select(StoreID uint, userID uint) (store.CoreStore, error) {
// 	var storeData database.Store
// 	tx := r.db.Where("id = ? AND user_id = ?", StoreID, userID).First(&storeData)
// 	if tx.Error != nil {
// 		return store.CoreStore{}, tx.Error
// 	}
// 	// tx = r.db.Preload("Store").First(&storeData, StoreID)
// 	// if tx.Error != nil {
// 	// 	return storer.CoreStore{}, tx.Error
// 	// }
// 	if tx.RowsAffected == 0 {
// 		return store.CoreStore{}, errors.New("project not found")
// 	}
// 	//Mapping Project to CorePproject
// 	coreProject := MapStoreToCoreStore(storeData)
// 	return coreProject, nil
// }

func (r *StoreQuery) Select(StoreID uint, userID uint) (store.CoreStore, error) {
	var storeData database.Store
	tx := r.db.Where("user_id = ?", userID).First(&storeData)
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
// func (r *StoreQuery) Update(StoreID uint, userID uint, storeData store.CoreStore) error {
// 	var Store database.Store
// 	tx := r.db.Where("id = ? AND user_id = ?", StoreID, userID).First(&Store)
// 	if tx.Error != nil {
// 		return tx.Error
// 	}
// 	if tx.RowsAffected == 0 {
// 		return errors.New("project not found")
// 	}

// 	//Mapping Project to CoreProject
// 	updatedProject := MapCoreStoreToStore(storeData)

// 	// Lakukan pembaruan data proyek dalam database
// 	tx = r.db.Model(&Store).Updates(updatedProject)
// 	if tx.Error != nil {
// 		return errors.New(tx.Error.Error() + " failed to update data")
// 	}
// 	return nil
// }

func (r *StoreQuery) Update(StoreID uint, userID uint, storeData store.CoreStore) error {
	var Store database.Store
	tx := r.db.Where("user_id = ?", userID).First(&Store)
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

// // Delete implements task.TaskDataInterface.
// func (r *StoreQuery) Delete(StoreID uint, userID uint) error {
// 	var Store database.Store
// 	tx := r.db.Where("id = ? AND user_id = ?", StoreID, userID).Delete(&Store)
// 	if tx.Error != nil {
// 		return tx.Error
// 	}
// 	if tx.RowsAffected == 0 {
// 		return errors.New("project not found")
// 	}
// 	return nil
// }

// Delete implements task.TaskDataInterface.
func (r *StoreQuery) Delete(StoreID uint, userID uint) error {
	var Store database.Store
	tx := r.db.Where("user_id = ?", userID).Delete(&Store)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("project not found")
	}
	return nil
}

func (r *StoreQuery) Photo(c echo.Context) *uploader.UploadResult {
	urlCloudinary := "cloudinary://979172954987629:PNgbXcjMn-VOd1AyTlN0yBSvnWU@dv3nso14b"
	fileHeader, _ := c.FormFile("image_toko")

	// var store database.Store
	// _ = c.Bind(&store)
	file, _ := fileHeader.Open()
	//log.Println(fileHeader.Filename)

	ctx := context.Background()
	cldService, _ := cloudinary.NewFromURL(urlCloudinary)
	resp, _ := cldService.Upload.Upload(ctx, file, uploader.UploadParams{})
	//log.Println(resp.SecureURL)

	return resp
}

//
// func (r *StoreQuery) Photo(c echo.Context) error {
// 	urlCloudinary := "cloudinary://979172954987629:PNgbXcjMn-VOd1AyTlN0yBSvnWU@dv3nso14b"
// 	fileHeader, err := c.FormFile("ImageToko")

// 	if err != nil {
// 		// Handle error
// 		return err
// 	}

// 	var store database.Store
// 	if err := c.Bind(&store); err != nil {
// 		// Handle error
// 		return err
// 	}

// 	if fileHeader == nil {
// 		// Foto tidak ditemukan atau tidak valid
// 		return errors.New("Invalid or missing file header")
// 	}

// 	file, err := fileHeader.Open()
// 	if err != nil {
// 		// Handle error
// 		return err
// 	}

// 	// Lanjutkan dengan operasi membuka file

// 	ctx := context.Background()
// 	cldService, err := cloudinary.NewFromURL(urlCloudinary)
// 	if err != nil {
// 		// Handle error
// 		return err
// 	}

// 	resp, err := cldService.Upload.Upload(ctx, file, uploader.UploadParams{})
// 	if err != nil {
// 		// Handle error
// 		return err
// 	}

// 	//log.Println(resp.SecureURL)
// 	return c.JSON(http.StatusOK, resp)
// }
