package handler

import (
	"Laptop/features/admin"
)

type LoginRequest struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type AdminRequest struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

// Mapping dari struct requet to struct core
func MapReqToCoreAdmin(req AdminRequest) admin.CoreAdmin {
	return admin.CoreAdmin{
		// Username:     req.Username,
		// NamaLengkap:  req.Name,
		Email:    req.Email,
		Password: req.Password,
		// NomorHP:      req.NomorHP,
		// Alamat:       req.Alamat,
		// JenisKelamin: req.JenisKelamin,
		// ImageProfil:  req.ImageProfil,
	}
}
