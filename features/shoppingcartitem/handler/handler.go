package handler

import (
	"Laptop/features/shoppingcartitem"
	"Laptop/utils/responses"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ItemHandler struct {
	itemService shoppingcartitem.ItemServiceInterface
}

func New(service shoppingcartitem.ItemServiceInterface) *ItemHandler {
	return &ItemHandler{
		itemService: service,
	}
}

// insert item
func (handler *ItemHandler) CreateItem(c echo.Context) error {
	newItem := ItemRequest{}
	errBind := c.Bind(&newItem)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse(http.StatusBadRequest, "error bind data. data not valid", nil))
	}

	itemCore := RequestToCore(newItem)

	errCreate := handler.itemService.Create(itemCore)
	if errCreate != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse(http.StatusInternalServerError, "error insert data"+errCreate.Error(), nil))
	}

	return c.JSON(http.StatusOK, responses.WebResponse(http.StatusOK, "success insert data", nil))
}
