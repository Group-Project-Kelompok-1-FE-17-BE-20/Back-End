package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	config "Laptop/app/configs"
	"Laptop/app/database"
	"Laptop/app/middlewares"
	"Laptop/features/payment"
	"Laptop/utils/responses"

	"github.com/labstack/echo/v4"
)

var log = middlewares.Log()

type paymentHandler struct {
	service payment.PaymentService
}

// type ErrorResponse struct {
// 	Message string `json:"message"`
// }

func New(us payment.PaymentService) payment.PaymentHandler {
	return &paymentHandler{
		service: us,
	}
}

func (tc *paymentHandler) Payment() echo.HandlerFunc {
	return func(c echo.Context) error {
		// Membuka koneksi ke database
		cfg := config.InitConfig()
		dbRaw := database.InitRawSql(cfg)

		//result, err := middlewares.ExtractTokenUserId(c)
		userID, err := middlewares.ExtractToken(c)
		if err != nil {
			log.Error("missing or malformed JWT")
			return c.JSON(http.StatusUnauthorized, responses.ResponseFormat(http.StatusUnauthorized, "", "Missing or Malformed JWT", nil, nil))
		}

		orderID, amount := tc.service.GetOrderItems(dbRaw, uint(userID))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, responses.WebResponse(http.StatusInternalServerError, "error get data"+err.Error(), nil))
		}

		order_id := strconv.Itoa(int(orderID))
		amountString := strconv.FormatFloat(amount, 'f', -1, 64)

		// mendapatkan data dari form data
		request := createPaymentRequest{}
		request.OrderID = order_id
		request.Amount = amountString

		errBind := c.Bind(&request)
		if errBind != nil {
			log.Error("error on bind request")
			return c.JSON(http.StatusBadRequest, responses.ResponseFormat(http.StatusBadRequest, "", "Bad request"+errBind.Error(), nil, nil))
		}

		fmt.Printf("log: %v\n", request)

		payment, err := tc.service.Payment(RequestToCore(request))
		if err != nil {
			if strings.Contains(err.Error(), "unsupported bank account") {
				return c.JSON(http.StatusBadRequest, responses.ResponseFormat(http.StatusBadRequest, "", "Bad request, unsupported bank account", nil, nil))
			}
			return c.JSON(http.StatusInternalServerError, responses.ResponseFormat(http.StatusInternalServerError, "", "Internal server error", nil, nil))
		}

		log.Sugar().Infoln(payment)

		return c.JSON(http.StatusOK, responses.ResponseFormat(http.StatusOK, "", "Successful Operation", paymentResp(payment), nil))

	}

}

func (tc *paymentHandler) Notification() echo.HandlerFunc {
	return func(c echo.Context) error {
		updateData := midtransCallback{}
		errBind := c.Bind(&updateData)
		if errBind != nil {
			return c.JSON(http.StatusBadRequest, responses.WebResponse(http.StatusBadRequest, "error bind data. data not valid", nil))
		}

		fmt.Println("isi update data: ", updateData)

		cfg := config.InitConfig()
		dbRaw := database.InitRawSql(cfg)

		updateDataCore := ReqMidToCore(updateData)
		errCall := tc.service.CallbackMid(dbRaw, updateDataCore)
		if errCall != nil {
			return c.JSON(http.StatusInternalServerError, responses.WebResponse(http.StatusInternalServerError, "error to update data "+errCall.Error(), nil))
		}

		return c.JSON(http.StatusOK, responses.WebResponse(http.StatusOK, "transaction success", nil))
	}
}

// func AmountHandler(c echo.Context) error {
// 	// Menerima nilai 'amount' dari permintaan HTTP
// 	amountStr := c.FormValue("amount")

// 	// Validasi dan Konversi ke Floating-Point
// 	amount, err := strconv.ParseFloat(amountStr, 64)

// 	// Penanganan Kesalahan Konversi
// 	if err != nil {
// 		// Tangani kesalahan, misalnya, kirim respons kesalahan ke klien
// 		log.Error("Error parsing amount:", zap.Error(err))
// 		return c.JSON(http.StatusBadRequest, ErrorResponse{"Invalid amount"})
// 	}

// 	// Penggunaan Nilai 'amount' yang Valid
// 	// Di sini Anda dapat melakukan apa pun yang diperlukan dengan nilai 'amount' yang telah diuji dan dikonversi.
// 	// Misalnya, menyimpannya ke database atau menggunakan nilainya dalam logika bisnis lainnya.

// 	// Kembalikan respons berhasil jika semua langkah sebelumnya berhasil
// 	return c.JSON(http.StatusOK, map[string]interface{}{"amount": amount})
// }
