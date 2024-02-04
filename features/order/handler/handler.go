package handler

import (
	config "Laptop/app/configs"
	"Laptop/app/database"
	"Laptop/app/middlewares"
	"Laptop/features/order"
	"Laptop/utils/responses"
	"fmt"
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
func (handler *OrderHandler) CreateOrder(c echo.Context) (order.Core, error) {
	// Mengambil ID pengguna dari token JWT yang terkait dengan permintaan
	userID := middlewares.ExtractTokenUserId(c)
	result, err := handler.orderService.GetCartID(userID)
	if err != nil {
		return order.Core{}, c.JSON(http.StatusInternalServerError, responses.WebResponse(http.StatusInternalServerError, "error read data, "+err.Error(), nil))
	}

	resGet, err := handler.orderService.GetAllCartItem(result)
	if err != nil {
		return order.Core{}, c.JSON(http.StatusInternalServerError, responses.WebResponse(http.StatusInternalServerError, "error read data, "+err.Error(), nil))
	}

	itemResponses := ResGetRequest(resGet)

	newOrder := OrderRequest{}
	newOrder.ShoppingCartID = result
	newOrder.Item = itemResponses
	newOrder.Status = "On Going"

	errBind := c.Bind(&newOrder)
	if errBind != nil {
		return order.Core{}, c.JSON(http.StatusBadRequest, responses.WebResponse(http.StatusBadRequest, "error bind data. data not valid", newOrder))
	}

	orderCore := RequestToCore(newOrder)

	errCreate := handler.orderService.Create(orderCore)
	if errCreate != nil {
		return order.Core{}, c.JSON(http.StatusInternalServerError, responses.WebResponse(http.StatusInternalServerError, "error insert data"+errCreate.Error(), nil))
	}

	return orderCore, c.JSON(http.StatusOK, responses.WebResponse(http.StatusOK, "success insert data", orderCore))
}

func (handler *OrderHandler) CreateOrderItem(c echo.Context) error {
	// Mengambil ID pengguna dari token JWT yang terkait dengan permintaan
	userID := middlewares.ExtractTokenUserId(c)

	// mendapatkan cart id
	cart_id, err := handler.orderService.GetCartID(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse(http.StatusInternalServerError, "error read data, "+err.Error(), nil))
	}

	order_id, err := handler.orderService.GetOrderID(cart_id)
	if order_id == 0 {
		// membuat order
		orderCore, errOr := handler.CreateOrder(c)
		if errOr != nil {
			return c.JSON(http.StatusInternalServerError, responses.WebResponse(http.StatusInternalServerError, "error make an order"+errOr.Error(), nil))
		}
		fmt.Println("data item: ", orderCore.ID, orderCore.Item)
	}

	// // mendapatkan order id dari order yang telah dibuat
	// orderID, err := handler.orderService.GetOrderID(cart_id)

	return c.JSON(http.StatusOK, responses.WebResponse(http.StatusOK, "success insert data", nil))
}

func (handler *OrderHandler) GetDetailOrder(c echo.Context) error {
	// Membuka koneksi ke database
	cfg := config.InitConfig()
	dbRaw := database.InitRawSql(cfg)

	userID := middlewares.ExtractTokenUserId(c)

	result, _, err := handler.orderService.DetailOrder(dbRaw, uint(userID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse(http.StatusInternalServerError, "error get data"+err.Error(), nil))
	}

	return c.JSON(http.StatusOK, responses.WebResponse(http.StatusOK, "success read data", result))
}

// func (handler *OrderHandler) CancelOrder(c echo.Context) error {
// 	history := HistoryRequest{}
// 	// Membuka koneksi ke database
// 	cfg := config.InitConfig()
// 	dbRaw := database.InitRawSql(cfg)

// 	userID := middlewares.ExtractTokenUserId(c)
// 	_, resID, err := handler.orderService.DetailOrder(dbRaw, uint(userID))
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, responses.WebResponse(http.StatusInternalServerError, "error get data"+err.Error(), nil))
// 	}

// 	date, err := handler.orderService.DateOrder(dbRaw, uint(userID))
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, responses.WebResponse(http.StatusInternalServerError, "error get data"+err.Error(), nil))
// 	}

// 	// cartID, err := handler.orderService.GetCartID(userID)
// 	// if err != nil {
// 	// 	return c.JSON(http.StatusInternalServerError, responses.WebResponse(http.StatusInternalServerError, "error read data, "+err.Error(), nil))
// 	// }

// 	// history.OrderID = resID
// 	// history.ShoppingCartID = cartID
// 	history.TglOrder = date
// 	// history.TotalBayar = 0.0
// 	// history.StatusOrder = "Cancel"

// 	historyCore := HistoryToCore(history)

// 	errDel := handler.orderService.Cancel(dbRaw, resID)
// 	if errDel != nil {
// 		return c.JSON(http.StatusInternalServerError, responses.WebResponse(http.StatusInternalServerError, "error cancel data. "+errDel.Error(), nil))
// 	}

// 	errCreate := handler.orderService.CreateHistory(historyCore)
// 	if errCreate != nil {
// 		return c.JSON(http.StatusInternalServerError, responses.WebResponse(http.StatusInternalServerError, "error insert data"+errCreate.Error(), nil))
// 	}

// 	return c.JSON(http.StatusOK, responses.WebResponse(http.StatusOK, "success cancel order", nil))
// }

func (handler *OrderHandler) OrderHistories(c echo.Context) error {
	userID := middlewares.ExtractTokenUserId(c)

	cfg := config.InitConfig()
	dbRaw := database.InitRawSql(cfg)

	allOrders, err := handler.orderService.GetAllPayments(dbRaw, userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse(http.StatusInternalServerError, "error read data, "+err.Error(), nil))
	}

	allOrdersResponse := CoreToResponseHistory(allOrders)

	return c.JSON(http.StatusOK, responses.WebResponse(http.StatusOK, "success read order histories", allOrdersResponse))
}
