package handler

import (
	"Laptop/features/user"
	"time"
)

type LoginResponse struct {
	ID    uint   `json:"id"`
	Email string `json:"email"`
	Token string `json:"token"`
	// Role  string `json:"role"`
}
type UserResponse struct {
	ID           uint   `json:"UserID"`
	Username     string `json:"username"`
	Name         string `json:"nama_lengkap"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	NomorHP      string `json:"nomor_hp"`
	Alamat       string `json:"alamat"`
	JenisKelamin string `json:"jenis_kelamin"`
	ImageProfil  string `json:"image_profil"`
	// NamaToko     string    `json:"nama_toko" `
	// AlamatToko   string    `json:"alamat_toko" `
	// ImageToko    string    `json:"image_toko"`
	CreatedAt time.Time `json:"created_at"`
}

// mapping from userCore to UserResponse
func MapCoreUserToRes(Core user.CoreUser) UserResponse {
	return UserResponse{
		ID:           Core.ID,
		Username:     Core.Username,
		Name:         Core.NamaLengkap,
		Email:        Core.Email,
		NomorHP:      Core.NomorHP,
		Alamat:       Core.Alamat,
		JenisKelamin: Core.JenisKelamin,
		ImageProfil:  Core.ImageProfil,
		// NamaToko:     Core.NamaToko,
		// AlamatToko:   Core.AlamatToko,
		// ImageToko:    Core.ImageToko,
		CreatedAt: Core.CreatedAt,
	}

}

// mapping from userCore to LoginResponse
func MapCoreUserToLogRes(Core user.CoreUser, jwtToken string) LoginResponse {
	return LoginResponse{
		ID:    Core.ID,
		Email: Core.Email,
		Token: jwtToken,
	}
}
