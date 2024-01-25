package handler

import (
	"Laptop/features/user"
)

type LoginRequest struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type UserRequest struct {
	ID           uint   `form:"id"`
	Username     string `json:"username" form:"username"`
	Name         string `json:"nama_lengkap" form:"nama_lengkap"`
	Email        string `json:"email" form:"email"`
	Password     string `json:"password" form:"password"`
	NomorHP      string `json:"nomor_hp" form:"nomor_hp"`
	Alamat       string `json:"alamat" form:"alamat"`
	JenisKelamin string `json:"jenis_kelamin" form:"jenis_kelamin"`
	ImageProfil  string `json:"image_profil" form:"image_profil"`
	// NamaToko     string `json:"nama_toko" form:"nama_toko"`
	// AlamatToko   string `json:"alamat_toko" form:"alamat_toko"`
	// ImageToko    string `json:"image_toko" form:"image_toko"`
}

// Mapping dari struct requet to struct core
func MapReqToCoreUser(req UserRequest) user.CoreUser {
	return user.CoreUser{
		ID:           req.ID,
		Username:     req.Username,
		NamaLengkap:  req.Name,
		Email:        req.Email,
		Password:     req.Password,
		NomorHP:      req.NomorHP,
		Alamat:       req.Alamat,
		JenisKelamin: req.JenisKelamin,
		ImageProfil:  req.ImageProfil,
		// NamaToko:     req.NamaToko,
		// AlamatToko:   req.AlamatToko,
		// ImageToko:    req.ImageToko,
	}
}
