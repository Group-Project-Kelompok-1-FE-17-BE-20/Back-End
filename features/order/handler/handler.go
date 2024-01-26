package handler

import (
	"Laptop/app/middlewares"
	"Laptop/features/order"
	"Laptop/utils/responses"
	"net/http"

	"github.com/labstack/echo/v4"
)

type OrderHandler struct {
	orderService order.OrderServiceInterface
}

func New(service order.OrderServiceInterface) *OrderHandler {
	return &OrderHandler{
		orderService: service,
	}
}

// insert order
func (handler *OrderHandler) CreateOrder(c echo.Context) error {
	// Mengambil ID pengguna dari token JWT yang terkait dengan permintaan
	userID := middlewares.ExtractTokenUserId(c)
	result, err := handler.orderService.GetCartID(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse(http.StatusInternalServerError, "error read data, "+err.Error(), nil))
	}

	resGet, err := handler.orderService.GetAllCartItem(result)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse(http.StatusInternalServerError, "error read data, "+err.Error(), nil))
	}

	itemResponses := ResGetRequest(resGet)

	newOrder := OrderRequest{}
	newOrder.ShoppingCartID = result
	newOrder.Item = itemResponses
	newOrder.Status = " "

	errBind := c.Bind(&newOrder)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse(http.StatusBadRequest, "error bind data. data not valid", newOrder))
	}

	orderCore := RequestToCore(newOrder)

	errCreate := handler.orderService.Create(orderCore)
	if errCreate != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse(http.StatusInternalServerError, "error insert data"+errCreate.Error(), nil))
	}

	return c.JSON(http.StatusOK, responses.WebResponse(http.StatusOK, "success insert data", orderCore))
}
