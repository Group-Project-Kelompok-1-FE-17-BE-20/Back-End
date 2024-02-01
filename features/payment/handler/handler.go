package handler

import (
	"fmt"
	"net/http"
	"strings"

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
		request := createPaymentRequest{}
		//result, err := middlewares.ExtractTokenUserId(c)
		_, err := middlewares.ExtractToken(c)

		if err != nil {
			log.Error("missing or malformed JWT")
			return c.JSON(http.StatusUnauthorized, responses.ResponseFormat(http.StatusUnauthorized, "", "Missing or Malformed JWT", nil, nil))
		}

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
		midtransResponse := midtransCallback{}
		errBind := c.Bind(&midtransResponse)
		if errBind != nil {
			log.Sugar().Errorf("error on binding notification input", errBind)
			return c.JSON(http.StatusBadRequest, responses.ResponseFormat(http.StatusBadRequest, "", "Bad request: "+errBind.Error(), nil, nil))
		}

		log.Sugar().Infof("callback midtrans status: %s, order ID: %s, transaction ID: %s",
			midtransResponse.TransactionStatus, midtransResponse.OrderID, midtransResponse.TransactionID)

		err := tc.service.UpdatePayment(RequestToCore(midtransResponse))
		if err != nil {
			if strings.Contains(err.Error(), "not found") {
				return c.JSON(http.StatusNotFound, responses.ResponseFormat(http.StatusNotFound, "", "The requested resource was not found", nil, nil))
			} else if strings.Contains(err.Error(), "no payment record has been updated") {
				return c.JSON(http.StatusNotFound, responses.ResponseFormat(http.StatusNotFound, "", "No payment record has been updated", nil, nil))
			}
			return c.JSON(http.StatusInternalServerError, responses.ResponseFormat(http.StatusInternalServerError, "", "Internal server error", nil, nil))
		}

		return c.JSON(http.StatusOK, responses.ResponseFormat(http.StatusOK, "", "Successful updated payment status", nil, nil))
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
