package handler

import (
	admin "Laptop/features/admin"
	"Laptop/utils/responses"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AdminHandler struct {
	AdminService admin.AdminServiceInterface
}

func New(service admin.AdminServiceInterface) *AdminHandler {
	return &AdminHandler{
		AdminService: service,
	}
}

func (h *AdminHandler) Login(c echo.Context) error {
	userInput := new(LoginRequest)
	err := c.Bind(&userInput)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse(http.StatusBadRequest, "error bind data", nil))
	}

	dataLogin, token, err := h.AdminService.Login(userInput.Email, userInput.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse(http.StatusInternalServerError, err.Error(), nil))
	}
	response := MapCoreAdminToLogRes(dataLogin, token)
	return c.JSON(http.StatusOK, responses.WebResponse(http.StatusFound, "login successfully", response))
}

// func (h *UserHandler) CreateUser(c echo.Context) error {
// 	NewUser := new(UserRequest)

// 	// mendapatkan data yang dikirim oleh FE melalui request body
// 	err := c.Bind(&NewUser)
// 	if err != nil {
// 		return c.JSON(http.StatusBadRequest, responses.WebResponse(http.StatusBadRequest, "error bind data", nil))
// 	}
// 	//mapping dari struct request to struct core
// 	input := MapReqToCoreUser(*NewUser)
// 	_, err = h.userService.Create(input)
// 	if err != nil {
// 		if strings.Contains(err.Error(), "validation") {
// 			return c.JSON(http.StatusBadRequest, responses.WebResponse(http.StatusBadRequest, err.Error(), nil))
// 		} else {
// 			return c.JSON(http.StatusInternalServerError, responses.WebResponse(http.StatusInternalServerError, "error insert data, "+err.Error(), nil))

// 		}
// 	}
// 	return c.JSON(http.StatusOK, responses.WebResponse(http.StatusCreated, "success create user", nil))
// }

func (h *AdminHandler) GetAllUser(c echo.Context) error {
	result, err := h.AdminService.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse(http.StatusInternalServerError, "error read data, "+err.Error(), nil))
	}
	var usersResponse []AdminResponse
	for _, v := range result {
		usersResponse = append(usersResponse, MapCoreAdminToRes(v))
	}
	return c.JSON(http.StatusOK, responses.WebResponse(http.StatusFound, "success read data", usersResponse))
}

// func (h *UserHandler) GetUserById(c echo.Context) error {
// 	// Mengambil ID pengguna dari token JWT yang terkait dengan permintaan
// 	userID := middlewares.ExtractTokenUserId(c)

// 	// Memeriksa apakah ID pengguna yang diambil dari token sama dengan ID yang diminta
// 	result, err := h.userService.GetById(uint(userID))
// 	if err != nil {
// 		if strings.Contains(err.Error(), "validation") {
// 			return c.JSON(http.StatusBadRequest, responses.WebResponse(http.StatusBadRequest, err.Error(), nil))
// 		} else {
// 			return c.JSON(http.StatusInternalServerError, responses.WebResponse(http.StatusInternalServerError, "error read data, "+err.Error(), nil))
// 		}
// 	}

// 	resultResponse := UserResponse{
// 		Username:     result.Username,
// 		Name:         result.NamaLengkap,
// 		Email:        result.Email,
// 		NomorHP:      result.NomorHP,
// 		Alamat:       result.Alamat,
// 		JenisKelamin: result.JenisKelamin,
// 		ImageProfil:  result.ImageProfil,
// 		CreatedAt:    result.CreatedAt,
// 	}
// 	return c.JSON(http.StatusOK, responses.WebResponse(http.StatusFound, "success read data", resultResponse))
// }

// func (h *UserHandler) UpdateUserById(c echo.Context) error {
// 	userID := middlewares.ExtractTokenUserId(c)
// 	userInput := UserRequest{}
// 	errBind := c.Bind(&userInput)
// 	if errBind != nil {
// 		return c.JSON(http.StatusBadRequest, responses.WebResponse(http.StatusBadRequest, "error bind data", nil))
// 	}
// 	//Mapping user reques to core user
// 	Core := MapReqToCoreUser(userInput)
// 	err := h.userService.UpdateById(uint(userID), Core)
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, responses.WebResponse(http.StatusInternalServerError, "error update data, "+err.Error(), nil))
// 	}

// 	// Get user data for response
// 	user, err := h.userService.GetById(uint(userID))
// 	if err != nil {
// 		return c.JSON(http.StatusBadRequest, responses.WebResponse(http.StatusBadRequest, "user not found", nil))
// 	}
// 	resultResponse := MapCoreUserToRes(user)
// 	return c.JSON(http.StatusOK, responses.WebResponse(http.StatusOK, "user updated successfully", resultResponse))
// }

// func (h *UserHandler) DeleteUserById(c echo.Context) error {
// 	userID := middlewares.ExtractTokenUserId(c)
// 	err := h.userService.DeleteById(uint(userID))
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, responses.WebResponse(http.StatusInternalServerError, "error delete data, "+err.Error(), nil))
// 	}
// 	return c.JSON(http.StatusOK, responses.WebResponse(http.StatusOK, "success delete user", nil))
// }
