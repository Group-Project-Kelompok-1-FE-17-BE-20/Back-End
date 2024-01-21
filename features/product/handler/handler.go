package handler

import (
	"Laptop/features/product"
	"Laptop/utils/responses"
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
	newProduct := ProductRequest{}
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

	var newUpdate product.Core
	if errBind := c.Bind(&newUpdate); errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse(http.StatusBadRequest, "error bind data. data not valid", nil))
	}

	errUpdates := handler.productService.Update(productID_int, newUpdate)
	if errUpdates != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse(http.StatusBadRequest, "error update data", nil))
	}

	return c.JSON(http.StatusOK, responses.WebResponse(http.StatusOK, "success update data", nil))
}
