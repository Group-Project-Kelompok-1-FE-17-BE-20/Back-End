package handler

import (
	"Laptop/app/middlewares"
	"Laptop/features/shoppingcartitem"
	"Laptop/utils/responses"
	"log"
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

func (handler *ItemHandler) GetNewCart(c echo.Context) error {
	userID := middlewares.ExtractTokenUserId(c)

	newCart := CartRequest{}
	errBind := c.Bind(&newCart)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse(http.StatusBadRequest, "error bind data. data not valid", nil))
	}

	// Tambahkan userID ke objek CartRequest
	newCart.UserID = userID
	newCart.Status = "On Going"

	cartCore := CartReqToCore(newCart)

	errCreate := handler.itemService.CreateCart(cartCore)
	if errCreate != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse(http.StatusInternalServerError, "error insert data"+errCreate.Error(), nil))
	}

	result, err := handler.itemService.GetCart(userID, newCart.Status)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse(http.StatusInternalServerError, "error read data, "+err.Error(), nil))
	}

	resultResponse := MapCoreStoreToStoreRes(result)
	return c.JSON(http.StatusOK, responses.WebResponse(http.StatusFound, "success create a cart", resultResponse))
}

// insert item
func (handler *ItemHandler) CreateItem(c echo.Context) error {
	// Mengambil ID pengguna dari token JWT yang terkait dengan permintaan
	userID := middlewares.ExtractTokenUserId(c)
	cartID, err := handler.itemService.GetCartID(userID)
	if cartID == 0 {
		errNew := handler.GetNewCart(c)
		if errNew != nil {
			return c.JSON(http.StatusInternalServerError, responses.WebResponse(http.StatusInternalServerError, "error make a new cart"+errNew.Error(), nil))
		}
	}

	cart_id, _ := handler.itemService.GetCartID(userID)

	productId := c.QueryParam("productId")
	product_int, err := strconv.Atoi(productId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse(http.StatusInternalServerError, "error convert data, "+err.Error(), nil))
	}

	res, err := handler.itemService.GetDataProduct(uint(product_int))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse(http.StatusInternalServerError, "error read data, "+err.Error(), nil))
	}

	qty := c.FormValue("quantity")
	qty_int, err := strconv.Atoi(qty)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse(http.StatusInternalServerError, "error convert data, "+err.Error(), nil))
	}

	newItem := ItemRequest{}
	newItem.ShoppingCartID = cart_id
	newItem.ProductID = uint(product_int)
	newItem.Tipe = res.Tipe
	newItem.Price = res.Price
	newItem.Processor = res.Processor
	newItem.RAM = res.RAM
	newItem.Storage = res.Storage
	newItem.TotalPrice = float64(qty_int) * newItem.Price
	newItem.Gambar = res.Gambar

	errBind := c.Bind(&newItem)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse(http.StatusBadRequest, "error bind data. data not valid", newItem))
	}
	log.Println(newItem)

	itemCore := RequestToCore(newItem)

	errCreate := handler.itemService.Create(itemCore)
	if errCreate != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse(http.StatusInternalServerError, "error insert data"+errCreate.Error(), nil))
	}

	return c.JSON(http.StatusOK, responses.WebResponse(http.StatusOK, "success insert data", newItem))
}

// update item
func (handler *ItemHandler) UpdateItem(c echo.Context) error {
	// mendapatkan productId
	productId := c.QueryParam("productId")
	intID, err := strconv.Atoi(productId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse(http.StatusInternalServerError, "error convert data, "+err.Error(), nil))
	}

	// Mendapatkan data item dari body request
	updatedItem := ItemRequest{}
	errBind := c.Bind(&updatedItem)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse(http.StatusBadRequest, "error bind data. data not valid", nil))
	}

	updatedItemCore := RequestToCore(updatedItem)

	errUpdate := handler.itemService.Update(uint(intID), updatedItemCore)
	if errUpdate != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse(http.StatusInternalServerError, "error updating project. "+errUpdate.Error(), nil))
	}

	return c.JSON(http.StatusOK, responses.WebResponse(http.StatusOK, "success updating item", nil))
}

// update item
func (handler *ItemHandler) DeleteItem(c echo.Context) error {
	// mendapatkan productId
	productId := c.QueryParam("productId")
	intID, err := strconv.Atoi(productId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse(http.StatusInternalServerError, "error convert data, "+err.Error(), nil))
	}

	result, errRead := handler.itemService.GetItemById(uint(intID))
	if errRead != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse(http.StatusInternalServerError, "error read data. "+errRead.Error(), nil))
	}

	errDelete := handler.itemService.Delete(result)
	if errDelete != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse(http.StatusInternalServerError, "error deleting project. "+errDelete.Error(), nil))
	}

	return c.JSON(http.StatusOK, responses.WebResponse(http.StatusOK, "success deleting item", nil))
}

// update item
func (handler *ItemHandler) GetItems(c echo.Context) error {
	userID := middlewares.ExtractTokenUserId(c)

	cart_id, _ := handler.itemService.GetCartID(userID)

	result, errRead := handler.itemService.GetCartItems(uint(cart_id))
	if errRead != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse(http.StatusInternalServerError, "error read data. "+errRead.Error(), nil))
	}

	return c.JSON(http.StatusOK, responses.WebResponse(http.StatusOK, "success read item", result))
}
