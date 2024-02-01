package handler

import (
	"Laptop/app/middlewares"
	"Laptop/features/store"
	"Laptop/utils/responses"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type StoreHandler struct {
	StoreService store.StoreServiceInterface
}

func New(service store.StoreServiceInterface) *StoreHandler {
	return &StoreHandler{
		StoreService: service,
	}
}
func (h *StoreHandler) CreateStore(c echo.Context) error {
	newStore := new(StoreRequest)
	// Mengambil ID pengguna dari token JWT yang terkait dengan permintaan
	newStore.UserID = middlewares.ExtractTokenUserId(c)
	//mendapatkan data yang dikirim oleh FE melalui request
	err := c.Bind(&newStore)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse(http.StatusBadRequest, "error bind data", nil))
	}

	//mapping dari request to CoreProject
	input := MapStoreReqToCoreStore(*newStore)
	_, err = h.StoreService.Create(input)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse(http.StatusInternalServerError, "error insert data, "+err.Error(), nil))
	}
	return c.JSON(http.StatusOK, responses.WebResponse(http.StatusCreated, "success create project", nil))
}

func (h *StoreHandler) GetAllStore(c echo.Context) error {
	// Mengambil ID pengguna dari token JWT yang terkait dengan permintaan
	userID := middlewares.ExtractTokenUserId(c)
	result, err := h.StoreService.GetAll(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse(http.StatusInternalServerError, "error read data, "+err.Error(), nil))
	}
	var storeResponse []StoreResponse
	for _, v := range result {
		storeResponse = append(storeResponse, MapCoreStoreToStoreRes(v))
	}
	return c.JSON(http.StatusOK, responses.WebResponse(http.StatusFound, "success read data", storeResponse))
}

func (h *StoreHandler) GetStoreById(c echo.Context) error {
	idParam := c.Param("user_id")
	idConv, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse(http.StatusBadRequest, "user id is not valid", nil))
	}
	// Mengambil ID pengguna dari token JWT yang terkait dengan permintaan
	userID := middlewares.ExtractTokenUserId(c)
	result, err := h.StoreService.GetById(uint(idConv), userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse(http.StatusInternalServerError, "error read data, "+err.Error(), nil))
	}

	resultResponse := MapCoreStoreToStoreRes(result)
	return c.JSON(http.StatusOK, responses.WebResponse(http.StatusFound, "success read data", resultResponse))
}

func (h *StoreHandler) UpdateStoreById(c echo.Context) error {
	idParam := c.Param("store_id")
	idConv, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse(http.StatusBadRequest, "user id is not valid", nil))
	}
	// Mengambil ID pengguna dari token JWT yang terkait dengan permintaan
	userID := middlewares.ExtractTokenUserId(c)
	StoreInput := StoreRequest{}
	errBind := c.Bind(&StoreInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse(http.StatusBadRequest, "error bind data", nil))
	}

	//StoreInput := StoreRequest{}
	StoreInput.UserID = userID
	StoreInput.NamaToko = c.FormValue("nama_toko")
	StoreInput.AlamatToko = c.FormValue("alamat_toko")

	oldPhoto, _ := c.FormFile("image_toko")
	if oldPhoto != nil {
		responURL := h.StoreService.Photo(c)
		log.Println(responURL.SecureURL)
		StoreInput.ImageToko = responURL.SecureURL
	}

	//Mapping store reques to core task
	Core := MapStoreReqToCoreStore(StoreInput)
	err = h.StoreService.UpdateById(uint(idConv), userID, Core)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse(http.StatusInternalServerError, "error update data, "+err.Error(), nil))
	}

	// Get store data for response
	updatedStore, err := h.StoreService.GetById(uint(idConv), userID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse(http.StatusBadRequest, "task not found", nil))
	}
	resultResponse := MapCoreStoreToStoreRes(updatedStore)
	return c.JSON(http.StatusOK, responses.WebResponse(http.StatusOK, "store updated successfully", resultResponse))
}

func (h *StoreHandler) DeleteStoreById(c echo.Context) error {
	// Mengambil ID pengguna dari token JWT yang terkait dengan permintaan
	userID := middlewares.ExtractTokenUserId(c)
	if userID == 0 {
		return c.JSON(http.StatusUnauthorized, responses.WebResponse(http.StatusUnauthorized, "Invalid user token", nil))
	}
	idParam := c.Param("store_id")
	idConv, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse(http.StatusBadRequest, "user id is not valid", nil))
	}
	fmt.Printf("idParam: %s\n", idParam)

	err = h.StoreService.DeleteById(uint(idConv), userID)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse(http.StatusInternalServerError, "error delete data, "+err.Error(), nil))
	}
	fmt.Printf("idConv: %d\n", idConv)
	return c.JSON(http.StatusOK, responses.WebResponse(http.StatusOK, "success delete task", nil))
}
