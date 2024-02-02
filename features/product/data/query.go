package data

import (
	"Laptop/app/database"
	"Laptop/features/product"
	"context"
	"errors"
	"mime/multipart"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"gorm.io/gorm"
)

type productQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) product.ProductDataInterface {
	return &productQuery{
		db: db,
	}
}

func (repo *productQuery) GetStoreID(userID uint) (uint, error) {
	var storeData database.Store
	tx := repo.db.Where("user_id = ?", userID).First(&storeData)
	if tx.Error != nil {
		return 0, tx.Error
	}

	storeID := storeData.ID
	return storeID, nil
}

func (repo *productQuery) Photo(fileHeader *multipart.FileHeader) *uploader.UploadResult {
	urlCloudinary := "cloudinary://377166738273893:ga3Zq7Ts84gJ-Ltn-gyMkTgHd40@dltcy9ghn"

	file, _ := fileHeader.Open()
	//log.Println(fileHeader.Filename)

	ctx := context.Background()
	cldService, _ := cloudinary.NewFromURL(urlCloudinary)
	resp, _ := cldService.Upload.Upload(ctx, file, uploader.UploadParams{})
	//log.Println(resp.SecureURL)

	return resp
}

func (repo *productQuery) Insert(input product.Core) error {
	// simpan ke DB
	newProductGorm := CoreToModel(input)

	tx := repo.db.Create(&newProductGorm) // proses query insert
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (repo *productQuery) Update(idParam int, newUpdate product.Core) error {
	newUpdateGorm := CoreToModel(newUpdate)

	txUpdates := repo.db.Model(&database.Product{}).Where("id = ?", idParam).Updates(newUpdateGorm)
	if txUpdates.Error != nil {
		return txUpdates.Error
	}

	return nil
}

func (repo *productQuery) Delete(input []product.Core, id int) error {
	allProductGorm := CoretoModelGorm(input)

	txDel := repo.db.Delete(&allProductGorm, id)
	if txDel.Error != nil {
		return txDel.Error
	}

	if txDel.RowsAffected == 0 {
		return errors.New("user's not found")
	}

	return nil
}

func (repo *productQuery) SelectAll() ([]product.Core, error) {
	var productsDataGorm []database.Product
	tx := repo.db.Find(&productsDataGorm) // select * from users;
	if tx.Error != nil {
		return nil, tx.Error
	}

	//mapping
	allProductCore := ModelGormToCore(productsDataGorm)

	return allProductCore, nil
}

func (repo *productQuery) GetSingleProduct(productID_int int) (product.Core, error) {
	var singleProductGorm database.Product
	tx := repo.db.First(&singleProductGorm, productID_int)
	if tx.Error != nil {
		return product.Core{}, tx.Error
	}

	singleProductCore := ModelToCore(singleProductGorm)

	return singleProductCore, nil
}

func (repo *productQuery) GetStoreProducts(store_id uint) ([]product.Core, error) {
	var productsDataGorm []database.Product
	tx := repo.db.Where("store_id = ?", store_id).Find(&productsDataGorm) // select * from users;
	if tx.Error != nil {
		return nil, tx.Error
	}

	//mapping
	allProductCore := ModelGormToCore(productsDataGorm)

	return allProductCore, nil
}
