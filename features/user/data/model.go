package data

import (
	"Laptop/features/user"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username     string `gorm:"column:username;"`
	NamaLengkap  string `gorm:"column:nama_lengkap;not null"`
	Email        string `gorm:"column:email;not null;unique"`
	Password     string `gorm:"column:password;not null"`
	NomorHP      string `gorm:"column:nomer_hp;"`
	Alamat       string `gorm:"column:alamat;"`
	JenisKelamin string `gorm:"column:jenis_kelamin;"`
	ImageProfil  string `gorm:"column:image_profil;"`
	NamaToko     string `gorm:"column:nama_toko;"`
	AlamatToko   string `gorm:"column:alamat_toko;"`
	ImageToko    string `gorm:"column:image_toko;"`
}

// mapping coreUser to User
func MapCoreUsertoUser(coreUser1 user.CoreUser) User {

	var user = &User{
		Username:     coreUser1.Username,
		NamaLengkap:  coreUser1.NamaLengkap,
		Email:        coreUser1.Email,
		Password:     coreUser1.Password,
		NomorHP:      coreUser1.NomorHP,
		Alamat:       coreUser1.Alamat,
		JenisKelamin: coreUser1.JenisKelamin,
		ImageProfil:  coreUser1.ImageProfil,
		NamaToko:     coreUser1.NamaToko,
		AlamatToko:   coreUser1.AlamatToko,
		ImageToko:    coreUser1.ImageToko,

		Model: gorm.Model{ID: uint(coreUser1.ID)},
	}
	return *user

}

// Mapping User to CoreUser
func MapUserToCoreUser(user1 User) user.CoreUser {

	coreUser := &user.CoreUser{
		ID:           user1.ID,
		Username:     user1.Username,
		NamaLengkap:  user1.NamaLengkap,
		Email:        user1.Email,
		Password:     user1.Password,
		NomorHP:      user1.NomorHP,
		Alamat:       user1.Alamat,
		JenisKelamin: user1.JenisKelamin,
		ImageProfil:  user1.ImageProfil,
		NamaToko:     user1.NamaToko,
		ImageToko:    user1.ImageToko,
		AlamatToko:   user1.AlamatToko,

		CreatedAt: user1.CreatedAt,
		UpdatedAt: user1.UpdatedAt,
		DeletedAt: user1.DeletedAt.Time,
	}
	return *coreUser

}

// Mapping dari []User ke []CoreUser

func ListMapUserToCoreUser(users []User) []user.CoreUser {
	coreUsers := make([]user.CoreUser, len(users))

	for i, users := range users {
		coreUser := &user.CoreUser{
			Username:     users.Username,
			NamaLengkap:  users.NamaLengkap,
			Email:        users.Email,
			Password:     users.Password,
			NomorHP:      users.NomorHP,
			Alamat:       users.Alamat,
			JenisKelamin: users.JenisKelamin,
			ImageProfil:  users.ImageProfil,
			NamaToko:     users.NamaToko,
			ImageToko:    users.ImageToko,
			AlamatToko:   users.AlamatToko,
		}
		coreUsers[i] = *coreUser
	}

	return coreUsers
}
