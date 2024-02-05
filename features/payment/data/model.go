package data

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	config "Laptop/app/configs"
	"Laptop/app/database"
	"Laptop/features/payment"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
)

func paymentModels(p database.Payment) payment.PaymentCore {
	return payment.PaymentCore{
		ID:          p.ID,
		OrderID:     p.OrderID,
		Amount:      p.Amount,
		BankAccount: p.BankAccount,
		VANumber:    p.VANumber,
		NamaLengkap: p.NamaLengkap,
		Alamat:      p.Alamat,
		Status:      p.Status,
		CreatedAt:   p.CreatedAt,
		UpdatedAt:   p.UpdatedAt,
	}
}

func paymentEntities(p payment.PaymentCore) database.Payment {
	return database.Payment{
		ID:          p.ID,
		OrderID:     p.OrderID,
		Amount:      p.Amount,
		BankAccount: p.BankAccount,
		VANumber:    p.VANumber,
		NamaLengkap: p.NamaLengkap,
		Alamat:      p.Alamat,
		Status:      p.Status,
		CreatedAt:   p.CreatedAt,
		UpdatedAt:   p.UpdatedAt,
	}
}

func CoreToModel(input payment.PaymentCore) database.Payment {
	return database.Payment{
		Status:  input.Status,
		OrderID: input.OrderID,
	}
}

func chargeMidtrans(request payment.PaymentCore) database.Payment {
	var bankMap = map[string]coreapi.BankTransferDetails{
		"bni": {
			Bank: midtrans.BankBni,
		},
		"bca": {
			Bank: midtrans.BankBca,
		},
		"bri": {
			Bank: midtrans.BankBri,
		},
	}

	var c = coreapi.Client{}
	c.New(config.MIDTRANS_SERVERKEY, midtrans.Sandbox)

	amount, err := strconv.ParseInt(request.Amount, 10, 64)
	if err != nil {
		return database.Payment{}
	}

	log.Info(request.BankAccount)
	bankTransfer, ok := bankMap[request.BankAccount]
	if !ok {
		return database.Payment{}
	}
	fmt.Printf("logbank: %v\n", request.BankAccount)
	fmt.Printf("logbank tranfer: %v\n", bankTransfer)

	req := &coreapi.ChargeReq{
		PaymentType:  coreapi.PaymentTypeBankTransfer,
		BankTransfer: &bankTransfer,
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  request.OrderID,
			GrossAmt: amount,
		},
	}

	resp, _ := c.ChargeTransaction(req)
	banks := make([]string, len(resp.VaNumbers))
	for i, bank := range resp.VaNumbers {
		banks[i] = bank.Bank
	}
	banksStr := strings.Join(banks, ",")

	fmt.Printf("log bankstr: %v\n", banksStr)

	vaNumbers := make([]string, len(resp.VaNumbers))
	for i, vaNumber := range resp.VaNumbers {
		vaNumbers[i] = vaNumber.VANumber
	}

	vaNumbersStr := strings.Join(vaNumbers, ",")
	createdAt, _ := time.Parse(time.RFC3339, resp.TransactionTime)

	//-------------------
	return database.Payment{
		ID:          resp.TransactionID,
		OrderID:     resp.OrderID,
		Amount:      resp.GrossAmount,
		NamaLengkap: request.NamaLengkap,
		Alamat:      request.Alamat,
		BankAccount: banksStr,
		VANumber:    vaNumbersStr,
		Status:      resp.TransactionStatus,
		CreatedAt:   createdAt,
	}
}
