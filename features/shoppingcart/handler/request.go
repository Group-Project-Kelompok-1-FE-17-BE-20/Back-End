package handler

import (
	"Laptop/features/shoppingcart"
)

type CartRequest struct {
	UserID uint   `json:"user_id" form:"user_id"`
	Status string `gorm:"type:string" json:"status" form:"status"`
}

// Mapping dari struct TaskRequest To struct Core Task
func CartReqToCore(req CartRequest) shoppingcart.CoreCart {
	return shoppingcart.CoreCart{
		UserID: req.UserID,
		Status: req.Status,
	}
}
