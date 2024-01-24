package handler

import (
	"Laptop/app/middlewares"
	"Laptop/features/shoppingcartitem"
	"Laptop/utils/responses"
	"net/http"
	"strconv"

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
	// Mengambil ID pengguna dari token JWT yang terkait dengan permintaan
	userID := middlewares.ExtractTokenUserId(c)
	result, err := handler.itemService.GetCartID(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse(http.StatusInternalServerError, "error read data, "+err.Error(), nil))
	}

	productId := c.QueryParam("productId")
	intID, err := strconv.Atoi(productId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse(http.StatusInternalServerError, "error convert data, "+err.Error(), nil))
	}

	res, err := handler.itemService.GetPrice(uint(intID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse(http.StatusInternalServerError, "error read data, "+err.Error(), nil))
	}

	newItem := ItemRequest{}
	newItem.ShoppingCartID = result
	newItem.ProductID = uint(intID)
	newItem.UnitPrice = res
	newItem.TotalPrice = newItem.UnitPrice * float64(newItem.Quantity)

	errBind := c.Bind(&newItem)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse(http.StatusBadRequest, "error bind data. data not valid", newItem))
	}

	itemCore := RequestToCore(newItem)

	errCreate := handler.itemService.Create(itemCore)
	if errCreate != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse(http.StatusInternalServerError, "error insert data"+errCreate.Error(), nil))
	}

	return c.JSON(http.StatusOK, responses.WebResponse(http.StatusOK, "success insert data", newItem))
}
