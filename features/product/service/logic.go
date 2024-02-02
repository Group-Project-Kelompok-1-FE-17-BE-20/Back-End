package service

import (
	"Laptop/features/product"
	"errors"
	"mime/multipart"

	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

type productService struct {
	productData product.ProductDataInterface
}

// dependency injection
func New(repo product.ProductDataInterface) product.ProductServiceInterface {
	return &productService{
		productData: repo,
	}
}

func (service *productService) GetStoreID(input uint) (uint, error) {
	// logic validation
	res, err := service.productData.GetStoreID(input)
	return res, err
}

func (service *productService) Photo(fileHeader *multipart.FileHeader) *uploader.UploadResult {
	res := service.productData.Photo(fileHeader)
	return res
}

func (service *productService) Create(input product.Core) error {
	// logic validation
	err := service.productData.Insert(input)
	return err
}

func (service *productService) Update(id int, input product.Core) error {
	//validasi
	if id <= 0 {
		return errors.New("invalid id")
	}
	//validasi inputan
	// ...
	err := service.productData.Update(id, input)
	return err
}

func (service *productService) GetAll() ([]product.Core, error) {
	// logic
	// memanggil func yg ada di data layer
	results, err := service.productData.SelectAll()
	return results, err
}

func (service *productService) Delete(input []product.Core, id int) error {
	err := service.productData.Delete(input, id)
	return err
}

func (service *productService) GetSingle(productID_int int) (product.Core, error) {
	// logic
	// memanggil func yg ada di data layer
	results, err := service.productData.GetSingleProduct(productID_int)
	return results, err
}

func (service *productService) GetStoreProducts(store_id uint) ([]product.Core, error) {
	// logic
	// memanggil func yg ada di data layer
	results, err := service.productData.GetStoreProducts(store_id)
	return results, err
}
