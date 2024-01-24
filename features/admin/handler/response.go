package handler

import (
	"Laptop/features/admin"
	"time"
)

type LoginResponse struct {
	Id    uint   `json:"id"`
	Email string `json:"email"`
	Token string `json:"token"`
	// Role  string `json:"role"`
}
type AdminResponse struct {
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
}

// mapping from userCore to UserResponse
func MapCoreAdminToRes(Core admin.CoreAdmin) AdminResponse {
	return AdminResponse{
		// Username:     Core.Username,
		// Name:         Core.NamaLengkap,
		Email: Core.Email,
		// NomorHP:      Core.NomorHP,
		// Alamat:       Core.Alamat,
		// JenisKelamin: Core.JenisKelamin,
		// ImageProfil:  Core.ImageProfil,
		CreatedAt: Core.CreatedAt,
	}

}

// mapping from userCore to LoginResponse
func MapCoreAdminToLogRes(Core admin.CoreAdmin, jwtToken string) LoginResponse {
	return LoginResponse{
		Id:    Core.ID,
		Email: Core.Email,
		Token: jwtToken,
	}
}
