package handler

import (
	"Laptop/features/payment"
)

type createPaymentRequest struct {
	OrderID     string `json:"order_id"`
	Amount      string `json:"amount"`
	NamaLengkap string `gorm:"not null" json:"nama_lengkap" form:"nama_lengkap"`
	Alamat      string `gorm:"type:string" json:"alamat" form:"alamat"`
	BankAccount string `json:"bank_account"`
}

type midtransCallback struct {
	// TransactionTime   string `json:"transaction_time"`
	TransactionStatus string `json:"transaction_status"`
	// TransactionID       string `json:"transaction_id"`
	// StatusMessage       string `json:"status_message"`
	// StatusCode          string `json:"status_code"`
	// SignatureKey        string `json:"signature_key"`
	// PaymentType         string `json:"payment_type"`
	OrderID string `json:"order_id"`
	// MerchantID          string `json:"merchant_id"`
	// MaskedCard          string `json:"masked_card"`
	// GrossAmount         string `json:"gross_amount"`
	// FraudStatus         string `json:"fraud_status"`
	// ECI                 string `json:"eci"`
	// Currency            string `json:"currency"`
	// ChannelResponseMsg  string `json:"channel_response_message"`
	// ChannelResponseCode string `json:"channel_response_code"`
	// CardType            string `json:"card_type"`
	// Bank                string `json:"bank"`
	// ApprovalCode        string `json:"approval_code"`
}

func ReqMidToCore(input midtransCallback) payment.PaymentCore {
	return payment.PaymentCore{
		OrderID: input.OrderID,
		Status:  input.TransactionStatus,
	}
}

func RequestToCore(data interface{}) payment.PaymentCore {
	res := payment.PaymentCore{}
	switch v := data.(type) {
	case createPaymentRequest:
		res.OrderID = v.OrderID
		res.Amount = v.Amount
		res.NamaLengkap = v.NamaLengkap
		res.Alamat = v.Alamat
		res.BankAccount = v.BankAccount
	case midtransCallback:
		//res.ID = v.TransactionID
		res.OrderID = v.OrderID
		res.Status = v.TransactionStatus
	default:
		return payment.PaymentCore{}
	}
	return res
}
