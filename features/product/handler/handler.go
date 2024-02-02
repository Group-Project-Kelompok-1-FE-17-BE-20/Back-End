package handler

import (
	"Laptop/app/middlewares"
	"Laptop/features/product"
	"Laptop/utils/responses"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ProductHandler struct {
	productService product.ProductServiceInterface
}

func New(service product.ProductServiceInterface) *ProductHandler {
	return &ProductHandler{
		productService: service,
	}
}

// insert product
func (handler *ProductHandler) CreateProduct(c echo.Context) error {
	// Mengambil ID pengguna dari token JWT yang terkait dengan permintaan
	userID := middlewares.ExtractTokenUserId(c)
	result, err := handler.productService.GetStoreID(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse(http.StatusInternalServerError, "error read data, "+err.Error(), nil))
	}

	inputPhoto, _ := c.FormFile("gambar")
	responURL := handler.productService.Photo(inputPhoto)
	log.Println(responURL.SecureURL)

	newProduct := ProductRequest{}
	newProduct.StoreID = result
	newProduct.Gambar = responURL.SecureURL

	errBind := c.Bind(&newProduct)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse(http.StatusBadRequest, "error bind data. data not valid", nil))
	}

	productCore := RequestToCore(newProduct)

	errCreate := handler.productService.Create(productCore)
	if errCreate != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse(http.StatusInternalServerError, "error insert data"+errCreate.Error(), nil))
	}

	return c.JSON(http.StatusOK, responses.WebResponse(http.StatusOK, "success insert data", nil))
}

// update product
func (handler *ProductHandler) UpdateProduct(c echo.Context) error {
	productID := c.Param("product_id")

	productID_int, errConv := strconv.Atoi(productID)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse(http.StatusBadRequest, "error convert id param", nil))
	}

	newUpdate := ProductRequest{}

	oldPrice := c.FormValue("price")
	if oldPrice != "" {
		newPrice, errConv := strconv.Atoi(oldPrice)
		if errConv != nil {
			return c.JSON(http.StatusBadRequest, responses.WebResponse(http.StatusBadRequest, "error convert data", nil))
		}
		newUpdate.Price = float64(newPrice)
	}

	oldStock := c.FormValue("stock")
	if oldStock != "" {
		newStock, errConv := strconv.Atoi(oldStock)
		if errConv != nil {
			return c.JSON(http.StatusBadRequest, responses.WebResponse(http.StatusBadRequest, "error convert data", nil))
		}
		newUpdate.Stock = newStock
	}

	oldPhoto, _ := c.FormFile("gambar")
	if oldPhoto != nil {
		responURL := handler.productService.Photo(oldPhoto)
		newUpdate.Gambar = responURL.SecureURL
	}

	newUpdate.Storage = c.FormValue("storage")
	newUpdate.RAM = c.FormValue("ram")
	newUpdate.Description = c.FormValue("description")
	newUpdate.Tipe = c.FormValue("model")
	newUpdate.Brand = c.FormValue("brand")
	newUpdate.Processor = c.FormValue("processor")
	newUpdate.Categories = c.FormValue("categories")

	fmt.Println("isi update: ", newUpdate)

	if errBind := c.Bind(&newUpdate); errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse(http.StatusBadRequest, "error bind data. data not valid", nil))
	}

	newUpdateCore := RequestToCore(newUpdate)

	errUpdates := handler.productService.Update(productID_int, newUpdateCore)
	if errUpdates != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse(http.StatusBadRequest, "error update data", nil))
	}

	return c.JSON(http.StatusOK, responses.WebResponse(http.StatusOK, "success update data", nil))
}

// delete product
func (handler *ProductHandler) Delete(c echo.Context) error {
	ProductID := c.Param("product_id")

	ProductID_int, errConv := strconv.Atoi(ProductID)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse(http.StatusBadRequest, "error convert id param", nil))
	}

	result, errRead := handler.productService.GetAll()
	if errRead != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse(http.StatusInternalServerError, "error read data. "+errRead.Error(), nil))
	}

	errDel := handler.productService.Delete(result, ProductID_int)
	if errDel != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse(http.StatusInternalServerError, "error delete data. "+errDel.Error(), nil))
	}

	return c.JSON(http.StatusOK, responses.WebResponse(http.StatusOK, "success delete data", nil))
}

// get all products
func (handler *ProductHandler) GetAllProducts(c echo.Context) error {
	result, errFind := handler.productService.GetAll()
	if errFind != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse(http.StatusInternalServerError, "error read data. "+errFind.Error(), nil))
	}

	return c.JSON(http.StatusOK, responses.WebResponse(http.StatusOK, "success read data.", result))
}

// read product by id
func (handler *ProductHandler) GetSingleProduct(c echo.Context) error {
	productID := c.Param("product_id")

	productID_int, errConv := strconv.Atoi(productID)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse(http.StatusBadRequest, "error convert id param", nil))
	}

	result, errFirst := handler.productService.GetSingle(productID_int)
	if errFirst != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse(http.StatusInternalServerError, "error read data. "+errFirst.Error(), nil))
	}

	return c.JSON(http.StatusOK, responses.WebResponse(http.StatusOK, "success read data.", result))
}

// read product by id
func (handler *ProductHandler) GetStoreProduct(c echo.Context) error {
	userID := middlewares.ExtractTokenUserId(c)
	storeID, err := handler.productService.GetStoreID(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse(http.StatusInternalServerError, "error read data, "+err.Error(), nil))
	}

	result, errFirst := handler.productService.GetStoreProducts(storeID)
	if errFirst != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse(http.StatusInternalServerError, "error read data. "+errFirst.Error(), nil))
	}

	return c.JSON(http.StatusOK, responses.WebResponse(http.StatusOK, "success read data.", result))
}
