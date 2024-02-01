package handler

import (
	"Laptop/app/middlewares"
	user "Laptop/features/user"
	"fmt"
	"log"

	"Laptop/utils/responses"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userService user.UserServiceInterface
}

func New(service user.UserServiceInterface) *UserHandler {
	return &UserHandler{
		userService: service,
	}
}

func (h *UserHandler) Login(c echo.Context) error {
	userInput := new(LoginRequest)
	err := c.Bind(&userInput)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse(http.StatusBadRequest, "error bind data", nil))
	}

	dataLogin, token, err := h.userService.Login(userInput.Email, userInput.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse(http.StatusInternalServerError, err.Error(), nil))
	}
	response := MapCoreUserToLogRes(dataLogin, token)
	return c.JSON(http.StatusOK, responses.WebResponse(http.StatusFound, "login successfully", response))
}
func (h *UserHandler) CreateUser(c echo.Context) error {
	NewUser := new(UserRequest)

	// mendapatkan data yang dikirim oleh FE melalui request body
	err := c.Bind(&NewUser)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse(http.StatusBadRequest, "error bind data", nil))
	}
	//mapping dari struct request to struct core
	input := MapReqToCoreUser(*NewUser)
	_, err = h.userService.Create(input)
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return c.JSON(http.StatusBadRequest, responses.WebResponse(http.StatusBadRequest, err.Error(), nil))
		} else {
			return c.JSON(http.StatusInternalServerError, responses.WebResponse(http.StatusInternalServerError, "error insert data, "+err.Error(), nil))

		}
	}
	return c.JSON(http.StatusOK, responses.WebResponse(http.StatusCreated, "success create user", nil))
}

func (h *UserHandler) GetAllUser(c echo.Context) error {
	result, err := h.userService.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse(http.StatusInternalServerError, "error read data, "+err.Error(), nil))
	}
	var usersResponse []UserResponse
	for _, v := range result {
		usersResponse = append(usersResponse, MapCoreUserToRes(v))
	}
	return c.JSON(http.StatusOK, responses.WebResponse(http.StatusFound, "success read data", usersResponse))
}

func (h *UserHandler) GetUserById(c echo.Context) error {
	// Mengambil ID pengguna dari token JWT yang terkait dengan permintaan
	userID := middlewares.ExtractTokenUserId(c)

	// Memeriksa apakah ID pengguna yang diambil dari token sama dengan ID yang diminta
	result, err := h.userService.GetById(uint(userID))
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return c.JSON(http.StatusBadRequest, responses.WebResponse(http.StatusBadRequest, err.Error(), nil))
		} else {
			return c.JSON(http.StatusInternalServerError, responses.WebResponse(http.StatusInternalServerError, "error read data, "+err.Error(), nil))
		}
	}

	resultResponse := UserResponse{
		ID:           result.ID,
		Username:     result.Username,
		Name:         result.NamaLengkap,
		Email:        result.Email,
		NomorHP:      result.NomorHP,
		Alamat:       result.Alamat,
		JenisKelamin: result.JenisKelamin,
		ImageProfil:  result.ImageProfil,
		CreatedAt:    result.CreatedAt,
	}
	return c.JSON(http.StatusOK, responses.WebResponse(http.StatusFound, "success read data", resultResponse))
}

func (h *UserHandler) UpdateUserById(c echo.Context) error {
	userID := middlewares.ExtractTokenUserId(c)
	userInput := UserRequest{}
	///-------

	//newProduct := UserRequest{}

	// oldPrice := c.FormValue("price")
	// if oldPrice != "" {
	// 	newPrice, errConv := strconv.Atoi(oldPrice)
	// 	if errConv != nil {
	// 		return c.JSON(http.StatusBadRequest, responses.WebResponse(http.StatusBadRequest, "error convert data", nil))
	// 	}
	// 	newUpdate.Price = float64(newPrice)
	// }

	// oldStock := c.FormValue("stock")
	// if oldStock != "" {
	// 	newStock, errConv := strconv.Atoi(oldStock)
	// 	if errConv != nil {
	// 		return c.JSON(http.StatusBadRequest, responses.WebResponse(http.StatusBadRequest, "error convert data", nil))
	// 	}
	// 	newUpdate.Stock = newStock
	// }

	oldPhoto, _ := c.FormFile("image_profil")
	if oldPhoto != nil {
		responURL := h.userService.Photo(c)
		log.Println(responURL.SecureURL)
		userInput.ImageProfil = responURL.SecureURL
	}

	userInput.ID = userID
	userInput.Username = c.FormValue("username")
	userInput.Name = c.FormValue("nama_lengkap")
	userInput.Email = c.FormValue("email")
	userInput.Alamat = c.FormValue("alamat")
	userInput.JenisKelamin = c.FormValue("jenis_kelamin")
	userInput.NomorHP = c.FormValue("nomor_hp")
	userInput.Password = c.FormValue("password") // jika password tidak di isi maka ganti dengan password lama
	// newPassword := c.FormValue("password")

	// // Memeriksa apakah password baru diberikan
	// if newPassword != "" {
	// 	userInput.Password = newPassword
	// } else {
	// 	// Jika password baru tidak diberikan, biarkan nilai password lama
	// 	user, err := h.userService.GetById(uint(userID))
	// 	if err == nil {
	// 		userInput.Password = user.Password
	// 	}
	// }

	fmt.Println("isi update: ", userInput)

	///-------------------------------

	errBind := c.Bind(&userInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse(http.StatusBadRequest, "error bind data", nil))
	}

	//Mapping user reques to core user
	Core := MapReqToCoreUser(userInput)
	err := h.userService.UpdateById(uint(userID), Core)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse(http.StatusInternalServerError, "error update data, "+err.Error(), nil))
	}

	// Get user data for response
	user, err := h.userService.GetById(uint(userID))
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse(http.StatusBadRequest, "user not found", nil))
	}
	resultResponse := MapCoreUserToRes(user)
	return c.JSON(http.StatusOK, responses.WebResponse(http.StatusOK, "user updated successfully", resultResponse))
}

func (h *UserHandler) DeleteUserById(c echo.Context) error {
	userID := middlewares.ExtractTokenUserId(c)
	err := h.userService.DeleteById(uint(userID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse(http.StatusInternalServerError, "error delete data, "+err.Error(), nil))
	}
	return c.JSON(http.StatusOK, responses.WebResponse(http.StatusOK, "success delete user", nil))
}

// toko atribut

// func (handler *ProductHandler) GetSingleProduct(c echo.Context) error {
// 	productID := c.Param("product_id")

// 	productID_int, errConv := strconv.Atoi(productID)
// 	if errConv != nil {
// 		return c.JSON(http.StatusBadRequest, responses.WebResponse(http.StatusBadRequest, "error convert id param", nil))
// 	}

// 	result, errFirst := handler.productService.GetSingle(productID_int)
// 	if errFirst != nil {
// 		return c.JSON(http.StatusInternalServerError, responses.WebResponse(http.StatusInternalServerError, "error read data. "+errFirst.Error(), nil))
// 	}

// 	return c.JSON(http.StatusOK, responses.WebResponse(http.StatusOK, "success read data.", result))
// }
