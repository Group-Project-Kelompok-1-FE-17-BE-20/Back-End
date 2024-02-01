package handler

import (
	"Laptop/features/payment"
	"Laptop/utils"
)

type paymentResponse struct {
	ID          string          `json:"id"`
	OrderID     string          `json:"order_id"`
	Amount      string          `json:"amount"`
	BankAccount string          `json:"bank_account"`
	VANumber    string          `json:"va_number"`
	NamaLengkap string          `json:"nama_lengkap"`
	Alamat      string          `json:"alamat"`
	Status      string          `json:"status"`
	CreatedAt   utils.LocalTime `json:"created_at"`
	UpdatedAt   utils.LocalTime `json:"updated_at"`
}

func paymentResp(p payment.PaymentCore) paymentResponse {
	return paymentResponse{
		ID:          p.ID,
		OrderID:     p.OrderID,
		Amount:      p.Amount,
		BankAccount: p.BankAccount,
		VANumber:    p.VANumber,
		NamaLengkap: p.NamaLengkap,
		Alamat:      p.Alamat,
		Status:      p.Status,
		CreatedAt:   utils.LocalTime(p.CreatedAt),
		UpdatedAt:   utils.LocalTime(p.UpdatedAt),
	}
}
