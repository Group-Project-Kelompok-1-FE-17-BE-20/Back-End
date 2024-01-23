package handler

import (
	"Laptop/app/middlewares"
	"Laptop/features/shoppingcart"
	"Laptop/utils/responses"
	"net/http"

	"github.com/labstack/echo/v4"
)

type CartHandler struct {
	CartService shoppingcart.CartServiceInterface
}

func New(service shoppingcart.CartServiceInterface) *CartHandler {
	return &CartHandler{
		CartService: service,
	}
}

func (h *CartHandler) GetCart(c echo.Context) error {
	userID := middlewares.ExtractTokenUserId(c)

	newCart := CartRequest{}
	errBind := c.Bind(&newCart)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse(http.StatusBadRequest, "error bind data. data not valid", nil))
	}

	// Tambahkan userID ke objek CartRequest
	newCart.UserID = userID
	newCart.Status = "Kosong"

	cartCore := CartReqToCore(newCart)

	errCreate := h.CartService.Create(cartCore)
	if errCreate != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse(http.StatusInternalServerError, "error insert data"+errCreate.Error(), nil))
	}

	result, err := h.CartService.GetCart(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse(http.StatusInternalServerError, "error read data, "+err.Error(), nil))
	}

	resultResponse := MapCoreStoreToStoreRes(result)
	return c.JSON(http.StatusOK, responses.WebResponse(http.StatusFound, "success read data", resultResponse))
}
